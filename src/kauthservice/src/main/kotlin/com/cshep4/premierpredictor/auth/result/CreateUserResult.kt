package com.cshep4.premierpredictor.auth.result

sealed class CreateUserResult {
    data class Success(val id: String) : CreateUserResult()
    data class Error(val message: String, val cause: Exception? = null, val internal: Boolean = false) : CreateUserResult()
}