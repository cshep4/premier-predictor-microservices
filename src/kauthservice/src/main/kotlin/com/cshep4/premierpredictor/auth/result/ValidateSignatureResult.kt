package com.cshep4.premierpredictor.auth.result

sealed class ValidateSignatureResult {
    object Success : ValidateSignatureResult()
    data class Error(val message: String, val cause: Exception? = null, val internal: Boolean = false) : ValidateSignatureResult()
}
