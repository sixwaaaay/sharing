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
using content.repository;
using Grpc.Core;
using Grpc.Net.Client;
using JetBrains.Annotations;

namespace content.Tests.repository;

[TestSubject(typeof(UserRepository))]
public class UserRepositoryTest
{
    
    private readonly ChannelBase _channel = GrpcChannel.ForAddress(Environment.GetEnvironmentVariable("USER_STRING") !);

    [Fact(DisplayName = "FindUserById")]
    public async void Test1()
    {
        var client = new UserRepository(_channel);
        var result = await client.FindById(1);
        
        Assert.NotNull(result);
        Assert.Equal("1", result.Id);
    }
    
    [Fact(DisplayName = "FindUserByIds")]
    public async void Test2()
    {

        var client = new UserRepository(_channel);
        var result = await client.FindAllByIds(new []{1L, 2L});
        Assert.NotNull(result);
        Assert.Equal(2, result.Count);
    }
    
    [Fact(DisplayName = "invalid token, error")]
    public async void Test3()
    {
        var client = new UserRepository(_channel)
        {
            Token = "token"
        };
        await Assert.ThrowsAsync<RpcException>(async () => await client.FindAllByIds(new []{493764627922944111, 11}));
    }
}