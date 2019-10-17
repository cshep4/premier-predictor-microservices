package com.cshep4.premierpredictor.auth

import com.cshep4.premierpredictor.auth.email.EmailClientBuilder
import com.cshep4.premierpredictor.auth.email.Emailer
import com.cshep4.premierpredictor.auth.repository.Repository
import com.cshep4.premierpredictor.auth.service.Service
import com.cshep4.premierpredictor.auth.tokeniser.Tokenizer
import com.cshep4.premierpredictor.auth.transport.grpc.GrpcService
import com.cshep4.premierpredictor.auth.transport.http.HttpHandler
import io.vertx.core.AbstractVerticle
import io.vertx.core.Handler
import io.vertx.grpc.VertxServerBuilder
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import java.io.IOException
import java.util.concurrent.TimeUnit.MILLISECONDS


class MainVerticle : AbstractVerticle() {
  private val port = System.getenv("PORT") ?: shutdown("port not set")
  private val healthPort = System.getenv("HEALTH_PORT") ?: shutdown("health port not set")
  private val emailAddr = System.getenv("EMAIL_ADDR") ?: shutdown("failed to get emailservice address")
  private val jwtSecret = System.getenv("JWT_SECRET") ?: shutdown("jwt secret not set")

  override fun start() {
    val emailer = initEmailer()
    val repository = initRepository()
    val tokeniser = Tokenizer(vertx, jwtSecret)

    val service = Service(repository, emailer, tokeniser)

    val grpcService = GrpcService(service)
    val handler = HttpHandler(service)

    GlobalScope.launch {
      try {
        VertxServerBuilder
          .forPort(vertx, port.toInt())
          .nettyBuilder()
          .keepAliveTime(60000, MILLISECONDS)
          .permitKeepAliveWithoutCalls(true)
          .addService(grpcService)
          .build()
          .start()

        System.out.println("gRPC service started at localhost:$port")
      } catch (e: IOException) {
        e.printStackTrace()
        System.exit(1)
      }
    }

    GlobalScope.launch {
      vertx.createHttpServer()
        .requestHandler(handler.route(vertx))
        .listen(healthPort.toInt()) { res ->
          if (res.succeeded()) {
            System.out.println("HTTP server listening at: http://localhost:$healthPort/")
          } else {
            res.cause().printStackTrace()
            System.exit(1)
          }
        }
    }
  }

  private fun initEmailer(): Emailer {
    val client = EmailClientBuilder
      .forTarget(vertx, emailAddr)
      .usePlaintext()
      .keepAliveTime(2 * 60000, MILLISECONDS)
      .keepAliveTimeout(20000, MILLISECONDS)
      .keepAliveWithoutCalls(true)
      .build()

//    val req = SendEmailRequest.newBuilder()
//      .setSender("Chris Shepherd")
//      .setRecipient("Chris Shepherd")
//      .setSenderEmail("shepapps4@gmail.com")
//      .setRecipientEmail("chris_shepherd2@hotmail.com")
//      .setSubject("Hello")
//      .setContent("Hello, test")
//      .build()
//    client.send(req) {
//      if (!it.succeeded()) {
//        it.cause().printStackTrace()
//      }
//    }
    return Emailer(client)
  }

  private fun initRepository(): Repository {
    return Repository()
      .init(vertx, Handler {
        if (it.failed()) {
          it.cause().printStackTrace()
          System.exit(1)
        }
      })
  }

  private fun shutdown(message: String): String {
    System.out.println(message)
    System.exit(1)
    return ""
  }
}
