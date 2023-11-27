/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.controller;

import io.sixwaaaay.sharingcomment.request.error.NoUserExitsError;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.*;

@ControllerAdvice
@ResponseBody
public class ExceptionController {

    @ResponseStatus(code = HttpStatus.BAD_REQUEST, reason = "malformed parameter")
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public void handleValidationExceptions(MethodArgumentNotValidException ex) {
    }

    @ResponseStatus(code = HttpStatus.UNAUTHORIZED)
    @ExceptionHandler(NoUserExitsError.class)
    public void handleNoUserExitsError(NoUserExitsError ex) {
    }
}
