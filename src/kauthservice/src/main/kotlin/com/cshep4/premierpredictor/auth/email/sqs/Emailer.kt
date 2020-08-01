package com.cshep4.premierpredictor.auth.email.sqs

import com.cshep4.premierpredictor.auth.email.Emailer
import com.cshep4.premierpredictor.auth.model.SendEmailRequest
import com.cshep4.premierpredictor.auth.result.SendEmailResult
import org.eclipse.microprofile.config.inject.ConfigProperty
import javax.inject.Named
import javax.inject.Singleton
import javax.ws.rs.Consumes
import javax.ws.rs.Produces
import javax.ws.rs.core.MediaType.APPLICATION_JSON

@Singleton
@Named("sqsEmailer")
class Emailer : Emailer {
    @ConfigProperty(name = "queue.url")
    var queueUrl: String = ""

//    @Inject
//    lateinit var sqs: SqsClient

    @Consumes(APPLICATION_JSON)
    @Produces(APPLICATION_JSON)
    override fun send(sendEmailRequest: SendEmailRequest): SendEmailResult = try {
//        sqs.sendMessage {
//            it.queueUrl(queueUrl).messageBody(emailArgs.toString())
//        }

        SendEmailResult.Success
    } catch (e: Exception) {
        SendEmailResult.Error(
                message = "could not send email",
                cause = e
        )
    }
}