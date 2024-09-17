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

CREATE TABLE my_videos
(
    id      BIGINT,
    user_id BIGINT,
    PRIMARY KEY (id)
) FROM mysql_content TABLE 'content.videos';

/* create a materialized view to aggregate the number of videos per user */
CREATE MATERIALIZED VIEW user_video_count AS
SELECT user_id, COUNT(*) AS video_count
FROM my_videos
GROUP BY user_id;


CREATE TABLE my_follows
(
    user_id bigint,
    target  bigint,
    id      bigint,
    PRIMARY KEY (user_id, target)
) FROM mysql_users TABLE 'users.follows';


/*
 aggregate the number of followers per user
*/

CREATE MATERIALIZED VIEW user_follower_count AS
SELECT target AS user_id, COUNT(*) AS followers
FROM my_follows
GROUP BY target;

/*
    aggregate the number of followings per user
*/

CREATE MATERIALIZED VIEW user_following_count AS
SELECT user_id, COUNT(*) AS following
FROM my_follows
GROUP BY user_id;


CREATE TABLE my_comments
(
    id        bigint,
    user_id   bigint,
    belong_to bigint,
    PRIMARY KEY (id)
) FROM mysql_comments TABLE 'comments.comments';

/* aggregate the number of comments per belong_to entity */
CREATE MATERIALIZED VIEW belong_to_comment_count AS
SELECT belong_to, COUNT(*) AS comment_count
FROM my_comments
GROUP BY belong_to;




/*
CREATE TABLE graph
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    relation   VARCHAR(255) NOT NULL,
    subject_id BIGINT       NOT NULL,
    object_id  BIGINT       NOT NULL,
    UNIQUE KEY id (relation, subject_id, object_id),
    KEY subject_id_scan (relation, subject_id, id DESC)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
*/

CREATE TABLE my_graph
(
    relation  VARCHAR(255),
    subject_id BIGINT,
    object_id BIGINT,
    PRIMARY KEY (relation, subject_id, object_id)
) FROM mysql_vote TABLE 'vote.graph';



/*
case: relation = 'vote:video' means one user add a vote to another video
case: relation = 'user:comment' means one user vote a comment
*/

/*
    aggregate the number of votes per video
*/

CREATE MATERIALIZED VIEW video_vote_count AS
SELECT object_id AS video_id, COUNT(*) AS vote_count
FROM my_graph
WHERE relation = 'vote:video'
GROUP BY object_id;

/*
    aggregate the number of votes to videos that a user has received

    note:
    'vote:video', subject_id (user_id) , object_id (video_id)
    record the vote count for each video
    however, we want to know the sum of video votes that a user who published the video has received
    so we join the video_vote_count with my_videos to get the user_id
*/

CREATE MATERIALIZED VIEW user_video_vote_count AS
SELECT my_videos.user_id, SUM(video_vote_count.vote_count) AS vote_received
FROM video_vote_count
INNER JOIN my_videos
    ON video_vote_count.video_id = my_videos.id
GROUP BY my_videos.user_id;

/*
aggregate the number of videos that a user has voted
*/
CREATE MATERIALIZED VIEW user_vote_count AS
SELECT subject_id AS user_id, COUNT(*) AS vote_count
FROM my_graph
WHERE relation = 'vote:video'
GROUP BY subject_id;

/*
aggregate the number of votes that a comment has received
*/
CREATE MATERIALIZED VIEW comment_vote_count AS
SELECT object_id AS comment_id, COUNT(*) AS vote_count
FROM my_graph
WHERE relation = 'user:comment'
GROUP BY object_id;

/*
    the number of videos, followers, followings for each user
*/
CREATE MATERIALIZED VIEW user_stats AS
SELECT
    user_video_count.user_id as user_id,
    video_count,
    followers,
    following
FROM user_video_count
INNER JOIN user_follower_count
    ON user_video_count.user_id = user_follower_count.user_id
INNER JOIN user_following_count
    ON user_video_count.user_id = user_following_count.user_id;

/*
    the number of comments, votes for each video
*/
CREATE MATERIALIZED VIEW video_stats AS
SELECT
    video_id,
    comment_count,
    vote_count
FROM
    belong_to_comment_count
INNER JOIN video_vote_count
    ON belong_to_comment_count.belong_to = video_vote_count.video_id;

