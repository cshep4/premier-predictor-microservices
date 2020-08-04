package com.cshep4.premierpredictor.auth.result

sealed class LoginResult {
    companion object {
        val USER_NOT_FOUND_ERROR = Error(message = "user not found")
        val PASSWORD_DOES_NOT_MATCH_ERROR = Error(message = "password does not match")
    }

    data class Success(val id: String, val token: String) : LoginResult()
    data class Error(val message: String, val cause: Exception? = null) : LoginResult()
}

