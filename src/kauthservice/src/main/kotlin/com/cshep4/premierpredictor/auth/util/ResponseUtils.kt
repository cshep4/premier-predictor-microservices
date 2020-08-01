package com.cshep4.premierpredictor.auth.util

import javax.ws.rs.core.Response
import javax.ws.rs.core.Response.Status.*

object ResponseUtils {
    fun badRequest(message: String? = null): Response = Response
            .status(BAD_REQUEST)
            .error(message)
            .build()

    fun conflict(message: String? = null): Response = Response
            .status(CONFLICT)
            .error(message)
            .build()

    fun ok(body: Any? = null): Response = Response
            .ok(body)
            .build()

    fun created(body: Any? = null): Response = Response
            .status(CREATED)
            .entity(body)
            .build()

    fun internal(message: String? = null): Response = Response
            .status(INTERNAL_SERVER_ERROR)
            .error(message)
            .build()

    fun unauthorized(message: String? = null): Response = Response
            .status(UNAUTHORIZED)
            .error(message)
            .build()

    fun Response.ResponseBuilder.error(message: String?): Response.ResponseBuilder {
        message ?: return this
        return this.entity(Error(message))
    }

    data class Error(val message: String)
}