using content.repository;
using Moq;
using StackExchange.Redis;

public class HistoryRepositoryTests
{
    [Fact]
    public async Task GetHistorys_ShouldReturnHistorys()
    {
        // Arrange
        var mockDb = new Mock<IDatabase>();
        var (userId, historyIds) = (123, new RedisValue[] { 1, 2, 3 });
        mockDb.Setup(db => db.ListRangeAsync($"history:{userId}", It.IsAny<long>(), It.IsAny<long>(), CommandFlags.None)).ReturnsAsync(historyIds);
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
        mockDb.Setup(db => db.ListRightPushAsync($"history:{userId}", historyId, When.Always, CommandFlags.None)).ReturnsAsync(expectedCount);
        var repository = new HistoryRepository(mockDb.Object);
        // Act
        var result = await repository.AddHistory(userId, historyId);
        // Assert
        Assert.Equal(expectedCount, result);
    }
}