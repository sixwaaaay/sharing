using System.Text.Json.Serialization;
using content.repository;
using Qdrant.Client;
using Qdrant.Client.Grpc;

namespace content.domainservice;

public record AddVideoHistory([property: JsonNumberHandling(JsonNumberHandling.AllowReadingFromString)] long VideoId);

public class HistoryService(HistoryRepository history, QdrantClient client, IDomainService domain)
{

    public async Task<Pagination<VideoDto>> GetHistory(long userId, long page = 0, int size = 10)
    {
        var historyList = await history.GetHistorys(userId, page, size);
        return new Pagination<VideoDto>() { Items = await domain.FindAllByIds(historyList), NextPage = historyList.Count >= size ? (page + 1).ToString() : null };
    }

    public async Task<bool> AddHistory(long userId, long videoId)
    {
        await history.AddHistory(userId, videoId);
        return true;
    }

    private readonly Random random = new();

    public async Task<Pagination<VideoDto>> Recommendation(long userId, ulong page = 0, ulong size = 10)
    {
        var historyList = await history.GetHistorys(userId, 0, 0);
        var positiveVectors = historyList.Count == 0 ? new Vector[] { new(Enumerable.Range(1, 1024).Select(_ => (float)random.NextDouble()).ToArray()) } : null;
        var result = await client.RecommendAsync("videos", positive: historyList.Select(x => (ulong)x).ToList(), /* history */
            positiveVectors: positiveVectors, /* if empty positive point */
            limit: size, offset: page * size
        );
        return new Pagination<VideoDto>() { Items = await domain.FindAllByIds([.. result.Select(x => (long)x.Id.Num)]), NextPage = 30 > size * (page + 1) ? (page + 1).ToString() : "0" };
    }
}

public static class HistoryServiceInject
{
    public static IServiceCollection AddHistoryService(this IServiceCollection services) => services
        .AddSingleton(sp => new QdrantClient(new Uri(sp.GetRequiredService<IConfiguration>().GetConnectionString("QdrantUrl")!),
            apiKey: sp.GetRequiredService<IConfiguration>().GetConnectionString("QdrantApiKey")!))
        .AddScoped<HistoryRepository>().AddScoped<HistoryService>();
}