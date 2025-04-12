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
using Npgsql;

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

    public DateTime CreatedAt { get; init; } = DateTime.UtcNow;

    public DateTime UpdatedAt { get; init; } = DateTime.UtcNow;

    public short Processed { get; init; } = 1;
}

public interface IVideoRepository
{
    Task<Video> FindById(long id);
    Task<IReadOnlyList<Video>> FindAllByIds(IReadOnlyList<long> ids);
    Task<IReadOnlyList<Video>> FindByUserId(long userId, long page, int size);
    Task<IReadOnlyList<Video>> FindRecent(long page, int size);
    Task<(long, IReadOnlyList<Video>)> DailyPopularVideos(long page, int size);
    Task<Video> Save(Video video);
    Task IncrementViewCount(long id);
}

public class VideoRepository(NpgsqlDataSource dataSource) : IVideoRepository
{
    const string columns = "id, user_id, title, des, cover_url, video_url, duration, view_count, like_count, created_at, updated_at, processed";

    public async Task<Video> FindById(long id)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        return await connection.QuerySingleAsync<Video>(
            $"SELECT {columns} FROM videos WHERE id = @id",
            new { id });
    }

    /// <summary> Find all the videos by the ids. </summary>
    /// <param name="ids"> The ids of the videos. </param>
    /// <returns> The list of videos. </returns>
    public async Task<IReadOnlyList<Video>> FindAllByIds(IReadOnlyList<long> ids)
    {
        if (ids.Count == 0)
        {
            return [];
        }
        await using var connection = await dataSource.OpenConnectionAsync();
        var videos = await connection.QueryAsync<Video>(
            $"SELECT {columns} FROM videos WHERE id = ANY(@ids) ORDER BY array_position(@ids, id)",
            new { ids });
        return videos.ToList();
    }

    /// <summary>  Find the videos by the user id. </summary>
    /// <param name="userId"> The user id. </param>
    /// <param name="page"> The page number. </param>
    /// <param name="size"> The size of the page. </param>
    /// <returns> The list of videos. </returns>
    public async Task<IReadOnlyList<Video>> FindByUserId(long userId, long page, int size)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        var videos = await connection.QueryAsync<Video>(
            $"SELECT {columns} FROM videos WHERE processed = 1 AND user_id = @userId AND id < @page ORDER BY id DESC LIMIT @size",
            new { userId, page, size });
        return videos.ToList();
    }

    /// <summary> Find the recent videos. </summary>
    /// <param name="page"> The page number. </param>
    /// <param name="size"> The size of the page. </param>
    /// <returns> The list of videos. </returns>
    public async Task<IReadOnlyList<Video>> FindRecent(long page, int size)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        var videos = await connection.QueryAsync<Video>(
            $"SELECT {columns} FROM videos WHERE processed = 1 AND id < @page ORDER BY id DESC LIMIT @size",
            new { page, size });
        return videos.ToList();
    }


    internal class RankedVideo
    {
        public long Id { get; set; }
        public long OrderNum { get; set; }
    };

    /// <summary> Get the daily popular videos. </summary>
    /// <param name="page"> The page number.</param>
    /// <param name="size">  The size of the page. </param>
    /// <returns> A tuple of the next page token and the list of videos. </returns>
    public async Task<(long, IReadOnlyList<Video>)> DailyPopularVideos(long page, int size)
    {
        var ranks = await queryRanks(page, size);
        var rankedVideos = ranks.ToList();
        var nextToken = rankedVideos.Count == size ? rankedVideos[^1].OrderNum : 0;
        var ids = rankedVideos.Select(r => r.Id).ToArray();
        var videos = await FindAllByIds(ids);
        return (nextToken, videos);
        async ValueTask<IEnumerable<RankedVideo>> queryRanks(long page, int size)
        {
            await using var connection = await dataSource.OpenConnectionAsync();
            return await connection.QueryAsync<RankedVideo>(
                "SELECT id, order_num FROM popular_videos WHERE order_num > @page ORDER BY order_num DESC LIMIT @size",
                new { page, size });
        }
    }

    /// <summary> Save the video. </summary>
    /// <param name="video"> The video to save. </param>
    /// <returns> The saved video. </returns>
    public async Task<Video> Save(Video video)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        var result = await connection.QuerySingleAsync<Video>(
            "INSERT INTO videos (user_id, title, des, cover_url, video_url, duration, view_count, like_count, created_at, updated_at, processed) VALUES (@UserId, @Title, @Des, @CoverUrl, @VideoUrl, @Duration, @ViewCount, @LikeCount, @CreatedAt, @UpdatedAt, @Processed) " +
            $"returning {columns}",
            video
        );
        return result;
    }

    public async Task IncrementViewCount(long id)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        await connection.ExecuteAsync(
            "UPDATE videos SET view_count = view_count + 1 WHERE id = @id",
            new { id }
        );
    }

}

/// <summary> The extensions for the video repository. </summary>
public static class VideoRepositoryExtensions
{
    public static IServiceCollection AddVideoRepository(this IServiceCollection services) =>
        services.AddSingleton<IVideoRepository, VideoRepository>();
}