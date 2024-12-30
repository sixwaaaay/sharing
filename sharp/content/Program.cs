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

using System.Text;
using System.Text.Json;
using System.Text.Json.Serialization;
using content.domainservice;
using content.endpoints;
using content.repository;
using Microsoft.AspNetCore.Mvc;
using Microsoft.IdentityModel.Tokens;
using Npgsql;
using OpenTelemetry.Metrics;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;

var builder = WebApplication.CreateSlimBuilder(args);

builder.Services.ConfigureHttpJsonOptions(options =>
{
    options.SerializerOptions.TypeInfoResolverChain.Insert(0, AppJsonSerializerContext.Default);
    options.SerializerOptions.PropertyNamingPolicy = JsonNamingPolicy.SnakeCaseLower;
});

var secret = builder.Configuration.GetSection("Secret").Value ?? throw new InvalidOperationException("Secret is null");

builder.Services.AddOpenApi();
builder.Services.AddAuthentication("Bearer").AddJwtBearer(
    option =>
    {
        option.TokenValidationParameters.IssuerSigningKey = new SymmetricSecurityKey
            (Encoding.UTF8.GetBytes(secret));
        option.TokenValidationParameters.ValidateAudience = false;
        option.TokenValidationParameters.ValidateIssuer = false;
    }
);

const string serviceName = "sharing.content";
var otelEndpoint = builder.Configuration.GetConnectionString("Otel_GrpcEndpoint");
builder.Services.AddOpenTelemetry().WithTracing(tcb =>
{
    tcb
        .AddSource(serviceName).AddNpgsql()
        .SetResourceBuilder(ResourceBuilder.CreateDefault().AddService(serviceName: serviceName))
        .AddHttpClientInstrumentation()
        .AddAspNetCoreInstrumentation()
        .AddOtlpExporter(o =>
        {
            if (!string.IsNullOrEmpty(otelEndpoint))
            {
                o.Endpoint = new Uri(otelEndpoint);
            }
        });
}).WithMetrics(mtb =>
{
    mtb
        .AddMeter("Npgsql")
        .AddAspNetCoreInstrumentation()
        .AddHttpClientInstrumentation()
        .AddPrometheusExporter();
});

builder.Services.AddSingleton(TracerProvider.Default.GetTracer(serviceName));
builder.Services.AddSingleton<MessageRequestValidator>().AddSingleton<VideoRequestValidator>();

builder.Services.AddAuthorization().AddProbe();
builder.Services.AddProblemDetails().AddResponseCompression();

builder.Services.AddNpgsqlDataSource(builder.Configuration.GetConnectionString("Default").EnsureNotNull("Connection string is null"));

builder.Services.AddVideoRepository().AddNotificationRepository();

builder.Services.AddUserRepository();

builder.Services.AddVoteRepository();

builder.Services.AddSearchClient();

builder.Services.AddDomainService().AddMessageDomain();

var app = builder.Build();

app.UseExceptionHandler();
app.UseStatusCodePages();
app.UseResponseCompression();
app.UseAuthorization();

app.UseToken();

app.MapPrometheusScrapingEndpoint();
app.MapEndpoints();

app.Run();

[JsonSourceGenerationOptions(PropertyNamingPolicy = JsonKnownNamingPolicy.SnakeCaseLower)]
[JsonSerializable(typeof(ProblemDetails))]
[JsonSerializable(typeof(VideoRequest))]
[JsonSerializable(typeof(VoteRequest))]
[JsonSerializable(typeof(VideoDto))]
[JsonSerializable(typeof(Pagination<VideoDto>))]
[JsonSerializable(typeof(IReadOnlyList<VideoDto>))]
[JsonSerializable(typeof(Pagination<MessageDto>))]
[JsonSerializable(typeof(MessageDto))]
[JsonSerializable(typeof(MessageRequest))]
[JsonSerializable(typeof(long?))]
internal partial class AppJsonSerializerContext : JsonSerializerContext;