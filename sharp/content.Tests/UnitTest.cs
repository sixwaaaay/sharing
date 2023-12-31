using content.repository;
using JetBrains.Annotations;
using MySqlConnector;
using Xunit.Abstractions;

namespace content.Tests;

public class UnitTest(ITestOutputHelper testOutputHelper)
{
    [TestSubject(typeof(VideoRepository))]
    [Fact(DisplayName = "Video Repository")]
    public async void Test1()
    {
        var environmentVariable = Environment.GetEnvironmentVariable("CONNECTION_STRING") !;
        await using var dataSource = new MySqlDataSource(environmentVariable);
        var videoRepository = (IVideoRepository)new VideoRepository(dataSource);
        var video = new Video
        {
            Id = 1,
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
            Processed = 1
        };
        video = await videoRepository.Save(video);
        Assert.NotNull(video);
        Assert.Equal(1, video.UserId);

        var video2 = await videoRepository.FindById(1);
        Assert.NotNull(video2);
        Assert.Equal(1, video2.Id);
        Assert.Equal(1, video2.UserId);


        var videos = await videoRepository.FindAllByIds([1, 2, 3, 4, 5, 6]);
        Assert.NotNull(videos);
        var list = videos.ToList();
        foreach (var v in list)
        {
            Assert.NotNull(v);
            Assert.Equal(1, v.UserId);
        }

        var videos2 = await videoRepository.FindByUserId(1, long.MaxValue, 10);
        Assert.NotNull(videos2);
        Assert.NotEmpty(videos2);

        var videos3 = await videoRepository.FindRecent(long.MaxValue, 10);
        Assert.NotNull(videos3);
        Assert.NotEmpty(videos3);
    }

    [Fact(DisplayName = "simultaneously get videos")]
    public async Task Test3()
    {
        var now = DateTime.Now;
        var tasks = new List<Task>();
        await using var dataSource = new MySqlDataSource(Environment.GetEnvironmentVariable("CONNECTION_STRING") !);
        for (var i = 0; i < 10000; i++)
        {
            var videoRepository = new VideoRepository(dataSource);
            var task = Task.Run(async () =>
            {
                var videos = await videoRepository.FindRecent(long.MaxValue, 10);
                Assert.NotNull(videos);
                Assert.NotEmpty(videos);
            });
            tasks.Add(task);
        }
        await Task.WhenAll(tasks);
        testOutputHelper.WriteLine($"{DateTime.Now - now}");
    }
}