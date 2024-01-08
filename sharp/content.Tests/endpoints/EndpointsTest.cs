using content.domainservice;
using content.endpoints;
using content.repository;
using JetBrains.Annotations;

namespace content.Tests.endpoints;

using Xunit;
using Moq;
using System.Security.Claims;
using System.Collections.Generic;
using System.Threading.Tasks;

[TestSubject(typeof(Endpoints))]
public class EndpointsTests
{
    private readonly Mock<IDomainService> _mockService = new();

    private readonly ClaimsPrincipal _user = new(new ClaimsIdentity(new Claim[]
    {
        new("id", "1"),
    }));

    [Fact]
    public async Task UserVideos_ReturnsExpectedVideos()
    {
        var expectedVideos = new Pagination<VideoDto>
        {
            Items = new List<VideoDto> { new VideoDto { Id = 1 }, new VideoDto { Id = 2 } },
            AllCount = 2
        };
        _mockService.Setup(s => s.FindByUserId(It.IsAny<long>(), It.IsAny<long>(), It.IsAny<int>()))
            .ReturnsAsync(expectedVideos);

        var result = await Endpoints.UserVideos(_mockService.Object, 1, 1, 2);

        Assert.Equal(expectedVideos, result);
    }

    [Fact]
    public async Task Videos_ReturnsExpectedVideos()
    {
        var expectedVideos = new Pagination<VideoDto>
        {
            Items = new List<VideoDto> { new VideoDto { Id = 1 }, new VideoDto { Id = 2 } },
            AllCount = 2
        };
        _mockService.Setup(s => s.FindRecent(It.IsAny<long>(), It.IsAny<int>())).ReturnsAsync(expectedVideos);

        var result = await Endpoints.Videos(_mockService.Object, 1, 2);

        Assert.Equal(expectedVideos, result);
    }

    [Fact]
    public async Task Likes_ReturnsExpectedVideos()
    {
        var expectedVideos = new List<VideoDto> { new VideoDto { Id = 1 }, new VideoDto { Id = 2 } };
        _mockService.Setup(s => s.VotedVideos(It.IsAny<long>(), It.IsAny<long>(), It.IsAny<int>()))
            .ReturnsAsync(expectedVideos);

        var result = await Endpoints.Likes(_mockService.Object, 1, 1, 2);

        Assert.Equal(expectedVideos, result);
    }

    [Fact]
    public void Vote_CallsServiceWithExpectedVoteTypeAndVideoId()
    {
        var request = new VoteRequest { Type = 1, VideoId = 1 };

        Endpoints.Vote(_mockService.Object, request);

        _mockService.Verify(s => s.Vote(VoteType.Vote, request.VideoId), Times.Once);
    }

    [Fact]
    public async void CreateVideo_CallsServiceWithExpectedVideo()
    {
        var request = new VideoRequest
        {
            Title = "Title",
            Des = "Description",
            CoverUrl = "CoverUrl",
            VideoUrl = "http://validurl.com"
        };

        await Endpoints.CreateVideo(_mockService.Object, new Probe(""), _user, request);

        _mockService.Verify(s => s.Save(It.Is<Video>(v =>
            v.Title == request.Title &&
            v.Des == request.Des &&
            v.CoverUrl == request.CoverUrl &&
            v.VideoUrl == request.VideoUrl &&
            1L == _user.UserId())), Times.Once);
    }

    [Fact]
    public void Validate_WithValidVideoRequest_DoesNotThrowException()
    {
        var request = new VideoRequest
        {
            Title = "Valid Title",
            Des = "Valid Description",
            VideoUrl = "http://validurl.com"
        };

        request.Validate();
    }

    [Fact]
    public void Validate_WithTitleNull_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Des = "Valid Description",
            VideoUrl = "http://validurl.com"
        };

        Assert.Throws<ArgumentException>(() => request.Validate());
    }

    [Fact]
    public void Validate_WithTitleEmpty_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = "",
            Des = "Valid Description",
            VideoUrl = "http://validurl.com"
        };

        Assert.Throws<ArgumentException>(() => request.Validate());
    }

    [Fact]
    public void Validate_WithTitleLengthGreaterThan50_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = new string('a', 51),
            Des = "Valid Description",
            VideoUrl = "http://validurl.com"
        };

        Assert.Throws<ArgumentException>(() => request.Validate());
    }

    [Fact]
    public void Validate_WithDesNull_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = "Valid Title",
            VideoUrl = "http://validurl.com"
        };

        Assert.Throws<ArgumentException>(() => request.Validate());
    }

    [Fact]
    public void Validate_WithDesEmpty_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = "Valid Title",
            Des = "",
            VideoUrl = "http://validurl.com"
        };

        Assert.Throws<ArgumentException>(() => request.Validate());
    }

    [Fact]
    public void Validate_WithDesLengthGreaterThan200_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = "Valid Title",
            Des = new string('a', 201),
            VideoUrl = "http://validurl.com"
        };

        Assert.Throws<ArgumentException>(() => request.Validate());
    }

    [Fact]
    public void Validate_WithVideoUrlNull_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = "Valid Title",
            Des = "Valid Description",
        };

        Assert.Throws<ArgumentException>(() => request.Validate());
    }

    [Fact]
    public void Validate_WithVideoUrlEmpty_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = "Valid Title",
            Des = "Valid Description",
            VideoUrl = ""
        };

        Assert.Throws<ArgumentException>(() => request.Validate());
    }

    [Fact]
    public void Validate_WithInvalidVideoUrl_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = "Valid Title",
            Des = "Valid Description",
            VideoUrl = "invalidurl"
        };

        Assert.Throws<ArgumentException>(() => request.Validate());
    }
}