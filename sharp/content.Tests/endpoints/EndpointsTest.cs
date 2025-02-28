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
using content.endpoints;
using content.repository;
using FluentValidation;
using Moq;
using System.Security.Claims;

namespace content.Tests.endpoints;
public class EndpointsTests
{
    private readonly Mock<IDomainService> _mockService = new();

    private readonly ClaimsPrincipal _user = new(new ClaimsIdentity([new("id", "1")]));

    [Fact]
    public async Task UserVideos_ReturnsExpectedVideos()
    {
        var expectedVideos = new Pagination<VideoDto>
        {
            Items = [new VideoDto { Id = "1" }, new VideoDto { Id = "2" }],
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
            Items = [new VideoDto { Id = "1" }, new VideoDto { Id = "2" }],
            AllCount = 2
        };
        _mockService.Setup(s => s.FindRecent(It.IsAny<long>(), It.IsAny<int>())).ReturnsAsync(expectedVideos);

        var result = await Endpoints.Videos(_mockService.Object, 1, 2);

        Assert.Equal(expectedVideos, result);
    }

    [Fact]
    public async Task DailyPopularVideos_ReturnsExpectedVideos()
    {
        var expectedVideos = new Pagination<VideoDto>
        {
            Items = [new VideoDto { Id = "1" }, new VideoDto { Id = "2" }],
            AllCount = 2
        };
        _mockService.Setup(s => s.DailyPopularVideos(It.IsAny<long>(), It.IsAny<int>())).ReturnsAsync(expectedVideos);

        var result = await Endpoints.DailyPopularVideos(_mockService.Object, 1, 2);

        Assert.Equal(expectedVideos, result);
    }


    [Fact]
    public async Task Likes_ReturnsExpectedVideos()
    {
        var expectedVideos = new Pagination<VideoDto>
        {
            Items = [new VideoDto { Id = "1" }, new VideoDto { Id = "2" }],
            AllCount = 2
        };
        _mockService.Setup(s => s.VotedVideos(It.IsAny<long>(), It.IsAny<long>(), It.IsAny<int>()))
            .ReturnsAsync(expectedVideos);

        var result = await Endpoints.Likes(_mockService.Object, 1, 1, 2);

        Assert.Equal(expectedVideos, result);
    }


    [Fact]
    public async Task CreateVideo_CallsServiceWithExpectedVideo()
    {
        var request = new VideoRequest
        {
            Title = "Title",
            Des = "Description",
            CoverUrl = "CoverUrl",
            VideoUrl = "https://validurl.com"
        };

        await Endpoints.CreateVideo(_mockService.Object, new Probe(""), _user, request, new VideoRequestValidator());

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
            VideoUrl = "https://validurl.com"
        };

        var result = new VideoRequestValidator().Validate(request);
        Assert.True(result.IsValid);
    }

    [Fact]
    public void Validate_WithTitleNull_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Des = "Valid Description",
            VideoUrl = "https://validurl.com"
        };

        Assert.Throws<ValidationException>(() => new VideoRequestValidator().ValidateAndThrow(request));
    }

    [Fact]
    public void Validate_WithTitleEmpty_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = "",
            Des = "Valid Description",
            VideoUrl = "https://validurl.com"
        };

        Assert.Throws<ValidationException>(() => new VideoRequestValidator().ValidateAndThrow(request));
    }

    [Fact]
    public void Validate_WithTitleLengthGreaterThan50_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = new string('a', 51),
            Des = "Valid Description",
            VideoUrl = "https://validurl.com"
        };

        Assert.Throws<ValidationException>(() => new VideoRequestValidator().ValidateAndThrow(request));
    }

    [Fact]
    public void Validate_WithDesNull_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = "Valid Title",
            VideoUrl = "https://validurl.com"
        };

        Assert.Throws<ValidationException>(() => new VideoRequestValidator().ValidateAndThrow(request));
    }

    [Fact]
    public void Validate_WithDesEmpty_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = "Valid Title",
            Des = "",
            VideoUrl = "https://validurl.com"
        };

        Assert.Throws<ValidationException>(() => new VideoRequestValidator().ValidateAndThrow(request));
    }

    [Fact]
    public void Validate_WithDesLengthGreaterThan200_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = "Valid Title",
            Des = new string('a', 201),
            VideoUrl = "https://validurl.com"
        };

        Assert.Throws<ValidationException>(() => new VideoRequestValidator().ValidateAndThrow(request));
    }

    [Fact]
    public void Validate_WithVideoUrlNull_ThrowsArgumentException()
    {
        var request = new VideoRequest
        {
            Title = "Valid Title",
            Des = "Valid Description",
        };

        Assert.Throws<ValidationException>(() => new VideoRequestValidator().ValidateAndThrow(request));
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

        Assert.Throws<ValidationException>(() => new VideoRequestValidator().ValidateAndThrow(request));
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

        Assert.Throws<ValidationException>(() => new VideoRequestValidator().ValidateAndThrow(request));
    }

    [Fact]
    public void EnsurePageAndSize_WithValidParameters_DoesNotThrowException()
    {
        // Arrange
        long? page = 1;
        int? size = 10;

        // Act
        var ex = Record.Exception(() => Endpoints.EnsurePageAndSize(page, size));

        // Assert
        Assert.Null(ex);
    }

    [Fact]
    public void EnsurePageAndSize_WithNegativePage_ThrowsArgumentOutOfRangeException()
    {
        // Arrange
        long? page = -1;
        int? size = 10;

        // Act & Assert
        Assert.Throws<ArgumentOutOfRangeException>(() => Endpoints.EnsurePageAndSize(page, size));
    }

    [Fact]
    public void EnsurePageAndSize_WithNegativeSize_ThrowsArgumentOutOfRangeException()
    {
        // Arrange
        long? page = 1;
        int? size = -1;

        // Act & Assert
        Assert.Throws<ArgumentOutOfRangeException>(() => Endpoints.EnsurePageAndSize(page, size));
    }

    [Fact]
    public void EnsurePageAndSize_WithSizeGreaterThan20_ThrowsArgumentOutOfRangeException()
    {
        // Arrange
        long? page = 1;
        int? size = 21;

        // Act & Assert
        Assert.Throws<ArgumentOutOfRangeException>(() => Endpoints.EnsurePageAndSize(page, size));
    }

}