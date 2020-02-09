package com.cshep4.premierpredictor.auth.result

sealed class LoginResult {
  data class Success(val id: String, val token: String) : LoginResult()
  data class Error(val message: String, val cause: Exception? = null) : LoginResult()
}
