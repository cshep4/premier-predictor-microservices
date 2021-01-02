package com.cshep4.premierpredictor.auth.result

import com.fasterxml.jackson.annotation.JsonAutoDetect
import com.fasterxml.jackson.annotation.JsonAutoDetect.Visibility.ANY
import com.fasterxml.jackson.annotation.JsonProperty

sealed class LoginResult {
    companion object {
        val USER_NOT_FOUND_ERROR = Error(message = "user not found")
        val PASSWORD_DOES_NOT_MATCH_ERROR = Error(message = "password does not match")
    }

    @JsonAutoDetect(fieldVisibility = ANY)
    data class Success(
            @JsonProperty("id") var id: String,
            @JsonProperty("token") var token: String
    ) : LoginResult()

    data class Error(val message: String, val cause: Exception? = null) : LoginResult()
}

