/*
 * Copyright (c) 2024 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

-- test data

create database compute;
\c compute

CREATE TABLE video_events
(
    video_id   BIGINT,
    event_type INT,
    event_time TIMESTAMP
);

DO
$do$
    DECLARE
        i INT;
    BEGIN
        FOR i IN 1..100000
            LOOP
                INSERT INTO video_events (video_id, event_type, event_time)
                VALUES ((random() * 5000)::int,
                        (random() * 3 + 1)::int,
                        timestamp '2024-01-01 00:00:00' +
                        ((random() * 365)::int || ' days')::interval);
            END LOOP;
    END
$do$;