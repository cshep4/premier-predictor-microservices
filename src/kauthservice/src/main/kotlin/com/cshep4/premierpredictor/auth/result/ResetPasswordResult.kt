package com.cshep4.premierpredictor.auth.result

sealed class ResetPasswordResult {
    companion object {
        val USER_NOT_FOUND_ERROR = Error(message = "user not found")
        val SIGNATURE_DOES_NOT_MATCH_ERROR = Error(message = "signature does not match")
    }

    object Success : ResetPasswordResult()
    data class Error(val message: String, val cause: Exception? = null) : ResetPasswordResult()
}
