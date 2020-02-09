package com.cshep4.premierpredictor.auth.model

data class EmailArgs(
  val sender: String,
  val recipient: String,
  val senderEmail: String,
  val recipientEmail: String,
  val subject: String,
  val content: String
)
