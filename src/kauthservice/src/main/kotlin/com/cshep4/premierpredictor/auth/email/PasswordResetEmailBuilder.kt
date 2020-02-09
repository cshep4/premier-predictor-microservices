package com.cshep4.premierpredictor.auth.email

import com.cshep4.premierpredictor.auth.model.EmailArgs

class PasswordResetEmailBuilder {
  companion object {
    const val RESET_PASSWORD_LINK = "https://premierpredictor.uk/reset-password"
    const val SENDER_EMAIL = "shepapps4@gmail.com"
  }

  private lateinit var sender: String
  private lateinit var recipient: String
  private lateinit var senderEmail: String
  private lateinit var recipientEmail: String
  private lateinit var subject: String
  private lateinit var content: String

  fun withSender(): PasswordResetEmailBuilder {
    this.sender = SENDER_EMAIL
    this.senderEmail = SENDER_EMAIL
    return this
  }

  fun withRecipient(recipient: String): PasswordResetEmailBuilder {
    this.recipient = recipient
    return this
  }

  fun withRecipientEmail(recipientEmail: String): PasswordResetEmailBuilder {
    this.recipientEmail = recipientEmail
    return this
  }

  fun withSubject(): PasswordResetEmailBuilder {
    this.subject = "Premier Predictor Password Reset"
    return this
  }

  fun withMessage(email: String, signature: String): PasswordResetEmailBuilder {
    val link = "$RESET_PASSWORD_LINK?email=$email&signature=$signature"

    content = """Hi,

                |We have received a request to reset your password.

                |To reset your password click on the following link or copy and paste this URL into your browser (link expires in 24 hours):

                |$link

                |If you don't want to reset your password then please ignore this message.

                |Regards,

                |The Premier Predictor Team""".trimMargin()

    return this
  }

  fun build() = EmailArgs(
    sender = this.sender,
    recipient = this.recipient,
    senderEmail = this.senderEmail,
    recipientEmail = this.recipientEmail,
    subject = this.subject,
    content = this.content
  )
}
