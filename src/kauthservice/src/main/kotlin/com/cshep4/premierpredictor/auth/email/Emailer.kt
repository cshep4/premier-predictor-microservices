package com.cshep4.premierpredictor.auth.email

import com.cshep4.premierpredictor.auth.model.SendEmailRequest

interface Emailer {
    fun send(sendEmailRequest: SendEmailRequest)
}