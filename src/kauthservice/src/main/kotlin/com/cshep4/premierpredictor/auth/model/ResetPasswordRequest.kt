package com.cshep4.premierpredictor.auth.model

import com.fasterxml.jackson.annotation.JsonProperty

data class ResetPasswordRequest(
        @JsonProperty("email")
        val email: String,
        @JsonProperty("signature")
        val signature: String,
        @JsonProperty("password")
        val password: String,
        @JsonProperty("confirmation")
        val confirmation: String = ""
)
