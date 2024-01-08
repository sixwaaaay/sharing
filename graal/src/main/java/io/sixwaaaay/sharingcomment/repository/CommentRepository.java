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

package io.sixwaaaay.sharingcomment.repository;


import io.sixwaaaay.sharingcomment.domain.Comment;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.data.domain.Limit;
import org.springframework.data.jdbc.repository.query.Modifying;
import org.springframework.data.jdbc.repository.query.Query;
import org.springframework.data.repository.CrudRepository;

import java.util.List;

public interface CommentRepository extends CrudRepository<Comment, Long> {

    /**
     * select ... from comment where
     * belong_to = ? and id < ? and reply_to is null
     * order by id desc
     * limit ?
     *
     * @param belongTo the id of the object which the comments belong to
     * @param id       the id of the earliest comment in the previous page
     * @return page of comments
     */
    @Cacheable("comments-main")
    List<Comment> findByBelongToAndIdLessThanAndReplyToNullOrderByIdDesc(Long belongTo, Long id, Limit limit);


    /**
     * select ... from comment where
     * belong_to = ?
     * reply_to = ?
     * order by id asc
     * limit ?
     *
     * @param replyTo the id of the comment which the comments belong to
     * @param id      the id of the latest comment in the previous page
     * @param limit   the limit
     * @return the list of comments
     */
    @Cacheable("comments-reply")
    List<Comment> findByBelongToAndReplyToAndIdGreaterThanOrderByIdAsc(Long belongTo, Long replyTo, Long id, Limit limit);


    /**
     * delete from comments where id = ? and user_id = ?
     *
     * @param id     the id of the comment
     * @param userId the id of the user
     */
    @Modifying
    @Query("delete from `comments` where `id` = :id and `user_id` = :userId")
    boolean deleteByIdAndUserId(Long id, Long userId);

    /**
     * update the reply_count of the specified comment
     *
     * @param id the id of the comment
     */
    @Modifying
    @Query("update `comments` set `reply_count` = `reply_count` + 1 where id = :id")
    void increaseReplyCount(Long id);

    /**
     * update the reply_count of the specified comment
     *
     * @param id the id of the comment
     */
    @Modifying
    @Query("update `comments` set `reply_count` = `reply_count` - 1 where id = :id")
    void decreaseReplyCount(Long id);


    /**
     * update the like_count of the specified comment
     *
     * @param id the id of the comment
     */
    @Modifying
    @Query("update `comments` set `like_count` = `like_count` + 1 where id = :id")
    void increaseLikeCount(Long id);

    /**
     * update the like_count of the specified comment
     *
     * @param id the id of the comment
     */
    @Modifying
    @Query("update `comments` set `like_count` = `like_count` - 1 where id = :id")
    void decreaseLikeCount(Long id);
}
