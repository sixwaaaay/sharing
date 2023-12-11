/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.api;

import io.sixwaaaay.sharingcomment.tools.JwtUtil;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;

@SpringBootTest
@AutoConfigureMockMvc
public class ApiTest {

    @Autowired
    private MockMvc mockMvc;

    @Autowired
    private JwtUtil jwtUtil;

    @Test
    public void getMainCommentListTest() throws Exception {
        var token = jwtUtil.generateToken("john", "111");
        mockMvc.perform(MockMvcRequestBuilders.get("/comments/main")
                        .param("belong_to", "1")
                        .param("page", "1")
                        .param("size", "10")
                        .header("Authorization", "Bearer " + token))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON));
        // without token
        mockMvc.perform(MockMvcRequestBuilders.get("/comments/main")
                        .param("belong_to", "1")
                        .param("page", "1")
                        .param("size", "10"))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON));
    }

    @Test
    public void getReplyCommentListTest() throws Exception {
        var token = jwtUtil.generateToken("n", "1111111");
        mockMvc.perform(MockMvcRequestBuilders.get("/comments/reply")
                        .param("reply_to", "1")
                        .param("page", "1")
                        .param("size", "10")
                        .header("Authorization", "Bearer " + token))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON));

        mockMvc.perform(MockMvcRequestBuilders.get("/comments/reply")
                        .param("reply_to", "1")
                        .param("page", "1")
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
                        .header("Authorization", "Bearer "+token))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON));

        mockMvc.perform(MockMvcRequestBuilders.post("/comments")
                        .content(json)
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(MockMvcResultMatchers.status().is4xxClientError());
    }

}


