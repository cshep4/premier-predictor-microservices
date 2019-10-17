package com.cshep4.premierpredictor.auth.email

import com.cshep4.premierpredictor.email.EmailServiceGrpc

class Emailer(private val client: EmailServiceGrpc.EmailServiceVertxStub) {
}
