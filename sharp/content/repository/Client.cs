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

using System.Net.Http.Headers;
using System.Security.Claims;
using System.Text.Json;
using System.Text.Json.Serialization;
using StackExchange.Redis;

namespace content.repository;

public class User
{
    public string Id { get; init; } = string.Empty;
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

public interface IUserRepository
{
    string? Token { get; set; }

    /// <summary> Find user information by id. </summary>
    /// <param name="id"> User id. </param>
    /// <returns> User information. </returns>
    Task<User> FindById(long id);

    /// <summary> Find user information by id list. </summary>
    /// <param name="ids"> User id list. </param>
    /// <returns> User information list. </returns>
    Task<IReadOnlyList<User>> FindAllByIds(IEnumerable<long> ids);
}

public class UserRepository(HttpClient client) : IUserRepository
{
    public string? Token { get; set; }

    /// <summary> Find user information by id. </summary>
    public async Task<User> FindById(long id)
    {
        var req = new HttpRequestMessage(HttpMethod.Get, $"/users/{id}");
        req.Headers.Authorization = new AuthenticationHeaderValue("Bearer", Token);
        Console.WriteLine(client.BaseAddress);
        using var resp = await client.SendAsync(req);
        resp.EnsureSuccessStatusCode();
        return await resp.Content.ReadFromJsonAsync(UserJsonContext.Default.User) ?? new();
    }


    /// <summary> Find user information by id list. </summary>
    public async Task<IReadOnlyList<User>> FindAllByIds(IEnumerable<long> ids)
    {
        var req = new HttpRequestMessage(HttpMethod.Get,
            $"/users?{string.Join("&", ids.Distinct().Select(id => $"ids={id}"))}");
        req.Headers.Authorization = new AuthenticationHeaderValue("Bearer", Token);
        using var resp = await client.SendAsync(req);
        resp.EnsureSuccessStatusCode();
        return await resp.Content.ReadFromJsonAsync(UserJsonContext.Default.IReadOnlyListUser) ?? [];
    }
}

[JsonSourceGenerationOptions(PropertyNamingPolicy = JsonKnownNamingPolicy.SnakeCaseLower)]
[JsonSerializable(typeof(IReadOnlyList<User>))]
public partial class UserJsonContext : JsonSerializerContext;

public interface IVoteRepository
{
    string? Token { get; set; }

    /// <summary> Get voted status of videos. </summary>
    /// <param name="videoIds"> Video ids. </param>
    /// <returns> Voted status of videos. </returns>
    Task<IReadOnlyList<long>> VotedOfVideos(List<long> videoIds);


    /// <summary> Scan voted videos, which means paging through all voted videos. </summary>
    /// <param name="page"> Page token. </param>
    /// <param name="size"> Page size. </param>
    /// <returns> Page token and voted videos. </returns>    
    Task<(long?, IReadOnlyList<long>)> VotedVideos(long page, int size);
}

public class VoteRepository(HttpClient client) : IVoteRepository
{
    public string? Token { get; set; }

    /// <summary> Get voted status of videos. </summary>
    public async Task<IReadOnlyList<long>> VotedOfVideos(List<long> videoIds)
    {
        if (string.IsNullOrEmpty(Token) || videoIds.Count == 0)
        {
            return [];
        }
        using var req = new HttpRequestMessage(HttpMethod.Post, "/graph/videos/likes");
        req.Content = JsonContent.Create(new InQuery(videoIds), VoteJsonContext.Default.InQuery);

        if (!string.IsNullOrEmpty(Token) && AuthenticationHeaderValue.TryParse(Token, out var auth))
        {
            req.Headers.Authorization = auth;
        }


        using var resp = await client.SendAsync(req);
        resp.EnsureSuccessStatusCode();

        var result = await resp.Content.ReadFromJsonAsync(VoteJsonContext.Default.ListInt64) ?? [];
        return result;
    }


    /// <summary> Scan voted videos, which means paging through all voted videos. </summary>
    public async Task<(long?, IReadOnlyList<long>)> VotedVideos(long page, int size)
    {
        using var req = new HttpRequestMessage(HttpMethod.Get, $"/graph/videos?page={page}&size={size}");
        if (!string.IsNullOrEmpty(Token) && AuthenticationHeaderValue.TryParse(Token, out var auth))
        {
            req.Headers.Authorization = auth;
        }

        using var resp = await client.SendAsync(req);

        resp.EnsureSuccessStatusCode();

        var result = await resp.Content.ReadFromJsonAsync(VoteJsonContext.Default.ScanResp) ?? new ScanResp();
        return (result.NextToken, result.TargetIds);
    }

    internal record ScanResp
    {
        public List<long> TargetIds { get; init; } = [];
        public long? NextToken { get; init; }
    }
}

public record InQuery(List<long> ObjectIds);

[JsonSourceGenerationOptions(PropertyNamingPolicy = JsonKnownNamingPolicy.SnakeCaseLower)]
[JsonSerializable(typeof(VoteRepository.ScanResp))]
[JsonSerializable(typeof(InQuery))]
internal partial class VoteJsonContext : JsonSerializerContext;

public class SearchClient(HttpClient client, IDatabase db)
{
    public async Task<IReadOnlyList<long>> SimilarSearch(long videoId)
    {
        var cache = await db.StringGetAsync($"similar:{videoId}");
        if (cache.HasValue)
        {
            return JsonSerializer.Deserialize(cache!, SearchContext.Default.ListInt64) ?? [];
        }

        var body = new RequestBody(videoId, ["id"]);
        var content = JsonContent.Create(body, SearchContext.Default.RequestBody);
        var req = new HttpRequestMessage(HttpMethod.Post, "/indexes/videos/similar") { Content = content };
        using var resp = await client.SendAsync(req);
        resp.EnsureSuccessStatusCode();

        var result = await resp.Content.ReadFromJsonAsync(SearchContext.Default.Response) ?? new Response();

        var ids = result.Hits.Select(h => h.Id).ToList();
        await db.StringSetAsync($"similar:{videoId}", JsonSerializer.Serialize(ids, SearchContext.Default.ListInt64), TimeSpan.FromSeconds(60));

        return ids;
    }
}

public record RequestBody(
    [property: JsonPropertyName("id")] long Id,
    [property: JsonPropertyName("attributesToRetrieve")] string[] AttributesToRetrieve,
    [property: JsonPropertyName("limit")] int Limit = 10);

public record Response()
{
    [JsonPropertyName("hits")] public IReadOnlyList<SimilarVideo> Hits { get; init; } = [];
}

public record SimilarVideo([property: JsonPropertyName("id")] long Id);


[JsonSerializable(typeof(Response))]
[JsonSerializable(typeof(RequestBody))]
[JsonSerializable(typeof(List<long>))]
public partial class SearchContext : JsonSerializerContext;

public static class Extension
{

    /// <summary> get connection string from configuration and ensure not null, trim end '/' </summary>
    public static string ConnString(this IConfiguration config, string name)
        => config.GetConnectionString(name).EnsureNotNull($"{name} connection string is null").TrimEnd('/');

    /// <summary> get connection string from service provider </summary>
    public static string GetConnString(this IServiceProvider sp, string name)
        => sp.GetRequiredService<IConfiguration>().ConnString(name);

    public static IServiceCollection AddVoteRepository(this IServiceCollection services) => services
        .AddHttpClient<IVoteRepository, VoteRepository>("Vote",
            (sp, client) => client.BaseAddress = new Uri(sp.GetConnString("Vote"))).Services;

    public static IServiceCollection AddSearchClient(this IServiceCollection services) => services
        .AddSingleton<IConnectionMultiplexer>(sp => ConnectionMultiplexer.Connect(sp.GetRequiredService<IConfiguration>().GetConnectionString("Redis").EnsureNotNull("Redis connection string is null")))
        .AddSingleton(sp => sp.GetRequiredService<IConnectionMultiplexer>().GetDatabase())
        .AddHttpClient<SearchClient>("Search", (sp, client) =>
        {
            var token = sp.GetRequiredService<IConfiguration>()["Token"];
            client.BaseAddress = new Uri(sp.GetConnString("Search"));
            client.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Bearer", token);
        }).Services;

    public static IServiceCollection AddUserRepository(this IServiceCollection services) =>
        services.AddHttpClient<IUserRepository, UserRepository>("User",
            (sp, client) => client.BaseAddress = new Uri(sp.GetConnString("User"))).Services;

    public static IApplicationBuilder UseToken(this IApplicationBuilder app) =>
        app.Use(async (context, next) =>
        {
            var userRepository = context.RequestServices.GetService<IUserRepository>();
            ArgumentNullException.ThrowIfNull(userRepository, nameof(userRepository));
            var authorization = context.Request.Headers.Authorization;

            userRepository.Token = authorization;

            var voteRepository = context.RequestServices.GetService<IVoteRepository>();
            ArgumentNullException.ThrowIfNull(voteRepository, nameof(voteRepository));
            voteRepository.Token = authorization;

            await next.Invoke();
        });

    public static long UserId(this ClaimsPrincipal user)
    {
        var id = user.Claims.FirstOrDefault(c => c.Type == "id")?.Value;
        return id == null ? 0 : long.Parse(id);
    }

    public static string EnsureNotNull(this string? value, string msg)
    {
        if (string.IsNullOrEmpty(value))
            throw new InvalidOperationException(msg);
        return value;
    }
}