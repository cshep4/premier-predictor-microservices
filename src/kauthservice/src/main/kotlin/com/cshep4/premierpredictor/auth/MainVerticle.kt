package com.cshep4.premierpredictor.auth

import com.cshep4.premierpredictor.auth.config.MongoConfig
import com.cshep4.premierpredictor.auth.email.Emailer
import com.cshep4.premierpredictor.auth.repository.Repository
import com.cshep4.premierpredictor.auth.service.Service
import com.cshep4.premierpredictor.auth.tokeniser.Tokenizer
import com.cshep4.premierpredictor.auth.transport.grpc.GrpcService
import com.cshep4.premierpredictor.auth.transport.http.HttpHandler
import io.vertx.core.AbstractVerticle
import io.vertx.core.Handler
import io.vertx.core.Vertx
import io.vertx.core.VertxOptions
import io.vertx.ext.web.client.WebClient
import io.vertx.grpc.VertxServerBuilder
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import java.io.IOException
import java.util.concurrent.TimeUnit.MILLISECONDS


class MainVerticle : AbstractVerticle() {
  private val port = System.getenv("PORT") ?: shutdown("port not set")
  private val healthPort = System.getenv("HEALTH_PORT") ?: shutdown("health port not set")
  private val jwtSecret = System.getenv("JWT_SECRET") ?: shutdown("jwt secret not set")
  private val emailUrl = System.getenv("EMAIL_URL") ?: shutdown("email url not set")

  private val mongoCfg = MongoConfig(
    uri = System.getenv("MONGO_URI") ?: "mongodb://localhost:27017"
  )

  override fun start() {
//    Json.mapper.registerModule(KotlinModule())

    val client = WebClient.create(Vertx.vertx(VertxOptions().setBlockedThreadCheckInterval(1000 * 60 * 60)))
    val emailer = Emailer(client, emailUrl)

//    val req = EmailArgs(
//      sender = "Chris S",
//      recipient = "C Shep",
//      senderEmail = "chris_shepherd2@hotmail.com",
//      recipientEmail = "shepapps4@gmail.com",
//      subject = "Hello 2",
//      content = "Hello test"
//    )
//    val result = emailer.send(req)

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

  private fun initRepository(): Repository {
    return Repository(mongoCfg)
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
