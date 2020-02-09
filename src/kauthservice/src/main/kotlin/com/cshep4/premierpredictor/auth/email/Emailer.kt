package com.cshep4.premierpredictor.auth.email

import com.cshep4.premierpredictor.auth.model.EmailArgs
import com.cshep4.premierpredictor.auth.result.SendEmailResult
import io.vertx.core.buffer.Buffer
import io.vertx.core.json.JsonObject
import io.vertx.ext.web.client.HttpResponse
import io.vertx.ext.web.client.WebClient
import io.vertx.kotlin.coroutines.awaitResult
import kotlinx.coroutines.runBlocking
import java.net.HttpURLConnection.HTTP_OK
import java.util.*


class Emailer(private val client: WebClient, private val emailUrl: String) {
  fun send(req: EmailArgs): SendEmailResult = runBlocking {
    val msg = JsonObject(
      mapOf(
        Pair("sender", req.sender),
        Pair("recipient", req.recipient),
        Pair("senderEmail", req.senderEmail),
        Pair("recipientEmail", req.recipientEmail),
        Pair("subject", req.subject),
        Pair("content", req.content),
        Pair("idempotencyKey", UUID.randomUUID().toString())
      )
    )

    try {
      val resp = awaitResult<HttpResponse<Buffer>> {
        client
          .postAbs(emailUrl)
          .sendJsonObject(msg, it)
      }

      if (resp.statusCode() != HTTP_OK) {
        return@runBlocking SendEmailResult.Error(
          message = "error sending email",
          cause = Exception("status: ${resp.statusMessage()}")
        )
      }

      SendEmailResult.Success
    } catch (e: Exception) {
      SendEmailResult.Error(
        message = "error sending email",
        cause = e
      )
    }
  }
}
