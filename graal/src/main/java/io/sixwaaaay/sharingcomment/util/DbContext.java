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

package io.sixwaaaay.sharingcomment.util;

/**
 * DbContext is a thread local variable which is used to determine
 * whether the current operation is a read or write operation.
 */
public class DbContext {
    /**
     * Thread local variable to store the current context.
     * default value is WRITE.
     */
    private static final ThreadLocal<DbContextEnum> CONTEXT = ThreadLocal.withInitial(() -> DbContextEnum.WRITE);


    /**
     * Set the current context.
     *
     * @param context the context to set.
     */
    public static void set(DbContextEnum context) {
        CONTEXT.set(context);
    }

    /**
     * Get the current context.
     *
     * @return the current context.
     */
    public static DbContextEnum get() {
        return CONTEXT.get();
    }
    
    /**
     * Clear the current context.
     */
    public static void clear() {
        CONTEXT.remove();
    }
}