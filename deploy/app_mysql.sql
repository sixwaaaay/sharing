-- MySQL
DROP DATABASE IF EXISTS `test`;

CREATE DATABASE `test`;

use `test`;

-- 用户表

CREATE TABLE `users` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT, -- 用户ID

    `username` varchar(32) NOT NULL, -- 用户名

    `password` varchar(255) NOT NULL, -- 密码

    `salt` char(6) NOT NULL, -- 密码盐

    `followed_count` int NOT NULL, -- 关注数

    `follower_count` int NOT NULL -- 粉丝数

);

-- 用户名唯一

ALTER TABLE users
    ADD UNIQUE (username);

--

CREATE TABLE `videos` (
    `id` bigint PRIMARY KEY AUTO_INCREMENT, -- 视频ID

    `user_id` bigint NOT NULL, -- 用户ID

    `play_url` varchar(255) NOT NULL, -- 播放地址

    `cover_url` varchar(255) NOT NULL, -- 封面地址

    `favorite_count` int NOT NULL, -- 点赞数

    `comment_count` int NOT NULL, -- 评论数

    `title` varchar(255) NOT NULL, -- 标题

    `created_at` timestamp NOT NULL -- 创建时间

    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 查询用户投稿的视频，按时间排列

CREATE INDEX `videos_user_id_created_at` ON `videos` (`user_id`, `created_at`);

-- 按时间排序

CREATE INDEX `videos_created_at` ON `videos` (`created_at`);

--

CREATE TABLE `favorites` (
    `id` bigint PRIMARY KEY AUTO_INCREMENT, -- 点赞记录ID

    `user_id` bigint NOT NULL, -- 用户ID

    `video_id` bigint NOT NULL, -- 视频ID

    `action` int NOT NULL, -- 操作类型，0：未点赞，1：点赞

    `update_at` timestamp NOT NULL -- 更新时间

    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 通过用户点赞记录查询视频，按时间排序

CREATE INDEX `favorites_index_0` ON `favorites` (`user_id`, `update_at`);

-- 对视频点赞，需要保证(user_id, video_id)唯一

ALTER TABLE `favorites`
    ADD CONSTRAINT `favorites_unique_0` UNIQUE (`user_id`, `video_id`);

--

CREATE TABLE `comments` (
    `id` bigint PRIMARY KEY AUTO_INCREMENT, -- 评论ID

    `user_id` bigint NOT NULL, -- 用户ID

    `video_id` bigint NOT NULL, -- 视频ID

    `content` varchar(255) NOT NULL, -- 评论内容

    `created_at` timestamp NOT NULL -- 创建时间

    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 查询视频的评论按时间排序

CREATE INDEX comments_index_0 ON comments (video_id, created_at);

-- 查询用户评论的视频 （暂时没有需求
-- 查询评论视频的用户 （暂时没有需求
-- 查询用户的评论的视频 （暂时没有需求
--

CREATE TABLE relations -- 指示关系边 follower -> followed

(
    `id` bigint PRIMARY KEY AUTO_INCREMENT, -- 关系ID

    `follower` bigint NOT NULL, -- 粉丝ID

    `followed` bigint NOT NULL, -- 关注ID

    `status` int NOT NULL, -- 关系状态，0：未关注，1：已关注

    `update_at` timestamp NOT NULL -- 更新时间;

    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 查询用户的关注，按时间排序

CREATE INDEX relation_index_0 ON relations (follower, update_at);

-- 查询用户的粉丝, 按时间排序

CREATE INDEX relation_index_1 ON relations (followed, update_at);

--
--    - 修改关注状态，添加关注或者取消关注，那么需要 (follower, followed) 唯一 (隐式索引)
--    - 查询指定 (follower, followed) 的状态

ALTER TABLE relations
    ADD CONSTRAINT relation_unique_0 UNIQUE (follower, followed);

