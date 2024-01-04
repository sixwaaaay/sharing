/*
 * Copyright (c) 2023 sixwaaaay.
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
DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS follows;

CREATE TABLE users (
    `id` bigint,
    `username` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `phone` varchar(255) NOT NULL DEFAULT '',
    `avatar_url` varchar(255) NOT NULL DEFAULT '',
    `bio` varchar(255) NOT NULL DEFAULT '',
    `bg_url` varchar(255) NOT NULL DEFAULT '',
    `likes_given` int(11) NOT NULL DEFAULT 0,
    `likes_received` int(11) NOT NULL DEFAULT 0,
    `videos_posted` int(11) NOT NULL DEFAULT 0,
    `following` int(11) NOT NULL DEFAULT 0,
    `followers` int(11) NOT NULL DEFAULT 0,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE key `email` (`email`)) engine = InnoDB DEFAULT charset = utf8mb4;

CREATE TABLE follows (
    `user_id` bigint NOT NULL,
    `target` bigint NOT NULL,
    `id` bigint NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`, `target`),
    KEY `user_created` (`user_id`, `id` DESC),
    KEY `target_created` (target, `id` DESC)) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

INSERT INTO users.users (id, username, PASSWORD, email, phone, avatar_url, bio, bg_url, likes_given, likes_received, videos_posted, FOLLOWING, followers)
    VALUES (1, '1xxxxx', '$2a$10$5XFK9i5DJHgpnlYE8ukH0utCUKQCdmexjHuZASrMD43QGxejZsbPy', '1@x.com', '', '', '', '', 0, 0, 0, 0, 0),
    (2, '2xxxxx', '$2a$10$5XFK9i5DJHgpnlYE8ukH0utCUKQCdmexjHuZASrMD43QGxejZsbPy', '2@x.com', '', '', '', '', 0, 0, 0, 0, 0),
    (3, '3xxxxx', '$2a$10$5XFK9i5DJHgpnlYE8ukH0utCUKQCdmexjHuZASrMD43QGxejZsbPy', '3@x.com', '', '', '', '', 0, 0, 0, 0, 0),
    (4, '4xxxxx', '$2a$10$5XFK9i5DJHgpnlYE8ukH0utCUKQCdmexjHuZASrMD43QGxejZsbPy', '4@x.com', '', '', '', '', 0, 0, 0, 0, 0),
    (5, '5xxxxx', '$2a$10$5XFK9i5DJHgpnlYE8ukH0utCUKQCdmexjHuZASrMD43QGxejZsbPy', '5@x.com', '', '', '', '', 0, 0, 0, 0, 0),
    (6, '6xxxxx', '$2a$10$5XFK9i5DJHgpnlYE8ukH0utCUKQCdmexjHuZASrMD43QGxejZsbPy', '6@x.com', '', '', '', '', 0, 0, 0, 0, 0),
    (11, '3fasdfasdfdsa', '$2a$10$5XFK9i5DJHgpnlYE8ukH0utCUKQCdmexjHuZASrMD43QGxejZsbPy', 'fdas@x.com', '', '', '', '', 0, 0, 0, 0, 0),
    (493764627922944111, '123xxxxx', '$2a$10$.7mS7uGar.ngCuZRrMSw7OQ0sVKAuiVK0NxIg0Vq9aZy9o.QRU.tC', '123', '', '', '', '', 0, 0, 0, 0, 0),
    (494639229293297775, '11xxxxx', '$2a$10$FMVKo4b2ffFrS.ScdgkeMej7CCcflOMdlPpOSsOIN/cLJwQaJliWe', '11@x.com', '', '', '', '', 0, 0, 0, 0, 0),
    (494639239661617263, '22xxxxx', '$2a$10$t80pjNC1IdxzQxLmpPF4suQhsXjeW77sgBKUZx3WYUl8yKxW8Gn9.', '22@x.com', '', '', '', '', 0, 0, 0, 0, 0),
    (494639250013159535, '33xxxxx', '$2a$10$GwDLfYEwobO7EkbdCzm/Xe2SmyFiL4xz7//BigG6xj31Hp7cAZX0G', '33@x.com', '', '', '', '', 0, 0, 0, 0, 0),
    (494639258401767535, '44xxxxx', '$2a$10$oiqZFbgq.tbiEMwVSQ8Ob.9Kl3uJj88O3yh3.tNjOHvtzCcJyce0i', '44@x.com', '', '', '', '', 0, 0, 0, 0, 0),
    (494639267713122415, '55xxxxx', '$2a$10$3IZ1AnSCBYCddSzoR8Fg7u8ZRDgs/wcBz38opzFTeoD9aVYf0LFjW', '55@x.com', '', '', '', '', 0, 0, 0, 0, 0),
    (494639286319054959, '66xxxxx', '$2a$10$5XFK9i5DJHgpnlYE8ukH0utCUKQCdmexjHuZASrMD43QGxejZsbPy', '66@x.com', '', '', '', '', 0, 0, 0, 0, 0);

