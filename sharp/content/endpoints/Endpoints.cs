using System.Security.Claims;
using content.domainservice;
using content.repository;

namespace content.endpoints;

public class VoteRequest {
  public long VideoId { get; set; }

  public short Type { get; set; }
}

public class VideoRequest {
  public string Title { get; set; } = string.Empty;
  public string Des { get; set; } = string.Empty;
  public string CoverUrl { get; set; } = string.Empty;
  public string VideoUrl { get; set; } = string.Empty;
  public int Duration { get; set; }
  public string Category { get; set; } = string.Empty;
  public string Tags { get; set; } = string.Empty;
}

public class Endpoints {
  public Task<Pagination<VideoDto>>
  UserVideos(IDomainService service, long userId, long? page,
             int size) => service.FindByUserId(userId, page ?? long.MaxValue,
                                               size);

  public Task<Pagination<VideoDto>>
  Videos(IDomainService service, long? page,
         int size) => service.FindRecent(page ?? long.MaxValue, size);

  public Task<IReadOnlyList<VideoDto>>
  Likes(IDomainService service, long userId, long page,
        int size) => service.VotedVideos(userId, page, size);

  public void Vote(IDomainService service, VoteRequest request) => service.Vote(
      request.Type switch {
        0 => VoteType.CancelVote,
        1 => VoteType.Vote,
        _ => throw new ArgumentOutOfRangeException(nameof(request.Type))
      },
      request.VideoId);

  public void CreateVideo(IDomainService service, ClaimsPrincipal user,
                          VideoRequest request) {
    var video =
        new Video { Title = request.Title,       Des = request.Des,
                    CoverUrl = request.CoverUrl, VideoUrl = request.VideoUrl,
                    Duration = request.Duration, Category = request.Category,
                    Tags = request.Tags,         UserId = user.UserId() };
    service.Save(video);
  }
}

public static class EndpointsExtension {
  public static IServiceCollection AddEndpoints(
      this IServiceCollection services) => services.AddSingleton<Endpoints>();

  public static void MapEndpoints(this IEndpointRouteBuilder endpoints) {
    var service = endpoints.ServiceProvider.GetService<Endpoints>() ??
                  throw new ArgumentNullException(nameof(Endpoints));
    endpoints.MapGet("/users/{userId:long}/videos", service.UserVideos);
    endpoints.MapGet("/users/{userId:long}/likes", service.Likes);
    endpoints.MapGet("/videos", service.Videos);
    endpoints.MapPost("/videos", service.CreateVideo).RequireAuthorization();
    endpoints.MapPost("/votes", service.Vote).RequireAuthorization();
    endpoints.MapPost("/votes/cancel", service.Vote).RequireAuthorization();
  }
}