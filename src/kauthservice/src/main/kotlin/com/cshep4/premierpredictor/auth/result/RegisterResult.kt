package com.cshep4.premierpredictor.auth.result

sealed class RegisterResult {
    companion object {
        val EMAIL_ALREADY_EXISTS_ERROR = Error(message = "email already exists")
    }

    data class Success(val id: String, val token: String) : RegisterResult()
    data class Error(val message: String, val cause: Exception? = null, val internal: Boolean = false) : RegisterResult()
}
