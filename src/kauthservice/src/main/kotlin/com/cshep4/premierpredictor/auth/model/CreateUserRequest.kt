package com.cshep4.premierpredictor.auth.model

import com.fasterxml.jackson.annotation.JsonProperty

data class CreateUserRequest(
        @JsonProperty("firstName")
        val firstName: String,
        @JsonProperty("surname")
        val surname: String,
        @JsonProperty("email")
        val email: String,
        @JsonProperty("password")
        val password: String,
        @JsonProperty("predictedWinner")
        val predictedWinner: String
)
