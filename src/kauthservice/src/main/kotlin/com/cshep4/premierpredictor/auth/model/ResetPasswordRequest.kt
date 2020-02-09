package com.cshep4.premierpredictor.auth.model

data class ResetPasswordArgs(
  val email: String,
  val signature: String,
  val password: String
)
