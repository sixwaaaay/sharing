using content.domainservice;
using content.repository;
using Xunit.Abstractions;

namespace content.Tests.repository;

public class MapperTest
(ITestOutputHelper testOutputHelper) {
  [Fact(DisplayName = "Video Mapper")]
  public void Test2() {
    var video = new Video { Id = 1,
                            UserId = 1,
                            Title = "title",
                            Des = "des",
                            CoverUrl = "coverUrl",
                            VideoUrl = "videoUrl",
                            Duration = 1,
                            Category = "category",
                            Tags = "tags",
                            ViewCount = 1,
                            LikeCount = 1,
                            CreatedAt = DateTime.Now,
                            UpdatedAt = DateTime.Now,
                            Processed = 1 };
    var videoDto = video.ToDto();
    Assert.NotNull(videoDto);
    testOutputHelper.WriteLine(videoDto.ToString());
  }
}
