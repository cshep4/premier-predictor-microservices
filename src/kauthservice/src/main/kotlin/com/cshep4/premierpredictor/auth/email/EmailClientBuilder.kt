package com.cshep4.premierpredictor.auth.email

import com.cshep4.premierpredictor.email.EmailServiceGrpc
import io.vertx.core.Vertx
import io.vertx.grpc.VertxChannelBuilder
import java.util.concurrent.TimeUnit

class EmailClientBuilder private constructor(vertx: Vertx, address: String) {
  private val builder: VertxChannelBuilder = VertxChannelBuilder.forTarget(vertx, address)

  companion object {
    fun forTarget(vertx: Vertx, address: String): EmailClientBuilder {
      return EmailClientBuilder(vertx, address)
    }
  }

  fun usePlaintext(): EmailClientBuilder {
    builder.usePlaintext()
    return this
  }

  fun keepAliveTime(keepAliveTime: Long, timeUnit: TimeUnit): EmailClientBuilder {
    builder.keepAliveTime(keepAliveTime, timeUnit)
    return this
  }

  fun keepAliveTimeout(keepAliveTime: Long, timeUnit: TimeUnit): EmailClientBuilder {
    builder.keepAliveTimeout(keepAliveTime, timeUnit)
    return this
  }

  fun keepAliveWithoutCalls(enable: Boolean): EmailClientBuilder {
    builder.keepAliveWithoutCalls(enable)
    return this
  }

  fun build(): EmailServiceGrpc.EmailServiceVertxStub {
    val channel = builder.build()
    return EmailServiceGrpc.newVertxStub(channel)
  }
}
