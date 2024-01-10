/*
 * Copyright (c) 2023-2024 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * 
 */
using Dapper;
using MySqlConnector;

[module: DapperAot]

namespace content.repository;

public record Video
{
    public long Id { get; init; }
    public long UserId { get; init; }
    public string Title { get; init; } = string.Empty;

    public string Des { get; init; } = string.Empty;

    public string CoverUrl { get; init; } = string.Empty;


    public string VideoUrl { get; init; } = string.Empty;

    public int Duration { get; init; }

    public int ViewCount { get; init; }

    public int LikeCount { get; init; }

    public DateTime CreatedAt { get; init; }

    public DateTime UpdatedAt { get; init; }

    public short Processed { get; init; } = 1;
}

public interface IVideoRepository
{
    Task<Video> FindById(long id);
    Task<IReadOnlyList<Video>> FindAllByIds(long[] ids);
    Task<IReadOnlyList<Video>> FindByUserId(long userId, long page, int size);
    Task<IReadOnlyList<Video>> FindRecent(long page, int size);
    Task<Video> Save(Video video);
    Task UpdateVoteCounter(long id, VoteType type);
}

public class VideoRepository(MySqlDataSource dataSource) : IVideoRepository
{
    public async Task<Video> FindById(long id)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        return await connection.QuerySingleAsync<Video>(
            "SELECT id, user_id, title, des, cover_url, video_url, duration, view_count, like_count, created_at, updated_at, processed " +
            "FROM videos WHERE id = @id",
            new { id });
    }

    [DapperAot(false)]
    public async Task<IReadOnlyList<Video>> FindAllByIds(long[] ids)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        var inClause = new InClause<long>(ids);

        var query =
            $"SELECT id, user_id, title, des, cover_url, video_url, duration, view_count, like_count, created_at, updated_at, processed " +
            $"FROM videos WHERE id in {inClause.Condition}";
        var command = connection.CreateCommand();
        command.CommandText = query;
        inClause.BindParam(command);
        var result = await command.ExecuteReaderAsync();
        var videos = new List<Video>();
        while (await result.ReadAsync())
        {
            videos.Add(new Video
            {
                Id = result.GetInt64(0),
                UserId = result.GetInt64(1),
                Title = result.GetString(2),
                Des = result.GetString(3),
                CoverUrl = result.GetString(4),
                VideoUrl = result.GetString(5),
                Duration = result.GetInt32(6),
                ViewCount = result.GetInt32(7),
                LikeCount = result.GetInt32(8),
                CreatedAt = result.GetDateTime(9),
                UpdatedAt = result.GetDateTime(10),
                Processed = result.GetInt16(11)
            });
        }

        return videos;
    }

    public async Task<IReadOnlyList<Video>> FindByUserId(long userId, long page, int size)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        var videos = await connection.QueryAsync<Video>(
            "SELECT id,user_id, title, des, cover_url, video_url, duration, view_count, like_count, created_at, updated_at, processed " +
            "FROM videos FORCE INDEX(user_created) WHERE processed = 1 AND  user_id = @userId AND id < @page ORDER BY id DESC LIMIT @size",
            new { userId, page, size });
        return videos.ToList();
    }

    public async Task<IReadOnlyList<Video>> FindRecent(long page, int size)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        var videos = await connection.QueryAsync<Video>(
            "SELECT id ,user_id, title, des, cover_url, video_url, duration, view_count, like_count, created_at, updated_at, processed " +
            "FROM videos FORCE INDEX (processed) WHERE processed = 1 AND id < @page ORDER BY id DESC LIMIT @size",
            new { page, size });
        return videos.ToList();
    }

    public async Task<Video> Save(Video video)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        await using var transaction = await connection.BeginTransactionAsync();
        var result = await connection.QuerySingleAsync<long>(
            "INSERT INTO videos (user_id, title, des, cover_url, video_url, duration, view_count, like_count, created_at, updated_at, processed) VALUES (@UserId, @Title, @Des, @CoverUrl, @VideoUrl, @Duration, @ViewCount, @LikeCount, @CreatedAt, @UpdatedAt, @Processed); SELECT LAST_INSERT_ID();",
            new
            {
                video.UserId,
                video.Title,
                video.Des,
                video.CoverUrl,
                video.VideoUrl,
                video.Duration,
                video.ViewCount,
                video.LikeCount,
                video.CreatedAt,
                video.UpdatedAt,
                video.Processed
            }, transaction);
        var affectedRows = await connection.ExecuteAsync(
            "UPDATE counter SET counter = counter + 1 WHERE id = @id",
            new { id = video.UserId }, transaction);
        if (affectedRows == 0)
        {
            await connection.ExecuteAsync(
                "INSERT INTO counter (id, counter) VALUES (@Id, 1)",
                new { id = video.UserId }, transaction);
        }

        await transaction.CommitAsync();
        return await FindById(result);
    }

    public async Task UpdateVoteCounter(long id, VoteType type)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        await connection.ExecuteAsync(
            type switch
            {
                VoteType.Vote => "UPDATE videos SET like_count = like_count + 1 WHERE id = @id",
                VoteType.CancelVote => "UPDATE videos SET like_count = like_count - 1 WHERE id = @id",
                _ => throw new ArgumentOutOfRangeException(nameof(type), type, null)
            },
            new { id });
    }
}

public static class VideoRepositoryExtensions
{
    public static IServiceCollection AddVideoRepository(this IServiceCollection services) =>
        services.AddSingleton<IVideoRepository, VideoRepository>();
}

/// <summary>
/// Represents a SQL IN clause for a list of values of type T.
/// </summary>
/// <typeparam name="T">The type of the values in the IN clause.</typeparam>
internal class InClause<T>(IEnumerable<T> values)
{
    /// <summary>
    /// Gets the parameters for the IN clause, each with a unique name and a value.
    /// </summary>
    private IEnumerable<(string, T)> Parameters =>
        values.Select((value, index) => ($"p{index}", value));

    /// <summary>
    /// Gets the condition for the IN clause, which can be used in a SQL query.
    /// </summary>
    public string Condition => $"({string.Join(", ", Parameters.Select(p => $"@{p.Item1}"))})";

    /// <summary>
    /// Adds the parameters for the IN clause to the specified SQL command.
    /// </summary>
    /// <param name="command">The SQL command to which the parameters will be added.</param>
    /// <returns>The same SQL command, for chaining calls.</returns>
    public MySqlCommand BindParam(MySqlCommand command)
    {
        foreach (var (key, value) in Parameters)
        {
            command.Parameters.AddWithValue(key, value);
        }

        return command;
    }
}