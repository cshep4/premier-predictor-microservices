package com.cshep4.premierpredictor.auth.service

import com.cshep4.premierpredictor.auth.email.Emailer
import com.cshep4.premierpredictor.auth.model.LoginRequest
import com.cshep4.premierpredictor.auth.model.LoginResponse
import com.cshep4.premierpredictor.auth.model.RegisterRequest
import com.cshep4.premierpredictor.auth.model.ResetPasswordRequest
import com.cshep4.premierpredictor.auth.repository.Repository
import com.cshep4.premierpredictor.auth.tokeniser.Tokenizer
import io.vertx.core.AsyncResult
import io.vertx.core.Future
import io.vertx.core.Handler

class Service(
  private val repository: Repository,
  private val emailer: Emailer,
  private val tokenizer: Tokenizer
) {

  fun login(req: LoginRequest, handler: Handler<AsyncResult<LoginResponse>>) {
  }

  fun register(req: RegisterRequest, handler: Handler<AsyncResult<Unit>>) {

  }

  fun validate(token: String): Boolean = tokenizer.validateToken(token)

  fun initiatePasswordReset(email: String, handler: Handler<AsyncResult<Unit>>) {

  }

  fun resetPassword(request: ResetPasswordRequest, handler: Handler<AsyncResult<Unit>>) {

  }
}
