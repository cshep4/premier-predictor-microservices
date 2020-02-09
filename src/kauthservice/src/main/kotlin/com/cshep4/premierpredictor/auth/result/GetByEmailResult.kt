package com.cshep4.premierpredictor.auth.result

import com.cshep4.premierpredictor.auth.model.User

sealed class GetByEmailResult {
  data class Success(val user: User) : GetByEmailResult()
  data class Error(val message: String, val cause: Exception? = null) : GetByEmailResult()
}
