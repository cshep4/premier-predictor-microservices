package com.cshep4.premierpredictor.auth.result

sealed class ResetPasswordResult {
  object Success : ResetPasswordResult()
  data class Error(val message: String, val cause: Exception? = null) : ResetPasswordResult()
}
