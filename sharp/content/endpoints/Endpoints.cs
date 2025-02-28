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

using System.Security.Claims;
using content.domainservice;
using content.repository;
using FluentValidation;

namespace content.endpoints;

public class VoteRequest
{
    public long VideoId { get; set; }

    public short Type { get; set; }
}

public class VideoRequest
{
    public string Title { get; set; } = string.Empty;
    public string Des { get; set; } = string.Empty;
    public string CoverUrl { get; set; } = string.Empty;
    public string VideoUrl { get; set; } = string.Empty;
}

public static class Endpoints
{
    public static Task<Pagination<VideoDto>> UserVideos(IDomainService service, long userId, long? page, int? size)
    {
        EnsurePageAndSize(page, size);
        return service.FindByUserId(userId, page ?? long.MaxValue, size ?? 10);
    }

    public static Task<Pagination<VideoDto>> Videos(IDomainService service, long? page, int? size)
    {
        EnsurePageAndSize(page, size);
        return service.FindRecent(page ?? long.MaxValue, size ?? 10);
    }

    public static Task<Pagination<VideoDto>> RecommendVideos(HistoryService service, ClaimsPrincipal user, long? page, int? size)
    {
        EnsurePageAndSize(page, size);
        return service.Recommendation(user.UserId(), (ulong?)page ?? 0, (ulong?)size ?? 10);
    }

    public static Task<Pagination<VideoDto>> HistoryVideos(HistoryService service,ClaimsPrincipal  user, long? page, int? size)
    {
        EnsurePageAndSize(page, size);
        return service.GetHistory(user.UserId());
    }

    public static Task<VideoDto> FindVideoById(IDomainService service, long id) => service.FindById(id);

    public static Task<Pagination<VideoDto>> DailyPopularVideos(IDomainService service, long? page, int? size)
    {
        EnsurePageAndSize(page, size);
        return service.DailyPopularVideos(page ?? long.MaxValue, size ?? 10);
    }


    public static Task<Pagination<VideoDto>> Likes(IDomainService service, long userId, long? page, int? size)
    {
        EnsurePageAndSize(page, size);
        return service.VotedVideos(userId, page ?? long.MaxValue, size ?? 10);
    }

    public static Task<IReadOnlyList<VideoDto>> SimilarVideos(IDomainService service, long id) =>
        service.FindSimilarVideos(id);

    public static async Task CreateVideo(IDomainService service, IProbe probe, ClaimsPrincipal user,
        VideoRequest request, VideoRequestValidator validator)
    {
        await validator.ValidateAndThrowAsync(request);

        var duration = await probe.GetVideoDuration(request.VideoUrl);
        var video = new Video
        {
            Title = request.Title,
            Des = request.Des,
            Duration = (int)(!string.IsNullOrWhiteSpace(duration) ? double.Parse(duration) : 0),
            CoverUrl = request.CoverUrl,
            VideoUrl = request.VideoUrl,
            UserId = user.UserId()
        };

        await service.Save(video);
    }



    public static Task AddHistory(HistoryService service,ClaimsPrincipal user , AddVideoHistory addVideo) {
        return service.AddHistory(user.UserId(), addVideo.VideoId);
    }

    public static void MapEndpoints(this IEndpointRouteBuilder endpoints)
    {
        endpoints.MapGet("/users/{userId:long}/videos", UserVideos).WithName("getUserVideos");
        endpoints.MapGet("/users/{userId:long}/likes", Likes).RequireAuthorization().WithName("getUserLikes");
        endpoints.MapGet("/videos", Videos).WithName("getVideos");
        endpoints.MapGet("/videos/recommend", RecommendVideos).WithName("getRecommendVideos");
        endpoints.MapGet("/videos/{id:long}", FindVideoById).WithName("getVideo");
        endpoints.MapGet("/videos/{id:long}/similar", SimilarVideos).WithName("getSimilarVideos");
        endpoints.MapPost("/videos/popular", DailyPopularVideos).WithName("getDailyPopularVideos");
        endpoints.MapPost("/videos", CreateVideo).RequireAuthorization().WithName("createVideo");
        endpoints.MapPost("/videos/history", AddHistory).RequireAuthorization().WithName("addVideoHistory");
        endpoints.MapGet("/videos/history", HistoryVideos).RequireAuthorization().WithName("getVideoHistory");
    }


    public static void EnsurePageAndSize(long? page, int? size)
    {
        if (page is < 0)
        {
            throw new ArgumentOutOfRangeException(nameof(page));
        }

        if (size is < 0 or > 20)
        {
            throw new ArgumentOutOfRangeException(nameof(size));
        }
    }
}

public class VideoRequestValidator : AbstractValidator<VideoRequest>
{
    public VideoRequestValidator()
    {
        RuleFor(x => x.Title).NotEmpty().MaximumLength(50)
            .WithMessage("title is null or empty or length greater than 50");
        RuleFor(x => x.Des).NotEmpty().MaximumLength(200)
            .WithMessage("description is null or empty or length greater than 200");
        RuleFor(x => x.VideoUrl).NotEmpty().Must(x => Uri.IsWellFormedUriString(x, UriKind.Absolute))
            .WithMessage("video url is null or empty or not a valid url");
    }
}