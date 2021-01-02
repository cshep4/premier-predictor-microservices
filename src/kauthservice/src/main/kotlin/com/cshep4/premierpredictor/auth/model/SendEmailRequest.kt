package com.cshep4.premierpredictor.auth.model

import com.fasterxml.jackson.annotation.JsonProperty
import io.quarkus.runtime.annotations.RegisterForReflection
import java.util.*

@RegisterForReflection
data class SendEmailRequest(
        @JsonProperty("sender")
        val sender: String,
        @JsonProperty("recipient")
        val recipient: String,
        @JsonProperty("senderEmail")
        val senderEmail: String,
        @JsonProperty("recipientEmail")
        val recipientEmail: String,
        @JsonProperty("subject")
        val subject: String,
        @JsonProperty("content")
        val content: String,
        @JsonProperty("htmlContent")
        val htmlContent: String,
        @JsonProperty("idempotencyKey")
        val idempotencyKey: String = UUID.randomUUID().toString()
)
