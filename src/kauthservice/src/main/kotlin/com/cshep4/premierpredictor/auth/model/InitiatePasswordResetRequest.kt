package com.cshep4.premierpredictor.auth.model

import com.fasterxml.jackson.annotation.JsonProperty

data class InitiatePasswordResetRequest(
        @JsonProperty("email")
        val email: String
)
