package com.cshep4.premierpredictor.auth.email.http

import com.cshep4.premierpredictor.auth.email.Emailer
import com.cshep4.premierpredictor.auth.model.SendEmailRequest
import com.cshep4.premierpredictor.auth.result.SendEmailResult
import org.eclipse.microprofile.rest.client.inject.RegisterRestClient
import org.eclipse.microprofile.rest.client.inject.RestClient
import javax.inject.Inject
import javax.inject.Named
import javax.inject.Singleton
import javax.ws.rs.Consumes
import javax.ws.rs.POST
import javax.ws.rs.Produces
import javax.ws.rs.core.MediaType.APPLICATION_JSON

@Singleton
@Named("httpEmailer")
class Emailer : Emailer {
    @RegisterRestClient(configKey = "email-api")
    interface EmailClient {
        @POST
        @Consumes(APPLICATION_JSON)
        @Produces(APPLICATION_JSON)
        fun send(sendEmailRequest: SendEmailRequest)
    }

    @Inject
    @RestClient
    lateinit var emailClient: EmailClient

    override fun send(sendEmailRequest: SendEmailRequest): SendEmailResult = try {
        emailClient.send(sendEmailRequest)

        SendEmailResult.Success
    } catch (e: Exception) {
        SendEmailResult.Error(
                message = "could not send email",
                cause = e
        )
    }

}