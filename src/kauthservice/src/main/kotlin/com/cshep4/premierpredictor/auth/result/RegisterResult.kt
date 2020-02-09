package com.cshep4.premierpredictor.auth.result

sealed class RegisterResult {
  data class Success(val id: String, val token: String) : RegisterResult()
  data class Error(val message: String, val cause: Exception? = null) : RegisterResult()
}

val EMAIL_EXISTS_ERROR = RegisterResult.Error(message = "email already exists")
