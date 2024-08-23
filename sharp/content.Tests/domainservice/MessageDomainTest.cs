namespace content.Tests.domainservice;

using Xunit;
using Moq;
using System.Collections.Generic;
using content.domainservice;
using content.repository;

public class MessageDomainTests
{
    private readonly Mock<INotificationRepository> _mockNotificationRepository;
    private readonly IMessageDomain _messageDomain;

    public MessageDomainTests()
    {
        _mockNotificationRepository = new Mock<INotificationRepository>();
        _messageDomain = new MessageDomain(_mockNotificationRepository.Object);
    }

    [Fact]
    public async Task FindMessages_ReturnsCorrectPagination_WhenMessagesExist()
    {
        // Arrange
        var messages = new List<Message>
        {
            new() { Id = 1, ReceiverId = 1, Status = 1 },
            new() { Id = 2, ReceiverId = 1, Status = 0 }
        };
        _mockNotificationRepository.Setup(repo => repo.FindByReceiverId(1, 1, 2, false)).ReturnsAsync(messages);

        // Act
        var result = await _messageDomain.FindMessages(1, 1, 2, false);

        // Assert
        Assert.Equal(2, result.AllCount);
        Assert.Equal(2, result.Items.Count);
        Assert.Equal("2", result.NextPage);
    }

    [Fact]
    public async Task Save_ReturnsCorrectDto_WhenMessageIsSaved()
    {
        // Arrange
        var message = new Message { Id = 1, ReceiverId = 1, Status = 1 };
        _mockNotificationRepository.Setup(repo => repo.Save(message)).ReturnsAsync(message);

        // Act
        var result = await _messageDomain.Save(message);

        // Assert
        Assert.Equal(message.Id, result.Id);
        Assert.Equal(message.ReceiverId, result.ReceiverId);
    }

    [Fact]
    public async Task MarkAsRead_UpdatesStatus_WhenCalled()
    {
        // Arrange
        _mockNotificationRepository.Setup(repo => repo.UpdateStatus(1, 1, 1)).Returns(ValueTask.CompletedTask);

        // Act
        await _messageDomain.MarkAsRead(1, 1);

        // Assert
        _mockNotificationRepository.Verify(repo => repo.UpdateStatus(1, 1, 1), Times.Once);
    }
}