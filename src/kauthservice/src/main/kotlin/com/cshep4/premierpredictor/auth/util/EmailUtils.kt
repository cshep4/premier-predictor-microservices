package com.cshep4.premierpredictor.auth.util

import com.cshep4.premierpredictor.auth.util.ServiceUtils.getEnv

object EmailUtils {
    const val RESET_PASSWORD_PATH = "/reset-password"

    val url = getEnv(
            key = "SERVICE_ADDR",
            default = "http://localhost:8080"
    )

    fun buildResetPasswordEmail(email: String, firstName: String, signature: String) =
            """Hi $firstName,
            
            |We have received a request to reset your password.

            |To reset your password click on the following link or copy and paste this URL into your browser (link expires in 24 hours):

            |$url$RESET_PASSWORD_PATH?email=$email&signature=$signature

            |If you don't want to reset your password then please ignore this message.

            |Regards,

            |The Premier Predictor Team""".trimMargin()
}