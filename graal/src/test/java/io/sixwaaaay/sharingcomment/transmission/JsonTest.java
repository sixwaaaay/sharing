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

package io.sixwaaaay.sharingcomment.transmission;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.sixwaaaay.sharingcomment.domain.Count;
import io.sixwaaaay.sharingcomment.domain.ReplyResult;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.*;


class JsonTest {
    private final ObjectMapper objectMapper = new ObjectMapper();
    @Test
    void testJson() {
        ScanVotedReq actualScanVotedReq = new ScanVotedReq();
        actualScanVotedReq.setLimit(1);
        actualScanVotedReq.setSubjectId(123L);
        actualScanVotedReq.setTargetType("Target Type");
        actualScanVotedReq.setToken(123L);
        actualScanVotedReq.setType("Type");
        assertEquals(1, actualScanVotedReq.getLimit().intValue());
        assertEquals(123L, actualScanVotedReq.getSubjectId().longValue());
        assertEquals("Target Type", actualScanVotedReq.getTargetType());
        assertEquals(123L, actualScanVotedReq.getToken().longValue());
        assertEquals("Type", actualScanVotedReq.getType());
        assertEquals(
                "ScanVotedReq(limit=1, subjectId=123, targetType=Target Type, token=123, type=Type)",
                actualScanVotedReq.toString());

        String json = "{\"limit\":1,\"token\":123,\"type\":\"Type\",\"subject_id\":123,\"target_type\":\"Target Type\"}";

        var jsonStr = objectMapper.valueToTree(actualScanVotedReq).toString();
        assertEquals(json, jsonStr);
    }

    @Test
    void test() {
        ScanVotedReply scanVotedReply = new ScanVotedReply();
        scanVotedReply.setTargetIds(List.of(1L, 1L, 1L));
        var json = "{\"target_ids\":[1,1,1]}";
        var jsonStr = objectMapper.valueToTree(scanVotedReply).toString();
        assertEquals(json, jsonStr);
    }

    @Test
    void testVoteReply() {
        VoteReply voteReply = new VoteReply();
        voteReply.setStatus("success");
        var json = "{\"status\":\"success\"}";
        var jsonStr = objectMapper.valueToTree(voteReply).toString();
        assertEquals(json, jsonStr);
    }

    @Test
    void testReplyResult() {
        var replyResult = new ReplyResult();
        replyResult.setNextPage(null);
        replyResult.setComments(List.of());

        var replyResult1 = new ReplyResult(null, List.of());
        var json = "{\"next_page\":null,\"comments\":[]}";
        var jsonStr = objectMapper.valueToTree(replyResult).toString();
        assertEquals(json, jsonStr);

        var jsonStr1 = objectMapper.valueToTree(replyResult1).toString();
        assertEquals(json, jsonStr1);
    }

    @Test
    void testCount() {
        var count = new Count();
        count.setId(1L);
        count.setCommentCount(1);
        var json = "{\"id\":1,\"commentCount\":1}";
        var jsonStr = objectMapper.valueToTree(count).toString();
        assertEquals(json, jsonStr);
    }
}