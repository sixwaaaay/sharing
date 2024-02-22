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
 */

package io.sixwaaaay.sharingcomment;

import io.sixwaaaay.sharingcomment.client.UserClient;
import io.sixwaaaay.sharingcomment.tools.JwtUtil;
import io.sixwaaaay.sharingcomment.util.TokenParser;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;

import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

@SpringBootTest
@AutoConfigureMockMvc
class SharingCommentApplicationTests {


    @Autowired
    private UserClient userClient;
    @Autowired
    private MockMvc mockMvc;

    @Autowired
    private JwtUtil jwtUtil;


    @Test
    public void getMainCommentListTest() throws Exception {
        var token = jwtUtil.generateToken("john", "111");
        mockMvc.perform(MockMvcRequestBuilders.get("/comments/main")
                        .param("belong_to", "1")
                        .param("size", "10")
                        .header("Authorization", "Bearer " + token))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON));
        // without token
        mockMvc.perform(MockMvcRequestBuilders.get("/comments/main")
                        .param("belong_to", "1")
                        .param("size", "10"))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON));
    }

    @Test
    public void getReplyCommentListTest() throws Exception {
        var token = jwtUtil.generateToken("n", "1111111");
        mockMvc.perform(MockMvcRequestBuilders.get("/comments/reply")
                        .param("belong_to", "1")
                        .param("reply_to", "1")
                        .param("size", "10")
                        .header("Authorization", "Bearer " + token))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON));

        mockMvc.perform(MockMvcRequestBuilders.get("/comments/reply")
                        .param("belong_to", "1")
                        .param("reply_to", "1")
                        .param("size", "10"))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON));
    }


    @Test
    public void createCommentTest() throws Exception {
        var token = jwtUtil.generateToken("n", "1111111");
        var json = "{ \"content\": \"This is a test comment\", \"reply_to\": 1, \"belong_to\": 1 }";
        mockMvc.perform(MockMvcRequestBuilders.post("/comments")
                        .content(json)
                        .contentType(MediaType.APPLICATION_JSON)
                        .header("Authorization", "Bearer " + token))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON));

        mockMvc.perform(MockMvcRequestBuilders.post("/comments")
                        .content(json)
                        .contentType(MediaType.APPLICATION_PROBLEM_JSON))
                .andExpect(MockMvcResultMatchers.status().is4xxClientError());
    }

    @Test
    public void deleteComments() throws Exception {
        var token = jwtUtil.generateToken("n", "1");
        var json = "{ \"content\": \"This is a test comment\", \"reply_to\": null, \"belong_to\": 1 }";
        mockMvc.perform(MockMvcRequestBuilders.delete("/comments/21")
                        .content(json)
                        .contentType(MediaType.APPLICATION_JSON)
                        .header("Authorization", "Bearer " + token))
                .andExpect(MockMvcResultMatchers.status().isOk());

        mockMvc.perform(MockMvcRequestBuilders.delete("/comments/1")
                        .content(json)
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(MockMvcResultMatchers.status().is4xxClientError());
    }

    @Test
    public void voteComment() throws Exception {
        var token = jwtUtil.generateToken("n", "1");
        mockMvc.perform(MockMvcRequestBuilders.post("/comments/action/like/1")
                        .contentType(MediaType.APPLICATION_JSON)
                        .header("Authorization", "Bearer " + token))
                .andExpect(MockMvcResultMatchers.status().isOk());

        mockMvc.perform(MockMvcRequestBuilders.delete("/comments/action/like/1")
                        .contentType(MediaType.APPLICATION_JSON)
                        .header("Authorization", "Bearer " + token))
                .andExpect(MockMvcResultMatchers.status().isOk());

        mockMvc.perform(MockMvcRequestBuilders.delete("/comments/action/like/1")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(MockMvcResultMatchers.status().is4xxClientError());

    }


    @Test
    void testGetUser() {
        var token = jwtUtil.generateToken("john", "111");

        var userReply = userClient.getUser(1L, token);
        Assertions.assertNotNull(userReply);
    }

    @Test
    void testGetUsers() {
        var token = jwtUtil.generateToken("john", "111");
        var usersReply = userClient.getManyUser(List.of(457232417502052951L, 457121784278309633L), token);
        Assertions.assertNotNull(usersReply);
    }

    @Autowired
    private TokenParser tokenParser;

    @Test
    public void testParseWithValidToken() {
        var l = 12314214521512131L;
        var id = String.valueOf(l);
        var johnDoe = "John Doe";
        var validToken = jwtUtil.generateToken(johnDoe, id);
        var principal = tokenParser.parse("Bearer " + validToken);
        assertTrue(principal.isPresent());
        assertEquals(johnDoe, principal.get().getName());
        assertEquals(l, principal.get().getId());
    }

    @Test
    public void testParseWithInvalidToken() {
        var invalidToken = "Invalid token";
        var principal = tokenParser.parse(invalidToken);
        assertFalse(principal.isPresent());
    }

    @Test
    public void testParseWithEmptyToken() {
        var emptyToken = "";
        var principal = tokenParser.parse(emptyToken);
        assertFalse(principal.isPresent());
    }
}
