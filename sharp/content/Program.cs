using System.Text.Json.Serialization;
using Microsoft.AspNetCore.Mvc;
using MySqlConnector;

var builder = WebApplication.CreateSlimBuilder(args);

builder.Services.ConfigureHttpJsonOptions(options =>
{
    options.SerializerOptions.TypeInfoResolverChain.Insert(0, AppJsonSerializerContext.Default);
});
builder.Services.AddProblemDetails();
builder.Services.AddMySqlDataSource(builder.Configuration.GetConnectionString("Default") ??
                                    throw new InvalidOperationException("Connection string is null"));

var app = builder.Build();
app.Run();

[JsonSourceGenerationOptions(PropertyNamingPolicy = JsonKnownNamingPolicy.SnakeCaseLower)]
[JsonSerializable(typeof(ProblemDetails))]
internal partial class AppJsonSerializerContext : JsonSerializerContext;