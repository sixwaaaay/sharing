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

package io.sixwaaaay.sharingcomment.util;

import lombok.Getter;

public class ShardEnum {
    // 最高位不使用， 次高位(62位) 到第60位(1位) 用于存储分片ID
    // 共3位，八个分片
    public static final long Chore = 0L; // aka 0b000L << 60;
    public static final long SHARD_ID_0 = 0b001L << 60;
    public static final long SHARD_ID_1 = 0b010L << 60;
    public static final long SHARD_ID_2 = 0b011L << 60;
    public static final long SHARD_ID_3 = 0b100L << 60;

    public static Shard getShard(String shard) {
        return switch (shard) {
            case "chore" -> Shard.ChoreShard;
            case "default" -> Shard.SHARD_0;
            case "video" -> Shard.Video;
            case "post" -> Shard.post;
            case "music" -> Shard.music;
            default -> throw new IllegalArgumentException("Unknown shard: " + shard);
        };
    }


    @Getter
    public enum Shard {
        ChoreShard(Chore),
        SHARD_0(SHARD_ID_0),
        Video(SHARD_ID_1),
        post(SHARD_ID_2),
        music(SHARD_ID_3);

        private final long shardId;

        Shard(long shardId) {
            this.shardId = shardId;
        }
    }

    /**
     * transform id to shard id
     *
     * @param id   id
     * @param shard shard
     * @return the id embed shard id
     */
    public static long transformId(long id, Shard shard) {
        return shard.getShardId() | id;
    }
}
