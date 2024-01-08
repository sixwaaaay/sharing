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
using content.repository;
using JetBrains.Annotations;
using Moq;

namespace content.Tests.repository;

[TestSubject(typeof(VoteRepository))]
public class VoteRepositoryTest
{
    private readonly HttpClient _httpClient = new HttpClient()
    {
        BaseAddress = new Uri(Environment.GetEnvironmentVariable("VOTE_STRING") !)
    };

    [Fact]
    public async Task UpdateVote_ShouldPostVote_WhenVoteTypeIsVote()
    {
        // Arrange
        var voteRepository = new VoteRepository(_httpClient)
        {
            CurrentUser = 3L
        };

        // Act
        await voteRepository.UpdateVote(1L, VoteType.Vote);
    }

    [Fact]
    public async Task UpdateVote_ShouldPostCancelVote_WhenVoteTypeIsCancelVote()
    {
        // Arrange
        var voteRepository = new VoteRepository(_httpClient)
        {
            CurrentUser = 3L
        };

        // Act
        await voteRepository.UpdateVote(1L, VoteType.CancelVote);
    }

    [Fact]
    public async Task UpdateVote_ShouldThrowException_WhenVoteTypeIsInvalid()
    {
        // Arrange
        var httpClient = new Mock<HttpClient>();
        var voteRepository = new VoteRepository(httpClient.Object);
        voteRepository.CurrentUser = 1;
        var videoId = 1L;
        var voteType = (VoteType)100; // Invalid vote type
        // Act & Assert
        await Assert.ThrowsAsync<ArgumentOutOfRangeException>(() => voteRepository.UpdateVote(videoId, voteType));
    }

    [Fact]
    public async Task VotedOfVideos_ShouldReturnEmptyList_WhenNoVideoIdsProvided()
    {
        // Arrange
        var httpClient = new Mock<HttpClient>();
        var voteRepository = new VoteRepository(httpClient.Object);
        voteRepository.CurrentUser = 1;

        // Act
        var result = await voteRepository.VotedOfVideos(new long[0]);

        // Assert
        Assert.Empty(result);
    }

    [Fact]
    public async Task VotedOfVideos_ShouldReturnEmptyList_WhenCurrentUserIsZero()
    {
        // Arrange
        var voteRepository = new VoteRepository(_httpClient)
        {
            CurrentUser = 2,
        };

        // Act
        var result = await voteRepository.VotedOfVideos(new long[] { 12345, 23456, 4567 });

        // Assert
        Assert.NotNull(result);
        Assert.Empty(result);

        result = await voteRepository.VotedOfVideos([]);

        // Assert
        Assert.NotNull(result);
        Assert.Empty(result);

        voteRepository.CurrentUser = 0;

        // Act 
        result = await voteRepository.VotedOfVideos([12345, 23456, 4567]);
        // Assert
        Assert.NotNull(result);
        Assert.Empty(result);
    }

    [Fact]
    public async Task VotedVideos()
    {
        // Arrange
        var voteRepository = new VoteRepository(_httpClient)
        {
            CurrentUser = 2,
        };
        await voteRepository.UpdateVote(5555L, VoteType.Vote);

        // Act
        var result = await voteRepository.VotedVideos(2, 0, 10);

        // Assert

        Assert.NotNull(result);
        Assert.NotEmpty(result);
    }
}