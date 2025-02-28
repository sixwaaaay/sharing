using System.Net;
using System.Net.Http.Json;
using System.Text.Json.Serialization;
using content.repository;
using Moq;
using Moq.Protected;
using StackExchange.Redis;

namespace content.Tests.repository;

public class ClientTest
{
    [Fact]
    public async Task VotedOfVideos_ReturnsEmptyList_WhenVideoIdsIsEmpty()
    {
        // Arrange
        var mockHttpMessageHandler = new Mock<HttpMessageHandler>();
        var httpClient = new HttpClient(mockHttpMessageHandler.Object);
        var voteRepository = new VoteRepository(httpClient);

        // Act
        var result = await voteRepository.VotedOfVideos([]);

        // Assert
        Assert.Empty(result);
    }

    [Fact]
    public async Task VotedOfVideos_ReturnsListOfVotedVideos()
    {
        // Arrange
        var videoIds = new List<long> { 1, 2, 3 };
        var expectedResponse = new List<long> { 1, 2 };

        var mockHttpMessageHandler = new Mock<HttpMessageHandler>();
        mockHttpMessageHandler.Protected()
            .Setup<Task<HttpResponseMessage>>(
                "SendAsync",
                ItExpr.IsAny<HttpRequestMessage>(),
                ItExpr.IsAny<CancellationToken>()
            )
            .ReturnsAsync(new HttpResponseMessage
            {
                StatusCode = HttpStatusCode.OK,
                Content = JsonContent.Create(expectedResponse, VoteJsonContext.Default.ListInt64)
            });

        var httpClient = new HttpClient(mockHttpMessageHandler.Object)
        {
            BaseAddress = new Uri("http://localhost:5151")
        };
        var voteRepository = new VoteRepository(httpClient)
        {
            Token = "test-token"
        };

        // Act
        var result = await voteRepository.VotedOfVideos(videoIds);

        // Assert
        Assert.Equal(expectedResponse, result);
    }

    [Fact]
    public async Task VotedOfVideos_ThrowsException_WhenResponseIsNotSuccessful()
    {
        // Arrange
        var videoIds = new List<long> { 1, 2, 3 };

        var mockHttpMessageHandler = new Mock<HttpMessageHandler>();
        mockHttpMessageHandler.Protected()
            .Setup<Task<HttpResponseMessage>>(
                "SendAsync",
                ItExpr.IsAny<HttpRequestMessage>(),
                ItExpr.IsAny<CancellationToken>()
            )
            .ReturnsAsync(new HttpResponseMessage
            {
                StatusCode = HttpStatusCode.BadRequest
            });

        var httpClient = new HttpClient(mockHttpMessageHandler.Object)
        {
            BaseAddress = new Uri("http://localhost:5151")
        };
        var voteRepository = new VoteRepository(httpClient)
        {
            Token = "test-token"
        };

        // Act & Assert
        await Assert.ThrowsAsync<HttpRequestException>(() => voteRepository.VotedOfVideos(videoIds));
    }



    [Fact]
    public async Task VotedVideos_ReturnsNextTokenAndTargetIds()
    {
        // Arrange
        var page = 1L;
        var size = 10;
        var expectedResponse = new ScanResp(2, [1, 2, 3]);

        var mockHttpMessageHandler = new Mock<HttpMessageHandler>();
        mockHttpMessageHandler.Protected()
            .Setup<Task<HttpResponseMessage>>(
                "SendAsync",
                ItExpr.IsAny<HttpRequestMessage>(),
                ItExpr.IsAny<CancellationToken>()
            )
            .ReturnsAsync(new HttpResponseMessage
            {
                StatusCode = HttpStatusCode.OK,
                Content = JsonContent.Create(expectedResponse, VoteJsonContext.Default.ScanResp)
            });

        var httpClient = new HttpClient(mockHttpMessageHandler.Object)
        {
            BaseAddress = new Uri("http://localhost:5151")
        };
        var voteRepository = new VoteRepository(httpClient)
        {
            Token = "test-token"
        };

        // Act
        var (nextToken, targetIds) = await voteRepository.VotedVideos(page, size);

        // Assert
        Assert.Equal(expectedResponse.TargetIds, targetIds);
        Assert.Equal(expectedResponse.NextToken, nextToken);
    }

    [Fact]
    public async Task VotedVideos_ThrowsException_WhenResponseIsNotSuccessful()
    {
        // Arrange
        var page = 1L;
        var size = 10;

        var mockHttpMessageHandler = new Mock<HttpMessageHandler>();
        mockHttpMessageHandler.Protected()
            .Setup<Task<HttpResponseMessage>>(
                "SendAsync",
                ItExpr.IsAny<HttpRequestMessage>(),
                ItExpr.IsAny<CancellationToken>()
            )
            .ReturnsAsync(new HttpResponseMessage
            {
                StatusCode = HttpStatusCode.BadRequest
            });

        var httpClient = new HttpClient(mockHttpMessageHandler.Object)
        {
            BaseAddress = new Uri("http://localhost:5151")
        };
        var voteRepository = new VoteRepository(httpClient)
        {
            Token = "test-token"
        };

        // Act & Assert
        await Assert.ThrowsAsync<HttpRequestException>(() => voteRepository.VotedVideos(page, size));
    }

    [Fact]
    public async Task SimilarSearch_ReturnsListOfVideoIds()
    {
        // Arrange
        var videoId = 1;
        var expectedResponse = new Response()
        {
            Hits = [
                new SimilarVideo(2), new SimilarVideo(3) 
            ]
        };

        var mockHttpMessageHandler = new Mock<HttpMessageHandler>();

        var client = new HttpClient(mockHttpMessageHandler.Object)
        {
            BaseAddress = new Uri("http://localhost:5151")
        };

        var searchClient = new SearchClient(client, ConnectionMultiplexer.Connect("localhost").GetDatabase());

        mockHttpMessageHandler.Protected()
            .Setup<Task<HttpResponseMessage>>(
                "SendAsync",
                ItExpr.IsAny<HttpRequestMessage>(),
                ItExpr.IsAny<CancellationToken>()
            )
            .ReturnsAsync(new HttpResponseMessage
            {
                StatusCode = HttpStatusCode.OK,
                Content = JsonContent.Create(expectedResponse, SearchContext.Default.Response)
            });
        
        // Act
        var result = await searchClient.SimilarSearch(videoId);

        // Assert
        Assert.Equal(expectedResponse.Hits.Select(h => h.Id).ToList(), result);


        // Again
        result = await searchClient.SimilarSearch(videoId);

        // Assert
        Assert.Equal(expectedResponse.Hits.Select(h => h.Id).ToList(), result);
    }
}




public record ScanResp(long? NextToken, List<long> TargetIds);

[JsonSourceGenerationOptions(PropertyNamingPolicy = JsonKnownNamingPolicy.SnakeCaseLower)]
[JsonSerializable(typeof(List<long>))]
[JsonSerializable(typeof(ScanResp))]
[JsonSerializable(typeof(InQuery))]
partial class VoteJsonContext : JsonSerializerContext;