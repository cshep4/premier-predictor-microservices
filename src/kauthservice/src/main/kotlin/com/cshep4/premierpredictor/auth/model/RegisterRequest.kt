package com.cshep4.premierpredictor.auth.model

data class RegisterRequest(
  val firstName: String,
  val surname: String,
  val email: String,
  val password: String,
  val confirmation: String,
  val predictedWinner: String
)
