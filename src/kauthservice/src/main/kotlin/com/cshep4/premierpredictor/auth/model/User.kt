package com.cshep4.premierpredictor.auth.model

data class User(
        val id: String,
        var firstName: String,
        var surname: String,
        var predictedWinner: String,
        var score: Int,
        var email: String,
        var password: String,
        val signature: String? = null
)
