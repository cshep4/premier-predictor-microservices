package com.cshep4.premierpredictor.auth.tokeniser

import com.cshep4.premierpredictor.auth.constant.SecurityConstants.EXPIRATION_TIME
import com.cshep4.premierpredictor.auth.constant.SecurityConstants.PASSWORD_RESET_EXPIRATION_TIME
import com.cshep4.premierpredictor.auth.extension.toJwt
import com.cshep4.premierpredictor.auth.factory.JWTProviderFactory
import com.cshep4.premierpredictor.auth.model.Token
import com.cshep4.premierpredictor.auth.result.ValidateTokenResult
import io.vertx.core.Vertx
import io.vertx.core.json.JsonObject.mapFrom
import io.vertx.ext.auth.User
import io.vertx.ext.jwt.JWTOptions
import io.vertx.kotlin.coroutines.awaitResult
import kotlinx.coroutines.runBlocking

class Tokenizer(vertx: Vertx, secret: String) {
  private val provider = JWTProviderFactory.create(vertx, secret)

  fun generateToken(email: String, subject: String): String {
    val jwt = Token(
      sub = subject + System.currentTimeMillis(),
      iss = email,
      exp = System.currentTimeMillis() + EXPIRATION_TIME
    )

    val opts = JWTOptions()
      .setAlgorithm("HS512")

    return provider.generateToken(mapFrom(jwt), opts)
  }

  fun validateToken(token: String): ValidateTokenResult = runBlocking {
    try {
      awaitResult<User> { provider.authenticate(token.toJwt(), it) }

      ValidateTokenResult.Success
    } catch (e: Exception) {
      ValidateTokenResult.Error(
        message = "invalid token: $token",
        cause = e
      )
    }
  }

  fun createSignature(email: String): String {
    val jwt = Token(
      sub = "user_signature" + System.currentTimeMillis(),
      iss = email,
      exp = System.currentTimeMillis() + PASSWORD_RESET_EXPIRATION_TIME
    )

    val opts = JWTOptions()
      .setAlgorithm("HS512")

    return provider.generateToken(mapFrom(jwt), opts)
  }

}
