package com.cshep4.premierpredictor.auth.model

import com.fasterxml.jackson.annotation.JsonProperty

data class User(
        @JsonProperty("id")
        val id: String,
        @JsonProperty("firstName")
        var firstName: String,
        @JsonProperty("surname")
        var surname: String,
        @JsonProperty("predictedWinner")
        var predictedWinner: String,
        @JsonProperty("score")
        var score: Int,
        @JsonProperty("email")
        var email: String,
        @JsonProperty("password")
        var password: String,
        @JsonProperty("signature")
        val signature: String? = null
)
