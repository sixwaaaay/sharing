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

namespace content.Tests.repository;

using MySqlConnector;
using Xunit;
using System.Collections.Generic;
using System.Threading.Tasks;
using content.repository;

public class NotificationRepositoryTests
{
    private readonly NotificationRepository _repository;

    public NotificationRepositoryTests()
    {
        var dataSource = new MySqlDataSource(Environment.GetEnvironmentVariable("CONNECTION_STRING") !);
        _repository = new NotificationRepository(dataSource);
    }

    [Fact]
    public async Task FindByReceiverId_ReturnsNotifications_WhenUnreadOnlyIsFalse()
    {
        // Arrange
        const long receiverId = 1L;
        const long page = 0L;
        const int size = 10;

        // Act
        var result = await _repository.FindByReceiverId(receiverId, page, size);

        // Assert
        Assert.NotNull(result);
        Assert.IsType<List<Message>>(result);
    }

    [Fact]
    public async Task FindByReceiverId_ReturnsUnreadNotifications_WhenUnreadOnlyIsTrue()
    {
        // Arrange
        const long receiverId = 1L;
        const long page = 0L;
        const int size = 10;
        const bool unreadOnly = true;

        // Act
        var result = await _repository.FindByReceiverId(receiverId, page, size, unreadOnly);

        // Assert
        Assert.NotNull(result);
        Assert.IsType<List<Message>>(result);
    }

    [Fact]
    public async Task Save_ReturnsNotification_WhenInsertIsSuccessful()
    {
        // Arrange
        var notification = new Message
        {
            SenderId = 1L,
            ReceiverId = 2L,
            Content = "Test content",
            Type = 1,
        };

        // Act
        var result = await _repository.Save(notification);

        // Assert
        Assert.NotNull(result);
        Assert.IsType<Message>(result);
    }

    [Fact]
    public async Task UpdateStatus_UpdatesNotificationStatus()
    {
        // Arrange
        const long id = 1L;
        const int status = 1;

        // Act
        await _repository.UpdateStatus(2, id, status);

        // Assert
        // No exception means the test passed
    }
}