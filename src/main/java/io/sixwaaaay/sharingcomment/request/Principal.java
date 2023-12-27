/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.request;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * The Principal class represents the principal user in the system.
 */
@Data
@AllArgsConstructor
@NoArgsConstructor
public class Principal {
    /**
     * The name of the principal user.
     */
    private String name;

    /**
     * The unique ID of the principal user.
     */
    private Long id;

    /**
     * The token of the principal user.
     */
    private String token;
}