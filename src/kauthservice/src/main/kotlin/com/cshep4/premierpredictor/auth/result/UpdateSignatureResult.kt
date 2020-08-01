package com.cshep4.premierpredictor.auth.result

sealed class UpdateSignatureResult {
    object Success : UpdateSignatureResult()
    data class Error(val message: String, val cause: Exception? = null, val internal: Boolean = false) : UpdateSignatureResult()
}
