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

drop table if exists users;
drop table if exists follows;
create table users
(
    `id`                bigint,
    `username`          varchar(255) not null,
    `password`          varchar(255) not null,
    `email`             varchar(255) not null,
    `phone`             varchar(255) not null default '',
    `avatar_url`        varchar(255) not null default '',
    `bio`               varchar(255) not null default '',
    `bg_url`            varchar(255) not null default '',
    `likes_given`       int(11)      not null default 0,
    `likes_received`    int(11)      not null default 0,
    `videos_posted`     int(11)      not null default 0,
    `following`         int(11)      not null default 0,
    `followers`         int(11)      not null default 0,
    `registration_time` datetime              default current_timestamp,
    PRIMARY KEY (`id`),
    unique key `email` (`email`)
) engine = InnoDB
  default charset = utf8mb4;

create table follows
(
    `user_id`    bigint NOT NULL,
    `target`     bigint NOT NULL,
    `id`         bigint NOT NULL,
    `created_at` datetime   NOT NULL,
    PRIMARY KEY (`user_id`, `target`),
    KEY `user_created` (`user_id`, `id` DESC),
    KEY `target_created` (target, `id` DESC)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
