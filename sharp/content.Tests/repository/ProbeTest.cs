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
using content.repository;
using JetBrains.Annotations;
using Xunit.Abstractions;

namespace content.Tests.repository;

[TestSubject(typeof(Probe))]
public class ProbeTest(ITestOutputHelper testOutputHelper)
{
    private readonly string _executablePath = Environment.GetEnvironmentVariable("FFPROBE_PATH") !;

    [Fact]
    public async void Method()
    {
        // Arrange
        var peg = new Probe(_executablePath);
        // Act
        var result =
            await peg.GetVideoDuration(@"https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.webm");
        // Assert
        Assert.NotNull(result);
        testOutputHelper.WriteLine(result);
    }

    [Fact(DisplayName = "Empty executable path")]
    public async void EmptyExecutablePath()
    {
        // Arrange
        var peg = new Probe(string.Empty);
        // Act
        var result =
            await peg.GetVideoDuration(@"https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.webm");
        
        var result2 =
            await peg.GetVideoResolution(@"https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.webm");
        // Assert
        Assert.Equal(string.Empty, result);
        Assert.Equal(string.Empty, result2);
    }

    [Fact(DisplayName = "empty url")]
    public async void EmptyUrl()
    {
        // Arrange
        var peg = new Probe(_executablePath);
        // Act && Assert
        await Assert.ThrowsAsync<ArgumentNullException>(async () =>
            await peg.GetVideoDuration(string.Empty));
    }

    [Fact(DisplayName = "resolution")]
    public async void Resolution()
    {
        // Arrange
        var peg = new Probe(_executablePath);
        // Act
        var result =
            await peg.GetVideoResolution(@"https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.webm");
        // Assert
        Assert.NotNull(result);
        testOutputHelper.WriteLine(result);
    }
    
    [Fact(DisplayName = "not a video")]
    public async void NotAVideo()
    {
        // Arrange
        var peg = new Probe(_executablePath);
        // Act && Assert
        await Assert.ThrowsAsync<Exception>(async () =>
            await peg.GetVideoResolution(@"https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.txt"));
    }
}