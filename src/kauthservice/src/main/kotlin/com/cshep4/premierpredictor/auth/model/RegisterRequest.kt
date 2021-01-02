package com.cshep4.premierpredictor.auth.model

import com.fasterxml.jackson.annotation.JsonProperty

data class RegisterRequest(
        @JsonProperty("firstName")
        val firstName: String,
        @JsonProperty("surname")
        val surname: String,
        @JsonProperty("email")
        val email: String,
        @JsonProperty("password")
        val password: String,
        @JsonProperty("confirmation")
        val confirmation: String = "",
        @JsonProperty("predictedWinner")
        val predictedWinner: String
)
