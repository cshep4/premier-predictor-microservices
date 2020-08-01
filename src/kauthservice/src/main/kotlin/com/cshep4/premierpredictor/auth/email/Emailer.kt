package com.cshep4.premierpredictor.auth.email

import com.cshep4.premierpredictor.auth.model.SendEmailRequest
import com.cshep4.premierpredictor.auth.result.SendEmailResult

interface Emailer {
    fun send(sendEmailRequest: SendEmailRequest): SendEmailResult
}