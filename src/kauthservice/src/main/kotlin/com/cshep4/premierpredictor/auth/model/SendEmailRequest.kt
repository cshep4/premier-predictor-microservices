package com.cshep4.premierpredictor.auth.model

import io.quarkus.runtime.annotations.RegisterForReflection
import java.util.*

@RegisterForReflection
data class SendEmailRequest(
        val sender: String,
        val recipient: String,
        val senderEmail: String,
        val recipientEmail: String,
        val subject: String,
        val content: String,
        val htmlContent: String,
        val idempotencyKey: String = UUID.randomUUID().toString()
)
