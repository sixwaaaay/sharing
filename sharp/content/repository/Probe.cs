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
using System.Diagnostics;

namespace content.repository;

public interface IProbe
{
    /// <summary>
    ///  get the resolution of video from given url
    /// </summary>
    /// <param name="url">
    ///  video url
    /// </param>
    /// <returns>
    ///  resolution of video
    /// </returns>
    public Task<string> GetVideoResolution(string url);

    /// <summary>
    /// get the duration of video from given url
    /// </summary>
    /// <param name="url">
    /// video url
    /// </param>
    /// <returns>
    /// duration of video
    /// </returns>
    public Task<string> GetVideoDuration(string url);
}

/// <summary>
///  use ffprobe to get video info
/// </summary>
/// <param name="executablePath">
/// ffprobe executable path
/// </param>
public class Probe(string executablePath) : IProbe
{
    public async Task<string> GetVideoResolution(string url)
    {
        EnsureUrlExists(url);
        if (string.IsNullOrEmpty(executablePath))
        {
            return string.Empty;
        }

        return await Process($"-v error -select_streams v:0 -show_entries stream=width,height -of csv=s=x:p=0 {url}");
    }

    private async Task<string> Process(string arguments)
    {
        var process = new Process();
        process.StartInfo.FileName = executablePath;
        process.StartInfo.Arguments = arguments;
        process.StartInfo.UseShellExecute = false;
        process.StartInfo.RedirectStandardOutput = true;
        process.Start();
        var output = await process.StandardOutput.ReadToEndAsync();
        if (string.IsNullOrEmpty(output))
        {
            throw new Exception("not a video");
        }

        await process.WaitForExitAsync();
        return output;
    }


    public async Task<string> GetVideoDuration(string url)
    {
        EnsureUrlExists(url);

        if (string.IsNullOrWhiteSpace(executablePath))
        {
            return string.Empty;
        }

        return await Process($"-v error -show_entries format=duration -of default=noprint_wrappers=1:nokey=1 {url}");
    }

    private static void EnsureUrlExists(string url)
    {
        if (string.IsNullOrEmpty(url))
        {
            throw new ArgumentNullException(nameof(url));
        }
    }
}

public static class ProbeExtensions
{
    public static IServiceCollection AddProbe(this IServiceCollection services) =>
        services.AddSingleton<IProbe>(new Probe(Environment.GetEnvironmentVariable("FFPROBE_PATH") ?? string.Empty));
}