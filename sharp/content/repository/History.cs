using StackExchange.Redis;

namespace content.repository;

/// <summary> history repository </summary>
public class HistoryRepository(IDatabase db)
{
    /// <summary> fetch historys of specified user </summary>
    public virtual async Task<List<long>> GetHistorys(long userId, long page = 0, int limit = 10)
    {
        var historys = await db.ListRangeAsync($"history:{userId}", -(page + 1) * limit, -1 - page * limit);
        return [.. historys.Select(x => (long)x)];
    }
    /// <summary> add history to specified user </summary>
    public virtual async Task<long> AddHistory(long userId, long historyId) => await db.ListRightPushAsync($"history:{userId}", historyId);
}