using System.Security.Claims;
using System.Text.Json.Serialization;
using Grpc.Core;
using Grpc.Core.Interceptors;
using Grpc.Net.Client;
using Riok.Mapperly.Abstractions;

namespace content.repository;

public class User
{
    public long Id { get; init; }
    public string Name { get; init; } = string.Empty;
    public bool IsFollow { get; init; }
    public string AvatarUrl { get; set; } = string.Empty;
    public string BgUrl { get; set; } = string.Empty;
    public string Bio { get; set; } = string.Empty;
    public int LikesGiven { get; init; }
    public int LikesReceived { get; init; }
    public int VideosPosted { get; init; }
    public int Following { get; init; }
    public int Followers { get; init; }
}

[Mapper]
public static partial class UserExtension
{
    public static partial User ProtoToUser(this GrpcUser.User user);

    public static partial IReadOnlyList<User> ProtoToUser(this IReadOnlyList<GrpcUser.User> users);
}

public interface IUserRepository
{
    string? Token { get; set; }

    Task<User> FindById(long id);
    Task<IReadOnlyList<User>> FindAllByIds(IEnumerable<long> ids);
}

public class UserRepository : IUserRepository
{
    private readonly GrpcUser.UserService.UserServiceClient _client;

    public UserRepository(ChannelBase channel)
    {
        _client = new GrpcUser.UserService.UserServiceClient(channel.Intercept(Func));
    }

    private Metadata Func(Metadata metadata)
    {
        if (!string.IsNullOrEmpty(Token))
        {
            metadata.Add("Authorization", Token);
        }

        return metadata;
    }

    public string? Token { get; set; }

    public async Task<User> FindById(long id)
    {
        var request = new GrpcUser.GetUserRequest()
        {
            UserId = id
        };
        var reply = await _client.GetUserAsync(request);
        return reply.User.ProtoToUser();
    }

    public async Task<IReadOnlyList<User>> FindAllByIds(IEnumerable<long> ids)
    {
        var request = new GrpcUser.GetUsersRequest()
        {
            UserIds = { ids.Distinct() }
        };
        var reply = await _client.GetUsersAsync(request);
        return reply.Users.ProtoToUser();
    }
}

public interface IVoteRepository
{
    long CurrentUser { get; set; }

    Task UpdateVote(long videoId, VoteType type);

    Task<IReadOnlyList<long>> VotedOfVideos(long[] videoIds);

    Task<IReadOnlyList<long>> VotedVideos(long userId, long page, int size);
}

public enum VoteType
{
    CancelVote,
    Vote
}

public class VoteRepository(HttpClient client) : IVoteRepository
{
    public long CurrentUser { get; set; }

    public async Task UpdateVote(long videoId, VoteType type)
    {
        if (CurrentUser == 0)
        {
            throw new InvalidOperationException("Current user is null.");
        }

        var row = new VoteRow()
        {
            Type = "vote",
            SubjectId = CurrentUser,
            TargetId = videoId,
        };
        var resp = type switch
        {
            VoteType.Vote => await client.PostAsJsonAsync("/item/add", row, VoteJsonContext.Default.VoteRow),
            VoteType.CancelVote => await client.PostAsJsonAsync("/item/delete", row, VoteJsonContext.Default.VoteRow),
            _ => throw new ArgumentOutOfRangeException(nameof(type), type, null)
        };
        resp.EnsureSuccessStatusCode();
    }

    public async Task<IReadOnlyList<long>> VotedOfVideos(long[] videoIds)
    {
        if (videoIds.Length == 0 || CurrentUser == 0)
        {
            return [];
        }

        var req = new ExistsReq()
        {
            Type = "vote",
            SubjectId = CurrentUser,
            TargetIds = videoIds
        };
        var resp = await client.PostAsJsonAsync("/item/exists", req, VoteJsonContext.Default.ExistsReq);
        resp.EnsureSuccessStatusCode();
        var result = await resp.Content.ReadFromJsonAsync(VoteJsonContext.Default.ExistsResp) ?? new ExistsResp();
        return result.Exists;
    }

    public async Task<IReadOnlyList<long>> VotedVideos(long userId, long page, int size)
    {
        var req = new ScanReq()
        {
            Type = "vote",
            SubjectId = userId,
            Token = page,
            Limit = size
        };
        var resp = await client.PostAsJsonAsync("/item/scan", req, VoteJsonContext.Default.ScanReq);
        resp.EnsureSuccessStatusCode();
        var result = await resp.Content.ReadFromJsonAsync(VoteJsonContext.Default.ScanResp) ?? new ScanResp();
        return result.TargetIds;
    }

    internal record VoteRow
    {
        public string Type { get; init; } = string.Empty;
        public long SubjectId { get; init; }
        public string TargetType { get; } = "video";
        public long TargetId { get; init; }
    }

    internal record ExistsReq
    {
        public string Type { get; init; } = string.Empty;
        public long SubjectId { get; init; }
        public string TargetType { get; } = "video";
        public long[] TargetIds { get; init; } = [];
    }

    internal record ExistsResp
    {
        public long[] Exists { get; init; } = [];
    }

    internal record ScanReq
    {
        public string Type { get; init; } = string.Empty;

        public long SubjectId { get; init; }
        public string TargetType { get; } = "video";
        public long Token { get; init; } = long.MaxValue;
        public int Limit { get; init; } = 10;
    }

    internal record ScanResp
    {
        public long[] TargetIds { get; init; } = [];
    }
}

[JsonSourceGenerationOptions(PropertyNamingPolicy = JsonKnownNamingPolicy.SnakeCaseLower)]
[JsonSerializable(typeof(VoteRepository.VoteRow))]
[JsonSerializable(typeof(VoteRepository.ExistsReq))]
[JsonSerializable(typeof(VoteRepository.ExistsResp))]
[JsonSerializable(typeof(VoteRepository.ScanReq))]
[JsonSerializable(typeof(VoteRepository.ScanResp))]
internal partial class VoteJsonContext : JsonSerializerContext;


public static class Extension
{
    public static IServiceCollection AddVoteRepository(this IServiceCollection services) =>
        services.AddScoped<IVoteRepository, VoteRepository>();

    public static IServiceCollection AddUserRepository(this IServiceCollection services) =>
        services.AddScoped<IUserRepository, UserRepository>();

    public static IServiceCollection AddGrpcUser(this IServiceCollection services) =>
        services.AddSingleton<ChannelBase>(sp => GrpcChannel.ForAddress(
            sp.GetRequiredService<IConfiguration>().GetConnectionString("User") ??
            throw new InvalidOperationException(@"User connection string is null.")));

    public static IServiceCollection AddVoteClient(this IServiceCollection services) =>
        services.AddSingleton<HttpClient>(sp =>
        {
            var connectionString = sp.GetRequiredService<IConfiguration>().GetConnectionString("Vote") ??
                                   throw new InvalidOperationException(@"Vote connection string is null.");
            return new HttpClient
            {
                BaseAddress = new Uri(connectionString.TrimEnd('/')),
            };
        });

    public static IApplicationBuilder UseToken(this IApplicationBuilder app) =>
        app.Use(async (context, next) =>
        {
            var userRepository = context.RequestServices.GetService<IUserRepository>() ??
                                 throw new NullReferenceException();
            userRepository.Token = context.Request.Headers.Authorization;

            var voteRepository = context.RequestServices.GetService<IVoteRepository>() ??
                                 throw new NullReferenceException();
            voteRepository.CurrentUser  = context.User.UserId();

            await next.Invoke();
        });
    
    public static long UserId(this ClaimsPrincipal user)
    {
        var id = user.Claims.FirstOrDefault(c => c.Type == "id")?.Value;
        return id == null ? 0 : long.Parse(id);
    }
}
