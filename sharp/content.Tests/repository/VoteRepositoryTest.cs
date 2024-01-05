using System.Net;
using content.repository;
using JetBrains.Annotations;
using Moq;

namespace content.Tests.repository;

[TestSubject(typeof(VoteRepository))]
public class VoteRepositoryTest {
  private readonly HttpClient _httpClient = new HttpClient() {
    BaseAddress = new Uri("http://localhost:8080"),
  };

  [Fact]
  public async Task UpdateVote_ShouldPostVote_WhenVoteTypeIsVote() {
    // Arrange
    var voteRepository = new VoteRepository(_httpClient) { CurrentUser = 3L };

    // Act
    await voteRepository.UpdateVote(1L, VoteType.Vote);
  }

  [Fact]
  public async Task UpdateVote_ShouldPostCancelVote_WhenVoteTypeIsCancelVote() {
    // Arrange
    var voteRepository = new VoteRepository(_httpClient) { CurrentUser = 3L };

    // Act
    await voteRepository.UpdateVote(1L, VoteType.CancelVote);
  }

  [Fact]
  public async Task UpdateVote_ShouldThrowException_WhenVoteTypeIsInvalid() {
    // Arrange
    var httpClient = new Mock<HttpClient>();
    var voteRepository = new VoteRepository(httpClient.Object);
    voteRepository.CurrentUser = 1;
    var videoId = 1L;
    var voteType = (VoteType)100; // Invalid vote type
    // Act & Assert
    await Assert.ThrowsAsync<ArgumentOutOfRangeException>(
        () => voteRepository.UpdateVote(videoId, voteType));
  }

  [Fact]
  public async
      Task VotedOfVideos_ShouldReturnEmptyList_WhenNoVideoIdsProvided() {
    // Arrange
    var httpClient = new Mock<HttpClient>();
    var voteRepository = new VoteRepository(httpClient.Object);
    voteRepository.CurrentUser = 1;

    // Act
    var result = await voteRepository.VotedOfVideos(new long[0]);

    // Assert
    Assert.Empty(result);
  }

  [Fact]
  public async
      Task VotedOfVideos_ShouldReturnEmptyList_WhenCurrentUserIsZero() {
    // Arrange
    var voteRepository = new VoteRepository(_httpClient) {
      CurrentUser = 2,
    };

    // Act
    var result =
        await voteRepository.VotedOfVideos(new long[] { 12345, 23456, 4567 });

    // Assert
    Assert.NotNull(result);
    Assert.Empty(result);

    result = await voteRepository.VotedOfVideos([]);

    // Assert
    Assert.NotNull(result);
    Assert.Empty(result);

    voteRepository.CurrentUser = 0;

    // Act
    result = await voteRepository.VotedOfVideos([12345, 23456, 4567]);
    // Assert
    Assert.NotNull(result);
    Assert.Empty(result);
  }

  [Fact]
  public async Task VotedVideos() {
    // Arrange
    var voteRepository = new VoteRepository(_httpClient) {
      CurrentUser = 2,
    };
    await voteRepository.UpdateVote(5555L, VoteType.Vote);

    // Act
    var result = await voteRepository.VotedVideos(2, 0, 10);

    // Assert

    Assert.NotNull(result);
    Assert.NotEmpty(result);
  }
}