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

package io.sixwaaaay.sharingcomment.request;

import io.sixwaaaay.sharingcomment.request.error.NoUserExitsError;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.security.authentication.AnonymousAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;

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

    /**
     *  The current token of the principal user.
     * @return the current token of the principal user.
     */
    public static String currentToken(){
        if (SecurityContextHolder.getContext().getAuthentication().getPrincipal() instanceof Principal principal) {
            return principal.getToken();
        } else {
            return null;
        }
    }

    /**
     * The current user ID of the principal user.
     * @return the current user ID of the principal user.
     */
    public static long currentUserId(){
        if (SecurityContextHolder.getContext().getAuthentication().getPrincipal() instanceof Principal principal) {
            return principal.getId();
        } else if (SecurityContextHolder.getContext().getAuthentication() instanceof AnonymousAuthenticationToken) {
            return 0;
        }
        throw NoUserExitsError.supply();
    }
}