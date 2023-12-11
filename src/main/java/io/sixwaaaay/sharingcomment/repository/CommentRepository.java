/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
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
    List<Comment> findByReplyToAndIdGreaterThanOrderByIdAsc(Long replyTo, Long id, Limit limit);


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
}
