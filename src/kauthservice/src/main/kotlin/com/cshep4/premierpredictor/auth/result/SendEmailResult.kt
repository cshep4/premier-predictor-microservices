package com.cshep4.premierpredictor.auth.result

sealed class SendEmailResult {
  object Success : SendEmailResult()
  data class Error(val message: String, val cause: Exception? = null) : SendEmailResult()
}
