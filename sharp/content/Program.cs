using System.Text;
using System.Text.Json;
using System.Text.Json.Serialization;
using content.domainservice;
using content.endpoints;
using content.repository;
using Microsoft.AspNetCore.Mvc;
using Microsoft.IdentityModel.Tokens;
using MySqlConnector;

var builder = WebApplication.CreateSlimBuilder(args);

builder.Services.ConfigureHttpJsonOptions(options =>
{
    options.SerializerOptions.TypeInfoResolverChain.Insert(0, AppJsonSerializerContext.Default);
    options.SerializerOptions.PropertyNamingPolicy = JsonNamingPolicy.SnakeCaseLower;
});

var secret = builder.Configuration.GetSection("Secret").Value?? throw new InvalidOperationException("Secret is null");

builder.Services.AddAuthentication("Bearer").AddJwtBearer(
    option =>
    {
        option.TokenValidationParameters.IssuerSigningKey = new SymmetricSecurityKey
            (Encoding.UTF8.GetBytes(secret));
        option.TokenValidationParameters.ValidateAudience = false;
        option.TokenValidationParameters.ValidateIssuer = false;
    }
);

builder.Services.AddAuthorization();

builder.Services.AddProblemDetails();

builder.Services.AddMySqlDataSource(builder.Configuration.GetConnectionString("Default") ??
                                    throw new InvalidOperationException("Connection string is null"));
builder.Services.AddVideoRepository();

builder.Services.AddGrpcUser();
builder.Services.AddUserRepository();

builder.Services.AddVoteClient();
builder.Services.AddVoteRepository();

builder.Services.AddDomainService();

builder.Services.AddResponseCompression();
var app = builder.Build();

app.UseResponseCompression();
app.UseAuthorization();

app.UseToken();

app.MapEndpoints();

app.Run();

[JsonSourceGenerationOptions(PropertyNamingPolicy = JsonKnownNamingPolicy.SnakeCaseLower)]
[JsonSerializable(typeof(ProblemDetails))]
[JsonSerializable(typeof(VideoRequest))]
[JsonSerializable(typeof(VoteRequest))]
[JsonSerializable(typeof(VideoDto))]
[JsonSerializable(typeof(Pagination<VideoDto>))]
[JsonSerializable(typeof(IReadOnlyList<VideoDto>))]
internal partial class AppJsonSerializerContext : JsonSerializerContext;