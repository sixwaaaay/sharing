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