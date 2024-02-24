/*
 * Copyright (c) 2024 sixwaaaay.
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


CREATE TABLE `notifications`
(
    `id`          BIGINT       NOT NULL AUTO_INCREMENT,
    `sender_id`   BIGINT       NOT NULL,
    `receiver_id` BIGINT       NOT NULL,
    `content`     varchar(255) NOT NULL,
    `type`        TINYINT unsigned NOT NULL DEFAULT 0,
    `status`      TINYINT unsigned NOT NULL DEFAULT 0,
    `created_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY           `query_by_receiver_id` (`status`,`receiver_id`,`id` ASC)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE videos
(
    id         BIGINT       NOT NULL AUTO_INCREMENT,
    user_id    BIGINT       NOT NULL,
    title      VARCHAR(255) NOT NULL,
    des        VARCHAR(255) NOT NULL,
    cover_url  VARCHAR(255) NOT NULL default '',
    video_url  VARCHAR(255) NOT NULL default '',
    duration   INT UNSIGNED,
    view_count INT UNSIGNED NOT NULL DEFAULT 0,
    like_count INT UNSIGNED NOT NULL DEFAULT 0,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    processed  TINYINT      NOT NULL DEFAULT 0,
    PRIMARY KEY (id),
    KEY        `user_created` (`processed`, `user_id`, `id` DESC),
    KEY        `processed` (`processed`, `id` DESC)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE counter
(
    id      BIGINT NOT NULL,
    counter INT UNSIGNED NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

INSERT INTO content.videos (id, user_id, title, des, cover_url, video_url, duration, view_count, like_count, created_at, updated_at, processed)
VALUES  (1, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 01:12:27', '2023-12-31 01:12:27', 1),
        (2, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 01:13:02', '2023-12-31 01:13:02', 1),
        (3, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 01:13:19', '2023-12-31 01:13:19', 1),
        (4, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 01:14:18', '2023-12-31 01:14:18', 1),
        (5, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 12:05:07', '2023-12-31 12:05:07', 1),
        (6, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 12:05:23', '2023-12-31 12:05:23', 1),
        (7, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 12:42:53', '2023-12-31 12:42:53', 1),
        (8, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 12:43:40', '2023-12-31 12:43:40', 1),
        (9, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 12:53:18', '2023-12-31 12:53:18', 1),
        (10, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 12:53:58', '2023-12-31 12:53:58', 1),
        (11, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 13:08:05', '2023-12-31 13:08:05', 1),
        (12, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 13:55:52', '2023-12-31 13:55:52', 1),
        (13, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 13:56:22', '2023-12-31 13:56:22', 1),
        (14, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 14:01:47', '2023-12-31 14:01:47', 1),
        (15, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 14:02:10', '2023-12-31 14:02:10', 1),
        (16, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 14:02:22', '2023-12-31 14:02:22', 1),
        (17, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 14:02:42', '2023-12-31 14:02:42', 1),
        (18, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 14:08:44', '2023-12-31 14:08:44', 1),
        (19, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 14:09:13', '2023-12-31 14:09:13', 1),
        (20, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 14:10:04', '2023-12-31 14:10:04', 1),
        (21, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 14:10:48', '2023-12-31 14:10:48', 1),
        (22, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 14:17:40', '2023-12-31 14:17:40', 1),
        (23, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 14:32:50', '2023-12-31 14:32:50', 1),
        (24, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 14:33:19', '2023-12-31 14:33:19', 1),
        (25, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 14:41:41', '2023-12-31 14:41:41', 1),
        (26, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 14:46:15', '2023-12-31 14:46:15', 1),
        (27, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 22:03:32', '2023-12-31 22:03:32', 1),
        (28, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 23:22:43', '2023-12-31 23:22:43', 1),
        (29, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 23:23:05', '2023-12-31 23:23:05', 1),
        (30, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2023-12-31 23:55:59', '2023-12-31 23:55:59', 1),
        (31, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2024-01-01 14:51:46', '2024-01-01 14:51:46', 1),
        (32, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2024-01-01 14:52:38', '2024-01-01 14:52:38', 1),
        (33, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2024-01-01 14:52:49', '2024-01-01 14:52:49', 1),
        (34, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2024-01-01 14:53:01', '2024-01-01 14:53:01', 1),
        (35, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2024-01-01 15:10:38', '2024-01-01 15:10:38', 1),
        (36, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2024-01-01 19:51:31', '2024-01-01 19:51:31', 1),
        (37, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2024-01-01 20:15:07', '2024-01-01 20:15:07', 1),
        (38, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2024-01-01 20:17:29', '2024-01-01 20:17:29', 1),
        (39, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2024-01-02 14:11:43', '2024-01-02 14:11:43', 1),
        (40, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2024-01-02 14:12:04', '2024-01-02 14:12:04', 1),
        (41, 1, 'title', 'des', 'coverUrl', 'videoUrl', 1, 1, 1, '2024-01-02 15:36:48', '2024-01-02 15:36:48', 1);

INSERT INTO content.counter (id, counter)
SELECT videos.user_id, count(*) FROM videos GROUP BY videos.user_id;

CREATE TABLE popular_videos
(
    order_num BIGINT NOT NULL
        PRIMARY KEY,
    id        BIGINT NULL,
    score     DOUBLE NULL
);



INSERT INTO popular_videos (id, order_num, score)
VALUES  ('16', 1, 90.077423201946),
        ('38', 2, 89.7837850170059),
        ('14', 3, 87.07611093956615),
        ('4', 4, 84.58457932303155),
        ('30', 5, 83.24402131442355),
        ('1', 6, 81.21239811294936),
        ('28', 7, 81.02940682417658),
        ('22', 8, 77.45784016089313),
        ('39', 9, 74.80318087600469),
        ('6', 10, 74.15545328907245),
        ('7', 11, 74.0317566078452),
        ('25', 12, 73.4048866419167),
        ('23', 13, 69.78511230068759),
        ('27', 14, 66.62686315050988),
        ('5', 15, 65.58183763696051),
        ('33', 16, 60.59655990507133),
        ('19', 17, 60.11935133498102),
        ('17', 18, 59.60366731472655),
        ('2', 19, 59.216842389867494),
        ('37', 20, 58.038582427463105),
        ('3', 21, 52.44702068385391),
        ('32', 22, 52.37177121673876),
        ('8', 23, 47.692426245373134),
        ('35', 24, 47.54776810067498),
        ('34', 25, 45.8674889485049),
        ('11', 26, 44.68464324687108),
        ('10', 27, 38.75704672071807),
        ('15', 28, 30.261150589455994),
        ('18', 29, 27.786070041159235),
        ('26', 30, 17.368304000579105),
        ('20', 31, 17.23854962479188),
        ('24', 32, 16.552044094122987),
        ('9', 33, 16.366864476694598),
        ('0', 34, 15.615050416081258),
        ('12', 35, 7.152079145565702),
        ('21', 36, 5.834697192380854),
        ('29', 37, 5.266447742717758),
        ('40', 38, 4.664552402370193),
        ('13', 39, 1.7064690605797517),
        ('31', 40, 0.42076641732898207),
        ('36', 41, 0.13637673122470895);