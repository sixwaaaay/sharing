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

CREATE TABLE `comments`
(
    `id`          bigint       NOT NULL AUTO_INCREMENT,
    `user_id`     bigint       NOT NULL,
    `content`     varchar(255) NOT NULL,
    `reply_to`    bigint,
    `belong_to`   bigint       NOT NULL,
    `created_at`  datetime     NOT NULL DEFAULT (CURRENT_TIMESTAMP(3)),
    `reply_count` int          NOT NULL DEFAULT 0,
    `like_count`  int          NOT NULL DEFAULT 0,
    `refer_to`    bigint,
    PRIMARY KEY (`id`),
    INDEX `finder` (`belong_to`, `reply_to`, `id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_bin;


INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '2', null, 1, '2023-11-25 14:50:09', 3, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '12', null, 1, '2023-11-25 14:50:26', 3, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '1', null, 1, '2023-11-25 14:51:21', 3, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '2', null, 1, '2023-11-25 14:51:21', 3, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '3', null, 1, '2023-11-25 14:51:21', 3, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '4', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '5', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '6', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '7', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '8', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '9', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '10', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '11', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '12', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '13', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '14', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '15', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '16', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '17', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '18', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '19', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '20', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '21', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '22', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '23', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '24', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '25', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '26', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '27', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (1, '28', null, 1, '2023-11-25 14:51:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy1', 1, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply2', 1, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy2', 1, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply3', 1, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy3', 1, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply4', 2, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy4', 2, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply5', 2, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy5', 2, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply6', 2, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy6', 2, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply7', 3, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy7', 3, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply8', 3, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy8', 3, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply9', 3, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy9', 3, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply10', 4, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy10', 4, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply11', 4, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy2', 4, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply3', 4, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy3', 4, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply4', 5, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy4', 5, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply5', 5, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'relpy5', 5, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (5, 'reply6', 5, 1, '2023-11-25 15:09:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', 62, 1, '2023-11-26 15:43:50', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', 62, 1, '2023-11-26 15:45:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', 62, 1, '2023-11-26 15:45:45', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', 62, 1, '2023-11-26 15:46:22', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', 62, 1, '2023-11-26 15:46:25', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', 62, 1, '2023-11-26 15:46:27', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', 62, 1, '2023-11-26 15:46:29', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', 62, 1, '2023-11-26 15:46:32', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', 62, 1, '2023-11-26 15:55:16', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', 62, 1, '2023-11-26 15:58:54', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', 62, 1, '2023-11-26 15:59:21', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 05:32:42', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 05:37:52', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 05:45:52', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 05:47:07', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 05:49:20', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 06:20:28', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 06:20:34', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 06:23:18', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 06:26:28', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 06:27:44', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 06:29:11', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 06:44:55', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 07:40:10', 0, 0);
INSERT INTO comments (user_id, content, reply_to, belong_to, created_at, reply_count, like_count)
VALUES (2, 'hello world', null, 1, '2023-11-27 07:40:48', 0, 0);
