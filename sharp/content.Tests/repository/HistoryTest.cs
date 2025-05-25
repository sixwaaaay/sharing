using content.repository;
using Moq;
using StackExchange.Redis;
using System;

public class HistoryRepositoryTests
{
    [Fact]
    public async Task GetHistorys_ShouldReturnHistorys()
    {
        // Arrange
        var mockDb = new Mock<IDatabase>();
        var (userId, historyIds) = (123, new RedisValue[] { 1, 2, 3 });
        mockDb.Setup(db => db.SortedSetRangeByRankAsync($"history:{userId}", It.IsAny<long>(), It.IsAny<long>(), Order.Descending, CommandFlags.None)).ReturnsAsync(historyIds);
        var repository = new HistoryRepository(mockDb.Object);

        // Act
        var result = await repository.GetHistorys(userId);

        // Assert
        Assert.Equal(historyIds.Length, result.Count);
        for (int i = 0; i < historyIds.Length; i++)
        {
            Assert.Equal((long)historyIds[i], result[i]);
        }
    }

    [Fact]
    public async Task AddHistory_ShouldReturnCount()
    {
        // Arrange
        var mockDb = new Mock<IDatabase>();
        var (userId, historyId, expectedCount) = (123, 456, 1);
        mockDb.Setup(db => db.SortedSetAddAsync($"history:{userId}", historyId, It.IsAny<double>(), CommandFlags.None)).ReturnsAsync(true);
        mockDb.Setup(db => db.SortedSetLengthAsync($"history:{userId}", double.NegativeInfinity, double.PositiveInfinity, Exclude.None, CommandFlags.None)).ReturnsAsync(expectedCount);
        var repository = new HistoryRepository(mockDb.Object);
        // Act
        var result = await repository.AddHistory(userId, historyId);
        // Assert
        Assert.Equal(expectedCount, result);
    }
}