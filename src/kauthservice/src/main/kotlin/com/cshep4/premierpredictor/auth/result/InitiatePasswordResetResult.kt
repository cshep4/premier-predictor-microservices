package com.cshep4.premierpredictor.auth.result

sealed class InitiatePasswordResetResult {
  object Success : InitiatePasswordResetResult()
  data class Error(val message: String, val cause: Exception? = null) : InitiatePasswordResetResult()
}
