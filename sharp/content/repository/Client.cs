using Grpc.Core;
using Grpc.Core.Interceptors;
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
