package com.cshep4.premierpredictor.auth.result

sealed class ResetPasswordRequestValidationResult {
  object Success : ResetPasswordRequestValidationResult()
  data class Error(val message: String) : ResetPasswordRequestValidationResult()
}
