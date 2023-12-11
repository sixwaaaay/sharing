/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
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