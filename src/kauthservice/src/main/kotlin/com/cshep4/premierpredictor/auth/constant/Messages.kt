package com.cshep4.premierpredictor.auth.constant

object Messages {
  const val TOKEN_VALIDATION_ERROR = "Token must be provided"

  const val LOGIN_VALIDATION_ERROR = "Email and password must be provided"
  const val LOGIN_INTERNAL_ERROR = "Somthing's gone wrong, please try again"
  const val LOGIN_INVALID_ERROR = "Invalid login details, please try again"

  const val REGISTER_VALIDATION_ERROR = "Invalid registration details, please try again"
  const val EMAIL_ALREADY_EXISTS_ERROR = "Email already exists, please try again"

  const val INITIATE_PASSWORD_RESET_VALIDATION_ERROR = "Email must be provided"

  const val PASSWORD_NOT_VALID_ERROR = "Could not reset password, password is not valid. Please try again"
  const val PASSWORDS_DO_NOT_MATCH_ERROR = "Could not reset password, password and confirmation don't match. Please try again"
  const val COULD_NOT_RESET_ERROR = "Could not reset password, Please try resending password reset email"
  const val SIGNATURE_EXPIRED_ERROR = "Could not reset password, reset link has expired. Please try resending password reset email"
}
