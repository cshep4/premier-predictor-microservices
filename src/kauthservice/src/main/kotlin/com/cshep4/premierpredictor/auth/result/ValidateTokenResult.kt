package com.cshep4.premierpredictor.auth.result

sealed class ValidateTokenResult {
    object Success : ValidateTokenResult()
    data class Error(val message: String, val cause: Exception? = null, val internal: Boolean = false) : ValidateTokenResult()
}
