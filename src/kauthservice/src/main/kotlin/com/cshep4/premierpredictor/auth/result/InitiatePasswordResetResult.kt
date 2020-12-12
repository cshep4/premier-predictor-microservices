package com.cshep4.premierpredictor.auth.result

sealed class InitiatePasswordResetResult {
  companion object {
    val USER_NOT_FOUND_ERROR = Error(message = "user not found")
  }


  object Success : InitiatePasswordResetResult()
  data class Error(val message: String, val cause: Exception? = null) : InitiatePasswordResetResult()
}
