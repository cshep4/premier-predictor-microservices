@file:JvmName("ResetPasswordRequestValidationResultKt")

package com.cshep4.premierpredictor.auth.result

sealed class UpdateUserResult {
  object Success : UpdateUserResult()
  data class Error(val message: String, val cause: Exception? = null) : UpdateUserResult()
}

val USER_NOT_FOUND_ERROR = UpdateUserResult.Error(message = "user not found")
