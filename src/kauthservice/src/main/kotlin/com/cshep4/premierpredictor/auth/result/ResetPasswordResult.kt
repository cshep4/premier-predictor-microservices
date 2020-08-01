package com.cshep4.premierpredictor.auth.result

sealed class ResetPasswordResult {
    companion object {
        val INVALID_SIGNATURE_ERROR = Error(message = "invalid signature")
        val SIGNATURE_DOES_NOT_MATCH_ERROR = Error(message = "signature does not match")
    }

    object Success : ResetPasswordResult()
    data class Error(val message: String, val cause: Exception? = null, val internal: Boolean = false) : ResetPasswordResult()
}
