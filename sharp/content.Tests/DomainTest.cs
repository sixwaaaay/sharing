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
namespace content.Tests;

using domainservice;
using JetBrains.Annotations;
using Xunit;
using Moq;
using content.repository;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

[TestSubject(typeof(DomainService))]
public class DomainTest
{
    [Fact]
    public async Task FindById_ReturnsVideoDto_WhenVideoExists()
    {
        // Arrange
        var mockVideoRepo = new Mock<IVideoRepository>();
        var mockUserRepo = new Mock<IUserRepository>();
        var mockVoteRepo = new Mock<IVoteRepository>();
        var video = new Video { Id = 1, UserId = 1 };
        var user = new User { Id = 1 };
        mockVideoRepo.Setup(repo => repo.FindById(1)).ReturnsAsync(video);
        mockUserRepo.Setup(repo => repo.FindById(1)).ReturnsAsync(user);
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object);

        // Act
        var result = await service.FindById(1);

        // Assert
        Assert.NotNull(result);
        Assert.Equal(1, result.Id);
        Assert.NotNull(result.Author);
        Assert.Equal(1, result.Author.Id);
    }

    [Fact]
    public async Task FindAllByIds_ReturnsVideoDtos_WhenVideosExist()
    {
        // Arrange
        var mockVideoRepo = new Mock<IVideoRepository>();
        var mockUserRepo = new Mock<IUserRepository>();
        var mockVoteRepo = new Mock<IVoteRepository>();
        var videos = new List<Video> { new() { Id = 1, UserId = 1 }, new() { Id = 2, UserId = 2 } };
        var users = new List<User> { new() { Id = 1 }, new() { Id = 2 } };
        mockVideoRepo.Setup(repo => repo.FindAllByIds(It.IsAny<long[]>())).ReturnsAsync(videos);
        mockUserRepo.Setup(repo => repo.FindAllByIds(It.IsAny<IEnumerable<long>>())).ReturnsAsync(users);
        mockVoteRepo.Setup(repo => repo.VotedOfVideos(It.IsAny<long[]>())).ReturnsAsync(new List<long>());
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object);

        // Act
        var result = await service.FindAllByIds([1, 2]);

        // Assert
        Assert.NotNull(result);
        Assert.Equal(2, result.Count());
    }

    [Fact]
    public async Task Save_ReturnsSavedVideoDto_WhenVideoIsSaved()
    {
        // Arrange
        var mockVideoRepo = new Mock<IVideoRepository>();
        var mockUserRepo = new Mock<IUserRepository>();
        var mockVoteRepo = new Mock<IVoteRepository>();
        var video = new Video { Id = 1, UserId = 1 };
        var user = new User { Id = 1 };
        mockVideoRepo.Setup(repo => repo.Save(It.IsAny<Video>())).ReturnsAsync(video);
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object);

        // Act
        await service.Save(new Video());
    }


    [Fact]
    public async Task FindByUserId_ReturnsExpectedVideos()
    {
        // Arrange
        var mockVideoRepo = new Mock<IVideoRepository>();
        var mockUserRepo = new Mock<IUserRepository>();
        var mockVoteRepo = new Mock<IVoteRepository>();
        var videos = new List<Video> { new() { Id = 1, UserId = 1 }, new() { Id = 2, UserId = 1 } };
        var user = new User { Id = 1 };
        var voteVideoIds = new List<long> { 1, 2 };
        mockVideoRepo.Setup(repo => repo.FindByUserId(1, 1, 2)).ReturnsAsync(videos);
        mockUserRepo.Setup(repo => repo.FindById(1)).ReturnsAsync(user);
        mockVoteRepo.Setup(repo => repo.VotedOfVideos(It.IsAny<long[]>())).ReturnsAsync(voteVideoIds);
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object);

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
        var mockVideoRepo = new Mock<IVideoRepository>();
        var mockUserRepo = new Mock<IUserRepository>();
        var mockVoteRepo = new Mock<IVoteRepository>();
        var videos = new List<Video> { new() { Id = 1, UserId = 1 }, new() { Id = 2, UserId = 2 } };
        var users = new List<User> { new() { Id = 1 }, new() { Id = 2 } };
        var voteVideoIds = new List<long> { 1 };
        
        mockVideoRepo.Setup(repo => repo.FindRecent(1, 2)).ReturnsAsync(videos);
        mockUserRepo.Setup(repo => repo.FindAllByIds(It.IsAny<IEnumerable<long>>())).ReturnsAsync(users);
        mockVoteRepo.Setup(repo => repo.VotedOfVideos(It.IsAny<long[]>())).ReturnsAsync(voteVideoIds);
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object);
        // Act
        var result = await service.FindRecent(1, 2);

        // Assert
        Assert.NotNull(result);
        Assert.Equal(2, result.Items.Count);
    }

    [Fact]
    public async Task VotedVideos_ReturnsExpectedVideos()
    {
        // Arrange
        var mockVideoRepo = new Mock<IVideoRepository>();
        var mockUserRepo = new Mock<IUserRepository>();
        var mockVoteRepo = new Mock<IVoteRepository>();
        var videoIds = new long[] { 1, 2 };
        var videos = new List<Video> { new() { Id = 1, UserId = 1 }, new() { Id = 2, UserId = 2 } };
        var users = new List<User> { new() { Id = 1 }, new() { Id = 2 } };
        var voteVideoIds = new List<long> { 1 };
        mockVoteRepo.Setup(repo => repo.VotedVideos(1, 1, 2)).ReturnsAsync(videoIds);
        mockVoteRepo.Setup(repo => repo.VotedOfVideos(It.IsAny<long[]>())).ReturnsAsync(voteVideoIds);
        mockVideoRepo.Setup(repo => repo.FindAllByIds(It.IsAny<long[]>())).ReturnsAsync(videos);
        mockUserRepo.Setup(repo => repo.FindAllByIds(It.IsAny<IEnumerable<long>>())).ReturnsAsync(users);
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object);

        // Act
        var result = await service.VotedVideos(1, 1, 2);

        // Assert
        Assert.NotNull(result);
        Assert.Equal(2, result.Count);
    }

    [Fact]
    public async Task Vote_UpdatesVoteSuccessfully()
    {
        // Arrange
        var mockVideoRepo = new Mock<IVideoRepository>();
        var mockUserRepo = new Mock<IUserRepository>();
        var mockVoteRepo = new Mock<IVoteRepository>();
        mockVoteRepo.Setup(repo => repo.UpdateVote(1, VoteType.Vote)).Returns(Task.CompletedTask);
        mockVideoRepo.Setup(repo => repo.UpdateVoteCounter(1, VoteType.Vote)).Returns(Task.CompletedTask);
        var service = new DomainService(mockVideoRepo.Object, mockUserRepo.Object, mockVoteRepo.Object);

        // Act
        await service.Vote(VoteType.Vote, 1);

        // Assert
        mockVoteRepo.Verify(repo => repo.UpdateVote(1, VoteType.Vote), Times.Once());
        mockVideoRepo.Verify(repo => repo.UpdateVoteCounter(1, VoteType.Vote), Times.Once());
    }
}