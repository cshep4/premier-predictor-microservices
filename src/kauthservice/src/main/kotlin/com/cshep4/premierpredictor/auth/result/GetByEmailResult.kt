package com.cshep4.premierpredictor.auth.result

import com.cshep4.premierpredictor.auth.model.User

sealed class GetByEmailResult {
  companion object {
    val USER_NOT_FOUND_ERROR = Error(message = "user not found")
  }

  data class Success(val user: User) : GetByEmailResult()
  data class Error(val message: String, val cause: Exception? = null, val internal: Boolean = false) : GetByEmailResult()
}
