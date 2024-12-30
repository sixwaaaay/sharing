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

using System.Net;
using System.Net.Http.Json;
using System.Text.Json.Serialization;
using content.repository;
using Moq;
using Moq.Protected;

namespace content.Tests.repository;

public class UserRepositoryTest
{
    [Fact(DisplayName = "FindUserById")]
    public async Task Test1()
    {
        // Arange
        var id = 1L;
        var user = new User { Id = id.ToString(), Name = "test", AvatarUrl = "https://example.com/avatar.png", BgUrl = "https://example.com/bg.png", Bio = "test user", };
        var mockHttpMessageHandler = new Mock<HttpMessageHandler>();
        mockHttpMessageHandler.Protected()
            .Setup<Task<HttpResponseMessage>>("SendAsync", ItExpr.IsAny<HttpRequestMessage>(), ItExpr.IsAny<CancellationToken>())
            .ReturnsAsync(new HttpResponseMessage { StatusCode = HttpStatusCode.OK, Content = JsonContent.Create(user, UserJsonContext.Default.User), });
        var httpClient = new HttpClient(mockHttpMessageHandler.Object) { BaseAddress = new Uri("http://localhost:5151") };
        var client = new UserRepository(httpClient);
        // Act
        var result = await client.FindById(id);

        // Assert
        Assert.NotNull(result);
        Assert.Equal(id.ToString(), result.Id);
        Assert.Equal("test", result.Name);
        Assert.Equal("https://example.com/avatar.png", result.AvatarUrl);
        Assert.Equal("https://example.com/bg.png", result.BgUrl);
        Assert.Equal("test user", result.Bio);
    }

    [Fact(DisplayName = "FindUserByIds")]
    public async Task Test2()
    {
        // Arange
        var ids = new List<long> { 1L, 2L };
        var users = new List<User>
        {
            new() {Id = "1",Name = "test1",AvatarUrl = "https://example.com/avatar1.png",BgUrl = "https://example.com/bg1.png",Bio = "test user1"},
            new(){Id = "2",Name = "test2",AvatarUrl = "https://example.com/avatar2.png",BgUrl = "https://example.com/bg2.png",Bio = "test user2"}
        };
        var mockHttpMessageHandler = new Mock<HttpMessageHandler>();
        mockHttpMessageHandler.Protected().Setup<Task<HttpResponseMessage>>("SendAsync", ItExpr.IsAny<HttpRequestMessage>(), ItExpr.IsAny<CancellationToken>())
            .ReturnsAsync(new HttpResponseMessage { StatusCode = HttpStatusCode.OK, Content = JsonContent.Create(users, UserJsonContext.Default.IReadOnlyListUser) });
        var httpClient = new HttpClient(mockHttpMessageHandler.Object) { BaseAddress = new Uri("http://localhost:5151") };
        var client = new UserRepository(httpClient);
        // Act
        var result = await client.FindAllByIds(ids);
        // Assert
        Assert.NotNull(result);
        Assert.Equal(2, result.Count);
        Assert.Equal("1", result[0].Id);
        Assert.Equal("https://example.com/avatar1.png", result[0].AvatarUrl);
        Assert.Equal("test user1", result[0].Bio);
        Assert.Equal("2", result[1].Id);
        Assert.Equal("https://example.com/bg2.png", result[1].BgUrl);
    }
}

[JsonSourceGenerationOptions(PropertyNamingPolicy = JsonKnownNamingPolicy.SnakeCaseLower)]
[JsonSerializable(typeof(IReadOnlyList<User>))]
partial class UserJsonContext : JsonSerializerContext;