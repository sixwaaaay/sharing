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

package io.sixwaaaay.sharingcomment.domain;


import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.data.annotation.Id;
import org.springframework.data.relational.core.mapping.Column;
import org.springframework.data.relational.core.mapping.Table;

/**
 * This class represents the count of all comments for a given object.
 * It is used to store the count in the database.
 */
@Table("counts")
@Data
@NoArgsConstructor
@AllArgsConstructor
public class Count {

    /**
     * The unique identifier for the count.
     */
    @Id
    private long id;

    /**
     * The count value.
     */
    @Column("comment_count")
    private int commentCount;
}