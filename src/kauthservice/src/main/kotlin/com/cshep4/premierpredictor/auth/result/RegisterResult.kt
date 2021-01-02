package com.cshep4.premierpredictor.auth.result

import com.fasterxml.jackson.annotation.JsonAutoDetect
import com.fasterxml.jackson.annotation.JsonAutoDetect.Visibility.ANY
import com.fasterxml.jackson.annotation.JsonProperty

sealed class RegisterResult {
    companion object {
        val EMAIL_ALREADY_EXISTS_ERROR = Error(message = "email already exists")
    }

    @JsonAutoDetect(fieldVisibility = ANY)
    data class Success(
            @JsonProperty("id")
            val id: String,
            @JsonProperty("token")
            val token: String
    ) : RegisterResult()

    data class Error(val message: String, val cause: Exception? = null) : RegisterResult()
}
