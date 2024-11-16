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

using content.domainservice;
using Moq;
using content.repository;

namespace content.Tests;

public class DomainTest
{

    private (Mock<IVideoRepository>, Mock<IUserRepository> userRepo, Mock<IVoteRepository> voteRepo, Mock<SearchClient> searchClient) Setup() =>
        (new Mock<IVideoRepository>(), new Mock<IUserRepository>(), new Mock<IVoteRepository>(), new Mock<SearchClient>(null!));

    [Fact]
    public async Task FindById_ReturnsVideoDto_WhenVideoExists()
    {
        // Arrange
        var (mockVideoRepo, mockUserRepo, mockVoteRepo, mockSearchClient) = Setup();
        var video = new Video { Id = 1, UserId = 1 };
        var user = new User { Id = "1" };
        mockVideoRepo.Setup(repo => repo.FindById(1)).ReturnsAsync(video);
        mockUserRepo.Setup(repo => repo.FindById(1)).ReturnsAsync(user);
        mockVoteRepo.Setup(repo => repo.VotedOfVideos(It.IsAny<List<long>>())).ReturnsAsync([1]);
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object, mockSearchClient.Object);

        // Act
        var result = await service.FindById(1);

        // Assert
        Assert.NotNull(result);
        Assert.Equal("1", result.Id);
        Assert.NotNull(result.Author);
        Assert.Equal("1", result.Author.Id);
        Assert.True(result.IsLiked);
    }

    [Fact]
    public async Task FindAllByIds_ReturnsVideoDtos_WhenVideosExist()
    {
        // Arrange
        var (mockVideoRepo, mockUserRepo, mockVoteRepo, mockSearchClient) = Setup();
        var videos = new List<Video> { new() { Id = 1, UserId = 1 }, new() { Id = 2, UserId = 2 } };
        var users = new List<User> { new() { Id = "1" }, new() { Id = "2" } };
        mockVideoRepo.Setup(repo => repo.FindAllByIds(It.IsAny<IReadOnlyList<long>>())).ReturnsAsync(videos);
        mockUserRepo.Setup(repo => repo.FindAllByIds(It.IsAny<IEnumerable<long>>())).ReturnsAsync(users);
        mockVoteRepo.Setup(repo => repo.VotedOfVideos(It.IsAny<List<long>>())).ReturnsAsync([]);
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object, mockSearchClient.Object);

        // Act
        var result = await service.FindAllByIds([1, 2]);

        // Assert
        Assert.NotNull(result);
        Assert.Equal(2, result.Count);
    }

    [Fact]
    public async Task Save_ReturnsSavedVideoDto_WhenVideoIsSaved()
    {
        // Arrange
        var (mockVideoRepo, mockUserRepo, mockVoteRepo, mockSearchClient) = Setup();
        var video = new Video { Id = 1, UserId = 1 };
        var user = new User { Id = "1" };
        mockVideoRepo.Setup(repo => repo.Save(It.IsAny<Video>())).ReturnsAsync(video);
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object, mockSearchClient.Object);

        // Act
        await service.Save(new Video());
    }


    [Fact]
    public async Task FindByUserId_ReturnsExpectedVideos()
    {
        // Arrange
        var (mockVideoRepo, mockUserRepo, mockVoteRepo, mockSearchClient) = Setup();
        var videos = new List<Video> { new() { Id = 1, UserId = 1 }, new() { Id = 2, UserId = 1 } };
        var user = new User { Id = "1" };
        var voteVideoIds = new List<long> { 1, 2 };
        mockVideoRepo.Setup(repo => repo.FindByUserId(1, 1, 2)).ReturnsAsync(videos);
        mockUserRepo.Setup(repo => repo.FindById(1)).ReturnsAsync(user);
        mockVoteRepo.Setup(repo => repo.VotedOfVideos(It.IsAny<List<long>>())).ReturnsAsync(voteVideoIds);
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object, mockSearchClient.Object);

        // Act
        var result = await service.FindByUserId(1, 1, 2);

        // Assert
        Assert.NotNull(result);
        Assert.Equal(2, result.Items.Count);
    }

    [Fact]
    public async Task FindRecent_ReturnsExpectedVideos()
    {
        // Arrange
        var (mockVideoRepo, mockUserRepo, mockVoteRepo, mockSearchClient) = Setup();
        var videos = new List<Video> { new() { Id = 1, UserId = 1 }, new() { Id = 2, UserId = 2 } };
        var users = new List<User> { new() { Id = "1" }, new() { Id = "2" } };
        var voteVideoIds = new List<long> { 1 };

        mockVideoRepo.Setup(repo => repo.FindRecent(1, 2)).ReturnsAsync(videos);
        mockUserRepo.Setup(repo => repo.FindAllByIds(It.IsAny<IEnumerable<long>>())).ReturnsAsync(users);
        mockVoteRepo.Setup(repo => repo.VotedOfVideos(It.IsAny<List<long>>())).ReturnsAsync(voteVideoIds);
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object, mockSearchClient.Object);
        // Act
        var result = await service.FindRecent(1, 2);

        // Assert
        Assert.NotNull(result);
        Assert.Equal(2, result.Items.Count);
    }

    [Fact]
    public async Task DailyPopularVideos_ReturnsExpectedVideos()
    {
        // Arrange
        var (mockVideoRepo, mockUserRepo, mockVoteRepo, mockSearchClient) = Setup();
        var videos = new List<Video> { new() { Id = 1, UserId = 1 }, new() { Id = 2, UserId = 2 } };
        var users = new List<User> { new() { Id = "1" }, new() { Id = "2" } };
        var voteVideoIds = new List<long> { 1 };
        mockVideoRepo.Setup(repo => repo.DailyPopularVideos(1, 2)).ReturnsAsync((2, videos));
        mockVoteRepo.Setup(repo => repo.VotedOfVideos(It.IsAny<List<long>>())).ReturnsAsync(voteVideoIds);
        mockVideoRepo.Setup(repo => repo.FindAllByIds(It.IsAny<long[]>())).ReturnsAsync(videos);
        mockUserRepo.Setup(repo => repo.FindAllByIds(It.IsAny<IEnumerable<long>>())).ReturnsAsync(users);
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object, mockSearchClient.Object);

        // Act
        var result = await service.DailyPopularVideos(1, 2);

        // Assert
        Assert.NotNull(result);
        Assert.Equal(2, result.Items.Count);
    }


    [Fact]
    public async Task VotedVideos_ReturnsExpectedVideos()
    {
        // Arrange
        var (mockVideoRepo, mockUserRepo, mockVoteRepo, mockSearchClient) = Setup();
        var videoIds = new long[] { 1, 2 };
        var videos = new List<Video> { new() { Id = 1, UserId = 1 }, new() { Id = 2, UserId = 2 } };
        var users = new List<User> { new() { Id = "1" }, new() { Id = "2" } };
        var voteVideoIds = new List<long> { 1 };
        mockVoteRepo.Setup(repo => repo.VotedVideos(1, 2)).ReturnsAsync((2, videoIds));
        mockVoteRepo.Setup(repo => repo.VotedOfVideos(It.IsAny<List<long>>())).ReturnsAsync(voteVideoIds);
        mockVideoRepo.Setup(repo => repo.FindAllByIds(It.IsAny<long[]>())).ReturnsAsync(videos);
        mockUserRepo.Setup(repo => repo.FindAllByIds(It.IsAny<IEnumerable<long>>())).ReturnsAsync(users);
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object, mockSearchClient.Object);

        // Act
        var result = await service.VotedVideos(1, 1, 2);

        // Assert
        Assert.NotNull(result);
        Assert.Equal(2, result.Items.Count);
    }
}