/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.request.error;

public class NoUserExitsError extends RuntimeException {
    public NoUserExitsError(String message) {
        super(message);
    }

    public static NoUserExitsError supply() {
        return new NoUserExitsError("no user exits");
    }
}
