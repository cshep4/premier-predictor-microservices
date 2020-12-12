package com.cshep4.premierpredictor.auth.model

data class CreateUserRequest(
        val firstName: String,
        val surname: String,
        val email: String,
        val password: String,
        val predictedWinner: String
)
