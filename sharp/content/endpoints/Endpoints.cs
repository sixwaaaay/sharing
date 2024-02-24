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
using Microsoft.AspNetCore.Mvc;

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

public class MessageRequest
{
    public long ReceiverId { get; set; }
    public string Content { get; set; } = string.Empty;
    public short Type { get; set; }
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


    public static void Vote(IDomainService service, VoteRequest request) => service.Vote(
        request.Type switch
        {
            0 => VoteType.CancelVote,
            1 => VoteType.Vote,
            _ => throw new ArgumentOutOfRangeException(nameof(request.Type))
        }, request.VideoId);

    public static async Task CreateVideo(IDomainService service, IProbe probe, ClaimsPrincipal user,
        VideoRequest request)
    {
        request.Validate();
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


    public static Task<Pagination<MessageDto>> FindMessages(IMessageDomain service, long receiverId, long? page,
        int? size, bool unreadOnly)
    {
        EnsurePageAndSize(page, size);
        return service.FindMessages(receiverId, page ?? 0, size ?? 10, unreadOnly);
    }

    public static Task<MessageDto> Save(IMessageDomain service, MessageRequest request, ClaimsPrincipal user)
    {
        request.Validate();
        var message = new Message
        {
            SenderId = user.UserId(),
            ReceiverId = request.ReceiverId,
            Content = request.Content,
            Type = request.Type
        };

        return service.Save(message);
    }

    public static Task MarkAsRead(IMessageDomain service, long id, ClaimsPrincipal user) => service.MarkAsRead(id, user.UserId());


    
    public static void Validate(this MessageRequest request)
    {
        if (string.IsNullOrWhiteSpace(request.Content) || request.Content.Length > 200)
        {
            throw new ArgumentException("Content is null or empty or length greater than 200", nameof(request.Content));
        }
    }
    
    
    public static void MapEndpoints(this IEndpointRouteBuilder endpoints)
    {
        endpoints.MapGet("/users/{userId:long}/videos", UserVideos);
        endpoints.MapGet("/users/{userId:long}/likes", Likes);
        endpoints.MapGet("/videos", Videos);
        endpoints.MapPost("/videos/popular", DailyPopularVideos);
        endpoints.MapPost("/videos", CreateVideo).RequireAuthorization();
        endpoints.MapPost("/votes", Vote).RequireAuthorization();
        endpoints.MapPost("/votes/cancel", Vote).RequireAuthorization();
        
        
        endpoints.MapGet("/messages", FindMessages).RequireAuthorization();
        endpoints.MapPost("/messages", Save).RequireAuthorization();
        endpoints.MapPost("/messages/{id:long}", MarkAsRead).RequireAuthorization();
    }

    public static void Validate(this VideoRequest request)
    {
        if (string.IsNullOrWhiteSpace(request.Title) || request.Title.Length > 50)
        {
            throw new ArgumentException("Title is null or empty or length greater than 50", nameof(request.Title));
        }

        if (string.IsNullOrWhiteSpace(request.Des) || request.Des.Length > 200)
        {
            throw new ArgumentException("Des is null or empty or length greater than 200", nameof(request.Des));
        }

        if (string.IsNullOrWhiteSpace(request.VideoUrl) ||
            !Uri.IsWellFormedUriString(request.VideoUrl, UriKind.Absolute))
        {
            throw new ArgumentException("Video url is null or empty", nameof(request.VideoUrl));
        }
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