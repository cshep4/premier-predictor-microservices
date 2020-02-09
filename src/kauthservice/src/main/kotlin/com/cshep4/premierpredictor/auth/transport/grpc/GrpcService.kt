package com.cshep4.premierpredictor.auth.transport.grpc

import com.cshep4.premierpredictor.auth.*
import com.cshep4.premierpredictor.auth.constant.Messages.EMAIL_ALREADY_EXISTS_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.INITIATE_PASSWORD_RESET_VALIDATION_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.LOGIN_INTERNAL_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.LOGIN_INVALID_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.LOGIN_VALIDATION_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.REGISTER_VALIDATION_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.TOKEN_VALIDATION_ERROR
import com.cshep4.premierpredictor.auth.exception.InternalException
import com.cshep4.premierpredictor.auth.extension.*
import com.cshep4.premierpredictor.auth.model.ResetPasswordArgs
import com.cshep4.premierpredictor.auth.result.*
import com.cshep4.premierpredictor.auth.service.Service
import com.cshep4.premierpredictor.request.EmailRequest
import com.cshep4.premierpredictor.request.LoginRequest
import com.google.protobuf.Empty
import io.grpc.Status.INTERNAL
import io.grpc.Status.INVALID_ARGUMENT
import io.grpc.Status.UNAUTHENTICATED
import io.vertx.core.Future

class GrpcService(private val service: Service) : AuthServiceGrpc.AuthServiceVertxImplBase() {

  override fun login(request: LoginRequest?, response: Future<LoginResponse>) {
    if (!request.isValid()) {
      return response.error(INVALID_ARGUMENT.withDescription(LOGIN_VALIDATION_ERROR))
    }

    when (val result = service.login(request!!.email, request.password)) {
      is LoginResult.Error -> {
        if (result.cause is InternalException) {
          System.out.println("Failed to login for email: ${request.email}: ${result.cause.message}")
          response.error(UNAUTHENTICATED.withDescription(LOGIN_INTERNAL_ERROR))
          return
        }

        response.error(UNAUTHENTICATED.withDescription(LOGIN_INVALID_ERROR))
      }
      is LoginResult.Success -> {
        val resp = LoginResponse.newBuilder()
          .setId(result.id)
          .setToken(result.token)
          .build()

        response.ok(resp)
      }
    }
  }

  override fun register(request: RegisterRequest?, response: Future<Empty>) {
    if (!request.isValid()) {
      return response.error(INVALID_ARGUMENT.withDescription(REGISTER_VALIDATION_ERROR))
    }

    val signUpUser = request!!.toSignUpUser()

    when (val result = service.register(signUpUser)) {
      EMAIL_EXISTS_ERROR -> response.error(INVALID_ARGUMENT.withDescription(EMAIL_ALREADY_EXISTS_ERROR))
      is RegisterResult.Error -> {
        if (result.cause is InternalException) {
          System.out.println("Failed to register for email: ${request.email}: ${result.cause.message}")
        }
        response.error(INTERNAL.withDescription("${result.message}, please try again"))
      }
      is RegisterResult.Success -> {
        val resp = LoginResponse.newBuilder()
          .setId(result.id)
          .setToken(result.token)
          .build()
        response.ok()
      }
    }
  }

  override fun validate(request: ValidateRequest?, response: Future<Empty>) {
    if (!request.isValid()) {
      return response.error(INVALID_ARGUMENT.withDescription(TOKEN_VALIDATION_ERROR))
    }

    when (val result = service.validate(request!!.token)) {
      is ValidateTokenResult.Error -> {
        val status = UNAUTHENTICATED
          .withDescription(result.message)
          .withCause(result.cause)
        response.error(status)
      }
      else -> response.ok()
    }
  }

  override fun initiatePasswordReset(request: EmailRequest?, response: Future<Empty>) {
    if (!request.isValid()) {
      return response.error(INVALID_ARGUMENT.withDescription(INITIATE_PASSWORD_RESET_VALIDATION_ERROR))
    }

    when (val result = service.initiatePasswordReset(request!!.email)) {
      is InitiatePasswordResetResult.Error -> {
        if (result.cause is InternalException) {
          System.out.println("Failed to initiate password reset for email: ${request.email}: ${result.cause.message}")
        }
        response.error(INTERNAL.withDescription(result.message).withCause(result.cause))
      }
      else -> response.ok()
    }
  }

  override fun resetPassword(request: ResetPasswordRequest?, response: Future<Empty>) {
    when (val result = request.validate()) {
      is ResetPasswordRequestValidationResult.Error -> return response.error(INVALID_ARGUMENT.withDescription(result.message))
    }

    val req = request!!.toResetPasswordArgs()

    when (val result = service.resetPassword(req)) {
      is ResetPasswordResult.Error -> {
        if (result.cause is InternalException) {
          System.out.println("Failed to reset password for email: ${req.email}: ${result.cause.message}")
        }
        response.error(INTERNAL.withDescription(result.message).withCause(result.cause))
      }
      else -> response.ok()
    }
  }
}
