package com.cshep4.premierpredictor.auth.result

sealed class UpdatePasswordResult {
    object Success : UpdatePasswordResult()
    data class Error(val message: String, val cause: Exception? = null, val internal: Boolean = false) : UpdatePasswordResult()
}
