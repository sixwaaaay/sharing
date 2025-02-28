using content.domainservice;
using content.repository;
using Moq;
using Qdrant.Client;
using Qdrant.Client.Grpc;

public class HistoryServiceTests
{
    [Fact]
    public async Task GetHistory_ReturnsVideos()
    {
        // Arrange
        var mockHistoryRepo = new Mock<HistoryRepository>(null!);
        var mockClient = new Mock<QdrantClient>("localhost",default!,default!,default!,default!,default!);
        var mockDomainService = new Mock<IDomainService>();
        var userId = 1L;
        var historyList = new List<long> { 1L, 2L };
        var videoDtos = new List<VideoDto> { new VideoDto(), new VideoDto() };

        mockHistoryRepo.Setup(x => x.GetHistorys(userId, It.IsAny<long>(), It.IsAny<int>())).ReturnsAsync(historyList);
        mockDomainService.Setup(x => x.FindAllByIds(historyList)).ReturnsAsync(videoDtos);

        var service = new HistoryService(mockHistoryRepo.Object, mockClient.Object, mockDomainService.Object);

        // Act
        var result = await service.GetHistory(userId);

        // Assert
        Assert.Equal(videoDtos, result.Items);
    }

    [Fact]
    public async Task AddHistory_ReturnsTrue()
    {
        // Arrange
        var mockHistoryRepo = new Mock<HistoryRepository>(null!);
        var mockClient = new Mock<QdrantClient>("localhost",default!,default!,default!,default!,default!);
        var mockDomainService = new Mock<IDomainService>();

        var userId = 1L;
        var videoId = 2L;

        mockHistoryRepo.Setup(x => x.AddHistory(userId, videoId)).Returns(Task.FromResult(1L));

        var service = new HistoryService(mockHistoryRepo.Object, mockClient.Object, mockDomainService.Object);

        // Act
        var result = await service.AddHistory(userId, videoId);

        // Assert
        Assert.True(result);
    }

/*     [Fact]
    public async Task Recommendation_ReturnsPagination()
    {
        // Arrange
        var mockHistoryRepo = new Mock<HistoryRepository>(null!);
        var mockClient = new Mock<QdrantClient>("localhost",default!,default!,default!,default!,default!);
        var mockDomainService = new Mock<IDomainService>();

        var userId = 1L;
        var page = 0UL;
        var size = 10UL;
        var historyList = new List<long> { 1L, 2L };
        var recommendationResults = new List<ScoredPoint> { new ScoredPoint { Id = new PointId { Num = 1 } }, new ScoredPoint { Id = new PointId { Num = 2 } } };
        var videoIds = recommendationResults.Select(x => (long)x.Id.Num).ToList();
        var videoDtos = new List<VideoDto> { new (), new () };
        
        mockHistoryRepo.Setup(x => x.GetHistorys(userId, It.IsAny<long?>(), It.IsAny<int?>())).ReturnsAsync(historyList);
        mockClient.Setup(x => x.RecommendAsync(
           It.IsAny<string>(),
           It.IsAny<IReadOnlyList<ulong>>(),
           It.IsAny<IReadOnlyList<ulong>?>(),
           It.IsAny<ReadOnlyMemory<Vector>?>(),
           It.IsAny<ReadOnlyMemory<Vector>?>(),
           It.IsAny<RecommendStrategy?>(),
           It.IsAny<Filter?>(),
           It.IsAny<SearchParams?>(),
           It.IsAny<ulong>(),
           It.IsAny<ulong>(),
           It.IsAny<WithPayloadSelector?>(),
           It.IsAny<WithVectorsSelector?>(),
           It.IsAny<float?>(),
           It.IsAny<string?>(),
           It.IsAny<LookupLocation?>(),
           It.IsAny<ReadConsistency?>(),
           It.IsAny<ShardKeySelector?>(),
           It.IsAny<TimeSpan?>(),
           It.IsAny<CancellationToken>()
       )).ReturnsAsync(recommendationResults);
        mockDomainService.Setup(x => x.FindAllByIds(videoIds)).ReturnsAsync(videoDtos);

        var service = new HistoryService(mockHistoryRepo.Object, mockClient.Object, mockDomainService.Object);

        // Act
        var result = await service.Recommendation(userId, page, size);

        // Assert
        Assert.NotNull(result);
        Assert.Equal(videoDtos, result.Items);
        Assert.Equal((page + 1).ToString(), result.NextPage);
    } */
}