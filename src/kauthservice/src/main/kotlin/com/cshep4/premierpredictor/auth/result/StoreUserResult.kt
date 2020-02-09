package com.cshep4.premierpredictor.auth.result

sealed class StoreUserResult {
  data class Success(val id: String) : StoreUserResult()
  data class Error(val message: String, val cause: Exception? = null) : StoreUserResult()
}
