package com.cshep4.premierpredictor.auth.result

sealed class HashResult {
    data class Success(val hash: String) : HashResult()
    data class Error(val message: String, val cause: Exception? = null, val internal: Boolean = true) : HashResult()
}
