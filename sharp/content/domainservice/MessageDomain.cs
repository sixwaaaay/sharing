using content.repository;
using Riok.Mapperly.Abstractions;

namespace content.domainservice;

public interface IMessageDomain
{
    Task<Pagination<MessageDto>> FindMessages(long receiverId, long page, int size, bool unreadOnly);
    Task<MessageDto> Save(Message message);
    Task MarkAsRead(long id, long userId);
}

public record MessageDto
{
    public long Id { get; init; }
    public long SenderId { get; init; }
    public long ReceiverId { get; init; }
    public string Content { get; init; } = string.Empty;
    public short Type { get; init; }
    public bool Read { get; set; }
    public DateTime CreatedAt { get; init; } = DateTime.Now;
}

[Mapper]
public static partial class MessageMapper
{
    public static partial MessageDto ToDto(this Message video);
}

public class MessageDomain(INotificationRepository notification) : IMessageDomain
{
    public async Task<Pagination<MessageDto>> FindMessages(long receiverId, long page, int size,
        bool unreadOnly = false)
    {
        var messages = await notification.FindByReceiverId(receiverId, page, size, unreadOnly);
        var messageDtos = messages.Select(m =>
        {
            var dto = m.ToDto();
            dto.Read = m.Status == 1;
            return dto;
        }).ToList();

        return new Pagination<MessageDto>
        {
            AllCount = messageDtos.Count,
            Items = messageDtos,
            NextPage = messageDtos.Count == size ? messageDtos[^1].Id.ToString() : null
        };
    }

    public async Task<MessageDto> Save(Message message)
    {
        message = await notification.Save(message);
        return message.ToDto();
    }

    public async Task MarkAsRead(long id, long userId) => await notification.UpdateStatus(userId, id, 1);
}

public static class MessageDomainExtensions
{
    public static IServiceCollection AddMessageDomain(this IServiceCollection services) =>
        services.AddSingleton<IMessageDomain, MessageDomain>();
}