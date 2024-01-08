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

import io.lettuce.core.dynamic.annotation.Param;
import io.sixwaaaay.sharingcomment.domain.Count;
import org.springframework.data.jdbc.repository.query.Modifying;
import org.springframework.data.jdbc.repository.query.Query;
import org.springframework.data.repository.CrudRepository;

/**
 * This interface represents the repository for the Count entity.
 * It extends CrudRepository to provide CRUD operations for the Count entity.
 * It includes methods to increase, decrease, and create a count.
 */
public interface CountRepository extends CrudRepository<Count, Long> {

    /**
     * Increases the comment count by 1 for the count with the provided id.
     *
     * @param id The id of the count to be increased.
     * @return true if the count was successfully increased, false otherwise.
     */
    @Modifying
    @Query("UPDATE `counts` SET `comment_count` = `comment_count` + 1 WHERE id = :id")
    boolean increaseCount(@Param("id") long id);

    /**
     * Decreases the comment count by 1 for the count with the provided id.
     *
     * @param id The id of the count to be decreased.
     */
    @Modifying
    @Query("UPDATE `counts` SET `comment_count` = `comment_count` - 1 WHERE id = :id")
    void decreaseCount(@Param("id") long id);

    /**
     * Creates a new count with the provided id and a comment count of 1.
     *
     * @param id The id of the count to be created.
     */
    @Modifying
    @Query("INSERT INTO `counts` (`id`, `comment_count`) VALUES (:id, 1)")
    void createCount(@Param("id") long id);
}