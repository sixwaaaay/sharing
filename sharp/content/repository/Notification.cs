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

using Dapper;
using MySqlConnector;

namespace content.repository;

public record Message
{
    public long Id { get; set; }
    public long SenderId { get; init; }
    public long ReceiverId { get; init; }
    public string Content { get; init; } = string.Empty;
    public short Type { get; init; }
    public short Status { get; init; }
    public DateTime CreatedAt { get; init; } = DateTime.Now;
}

public interface INotificationRepository
{
    ValueTask<IReadOnlyList<Message>> FindByReceiverId(long receiverId, long page, int size, bool unreadOnly);

    ValueTask<Message> Save(Message notification);
    ValueTask UpdateStatus(long userId, long id, short status);
}

public class NotificationRepository(MySqlDataSource dataSource) : INotificationRepository
{
    public async ValueTask<IReadOnlyList<Message>> FindByReceiverId(long receiverId, long page, int size,
        bool unreadOnly = false)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        if (unreadOnly)
        {
            var messageNotifications = await connection.QueryAsync<Message>(
                "SELECT id, sender_id, receiver_id, content, type, status, created_at " +
                "FROM notifications WHERE receiver_id = @receiverId AND status = 0 " +
                "AND id > @page limit @size", new { receiverId, page, size });
            return messageNotifications.ToList();
        }

        var notifications = await connection.QueryAsync<Message>(
            "SELECT id, sender_id, receiver_id, content, type, status, created_at " +
            "FROM notifications WHERE receiver_id = @receiverId " +
            "AND id > @page limit @size",
            new { receiverId, page, size });
        return notifications.ToList();
    }

    public async ValueTask<Message> Save(Message notification)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        var result = await connection.ExecuteAsync(
            "INSERT INTO notifications (sender_id, receiver_id, content, type, status, created_at) " +
            "VALUES (@SenderId, @ReceiverId, @Content, @Type, @Status, @CreatedAt)", notification);
        if (result == 0)
        {
            throw new Exception("Insert notification failed");
        }

        notification.Id = await connection.QuerySingleAsync<long>("SELECT LAST_INSERT_ID()");
        return notification;
    }

    public async ValueTask UpdateStatus(long userId, long id, short status)
    {
        await using var connection = await dataSource.OpenConnectionAsync();
        await connection.ExecuteAsync(
            "UPDATE notifications SET status = @status WHERE id = @id AND receiver_id = @userId",
            new { userId, id, status });
    }
}

public static class NotificationRepositoryExtension
{
    public static IServiceCollection AddNotificationRepository(this IServiceCollection services) =>
        services.AddSingleton<INotificationRepository, NotificationRepository>();
}