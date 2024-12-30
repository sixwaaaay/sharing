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

import com.fasterxml.jackson.annotation.JsonProperty;

public class ShardEnum {
    /*
    the highest bit is not used,
    the second-highest bit (62) to the 60th bit is used to store the shard ID.
    in total, 3 bits, so there are at most 8 shards.
    */
    public static final long Chore = 0L; // aka 0b000L << 60;
    public static final long SHARD_ID_0 = 0b001L << 60;
    public static final long SHARD_ID_1 = 0b010L << 60;
    public static final long SHARD_ID_2 = 0b011L << 60;
    public static final long SHARD_ID_3 = 0b100L << 60;

    public enum Shard {
        @JsonProperty("chore")
        chore(Chore),
        @JsonProperty("default")
        SHARD(SHARD_ID_0),
        @JsonProperty("video")
        video(SHARD_ID_1),
        @JsonProperty("post")
        post(SHARD_ID_2),
        @JsonProperty("music")
        music(SHARD_ID_3);

        private final long shardId;

        Shard(long shardId) {
            this.shardId = shardId;
        }

        /**
         * transform id to shard id
         *
         * @param id id
         * @return the id embed shard id
         */
        public long transform(long id) {
            return this.shardId | id;
        }
    }
}
