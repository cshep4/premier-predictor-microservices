package com.cshep4.premierpredictor.auth.model

data class ResetPasswordRequest(
        val email: String,
        val signature: String,
        val password: String,
        val confirmation: String = ""
)
