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

using content.repository;
using Riok.Mapperly.Abstractions;

namespace content.domainservice;

public record Pagination<T> where T : class
{
    public int AllCount { get; init; }
    public long? NextPage { get; init; }
    public IReadOnlyList<T> Items { get; init; } = [];
}

public interface IDomainService
{
    Task<VideoDto> FindById(long id);
    Task<IReadOnlyList<VideoDto>> FindAllByIds(long[] ids);
    Task<Pagination<VideoDto>> FindByUserId(long userId, long page, int size);
    Task<Pagination<VideoDto>> FindRecent(long page, int size);
    Task<Pagination<VideoDto>> VotedVideos(long userId, long page, int size);
    Task<Pagination<VideoDto>> DailyPopularVideos(long page, int size);
    Task Save(Video video);

    Task Vote(VoteType type, long videoId);
}

public class DomainService(IVideoRepository videoRepo, IUserRepository userRepo, IVoteRepository voteRepo)
    : IDomainService
{
    public async Task<VideoDto> FindById(long id)
    {
        var video = await videoRepo.FindById(id);
        var videoToVideoDto = video.ToDto();
        var user = await userRepo.FindById(video.UserId);
        videoToVideoDto.Author = user;
        return videoToVideoDto;
    }

    public async Task<IReadOnlyList<VideoDto>> FindAllByIds(long[] ids)
    {
        var videos = await videoRepo.FindAllByIds(ids);
        return await CurrentUserVotedVideos(videos);
    }

    private async Task<IReadOnlyList<VideoDto>> CurrentUserVotedVideos(IReadOnlyList<Video> videos)
    {
        var userIds = videos.Select(v => v.UserId);
        var userTask = userRepo.FindAllByIds(userIds);
        var voteVideoIdsTask = voteRepo.VotedOfVideos(videos.Select(v => v.Id).ToArray());
        var users = await userTask;
        var voteVideoIds = await voteVideoIdsTask;
        return Compose(users, videos, voteVideoIds).ToList();
    }


    private static IEnumerable<VideoDto> Compose(IEnumerable<User> users, IEnumerable<Video> videos,
        IEnumerable<long> voteVideoIds)
    {
        var userDict = users.ToDictionary(u => u.Id);
        var voteSet = new HashSet<long>(voteVideoIds);
        return videos.Select(video =>
        {
            var dto = video.ToDto();
            if (userDict.TryGetValue(video.UserId, out var value))
            {
                dto.Author = value;
            }

            dto.IsLiked = voteSet.Contains(video.Id);
            return dto;
        });
    }

    public async Task<Pagination<VideoDto>> FindByUserId(long userId, long page, int size)
    {
        var videos = await videoRepo.FindByUserId(userId, page, size);
        var findById = userRepo.FindById(userId);
        var votedOfVideos = voteRepo.VotedOfVideos(videos.Select(v => v.Id).ToArray());
        var user = await findById;
        var voteVideoIds = await votedOfVideos;
        var videoDtos = Compose([user], videos, voteVideoIds).ToList();
        return new Pagination<VideoDto>()
        {
            Items = videoDtos,
            NextPage = videoDtos.Count == size ? videoDtos[^1].Id : null
        };
    }

    public async Task<Pagination<VideoDto>> FindRecent(long page, int size)
    {
        var videos = await videoRepo.FindRecent(page, size);
        var videoDtos = await CurrentUserVotedVideos(videos);
        return new Pagination<VideoDto>()
        {
            Items = videoDtos,
            NextPage = videoDtos.Count == size ? videoDtos[^1].Id : default
        };
    }

    public Task Save(Video video) => videoRepo.Save(video);

    public async Task<Pagination<VideoDto>> VotedVideos(long userId, long page, int size)
    {
        var (token, videoIds) = await voteRepo.VotedVideos(userId, page, size);
        return new Pagination<VideoDto>
        {
            Items = await FindAllByIds(videoIds.ToArray()),
            NextPage = token
        };
    }

    public async Task<Pagination<VideoDto>> DailyPopularVideos(long page, int size)
    {
        var (token,videos) = await videoRepo.DailyPopularVideos(page, size);
        var videoDtos = await CurrentUserVotedVideos(videos);
        return new Pagination<VideoDto>
        {
            Items = videoDtos,
            NextPage = token
        };
    }

    public async Task Vote(VoteType type, long videoId)
    {
        await voteRepo.UpdateVote(videoId, type);
        await videoRepo.UpdateVoteCounter(videoId, type);
    }
}

[Mapper]
public static partial class VideoMapper
{
    [MapperIgnoreSource(nameof(Video.UserId))]
    [MapperIgnoreTarget(nameof(VideoDto.Author))]
    [MapperIgnoreTarget(nameof(VideoDto.IsLiked))]
    public static partial VideoDto ToDto(this Video video);
}

public record VideoDto
{
    public long Id { get; init; }
    public User? Author { get; set; }
    public string Title { get; init; } = string.Empty;
    public string Des { get; init; } = string.Empty;
    public string CoverUrl { get; init; } = string.Empty;
    public string VideoUrl { get; init; } = string.Empty;
    public int Duration { get; init; }
    public int ViewCount { get; init; }
    public int LikeCount { get; init; }
    public DateTime CreatedAt { get; init; }
    public DateTime UpdatedAt { get; init; }
    public short Processed { get; init; }
    public bool IsLiked { get; set; }
}

public static class DomainServiceExtensions
{
    public static IServiceCollection AddDomainService(this IServiceCollection services) =>
        services.AddScoped<IDomainService, DomainService>();
}