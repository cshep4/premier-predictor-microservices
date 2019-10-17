package com.cshep4.premierpredictor.auth.transport.grpc

import com.cshep4.premierpredictor.auth.*
import com.cshep4.premierpredictor.auth.extension.error
import com.cshep4.premierpredictor.auth.extension.ok
import com.cshep4.premierpredictor.auth.extension.unauthenticated
import com.cshep4.premierpredictor.auth.service.Service
import com.cshep4.premierpredictor.request.EmailRequest
import com.cshep4.premierpredictor.request.LoginRequest
import com.google.protobuf.Empty
import io.grpc.Status
import io.grpc.Status.UNAUTHENTICATED
import io.vertx.core.Future
import io.vertx.core.Handler

class GrpcService(private val service: Service) : AuthServiceGrpc.AuthServiceVertxImplBase() {
  override fun login(request: LoginRequest?, response: Future<LoginResponse>?) {
    super.login(request, response)
  }

  override fun register(request: RegisterRequest?, response: Future<Empty>?) {
    super.register(request, response)
  }

  override fun validate(request: ValidateRequest?, response: Future<Empty>) = when {
    request?.token.isNullOrEmpty() -> response.error(Status.UNAUTHENTICATED.withDescription("Token must be provided."))
    !service.validate(request!!.token) -> response.unauthenticated()
    else -> response.ok()
  }

  override fun initiatePasswordReset(request: EmailRequest?, response: Future<Empty>?) {
    super.initiatePasswordReset(request, response)
  }

  override fun resetPassword(request: ResetPasswordRequest?, response: Future<Empty>?) {
    super.resetPassword(request, response)
  }
}
