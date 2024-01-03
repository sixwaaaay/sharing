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
}