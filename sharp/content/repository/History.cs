using StackExchange.Redis;
using System;

namespace content.repository;

/// <summary> history repository </summary>
public class HistoryRepository(IDatabase db)
{
    /// <summary> fetch historys of specified user </summary>
    public virtual async Task<List<long>> GetHistorys(long userId, long page = 0, int limit = 10)
    {
        var historys = await db.SortedSetRangeByRankAsync($"history:{userId}", -(page + 1) * limit, -1 - page * limit, Order.Descending);
        return [.. historys.Select(x => (long)x)];
    }
    /// <summary> add history to specified user </summary>
    public virtual async Task<long> AddHistory(long userId, long historyId)
    {
        var score = DateTimeOffset.UtcNow.ToUnixTimeMilliseconds();
        var added = await db.SortedSetAddAsync($"history:{userId}", historyId, score, CommandFlags.None);
        return added ? await db.SortedSetLengthAsync($"history:{userId}") : 0;
    }
}