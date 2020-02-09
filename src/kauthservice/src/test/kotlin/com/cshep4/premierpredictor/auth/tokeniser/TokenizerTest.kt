package com.cshep4.premierpredictor.auth.tokeniser

import com.cshep4.premierpredictor.auth.extension.toJwt
import com.cshep4.premierpredictor.auth.factory.JWTProviderFactory
import com.cshep4.premierpredictor.auth.result.ValidateTokenResult
import io.vertx.core.Vertx
import io.vertx.ext.auth.User
import io.vertx.kotlin.coroutines.awaitResult
import kotlinx.coroutines.runBlocking
import org.hamcrest.CoreMatchers.`is`
import org.hamcrest.MatcherAssert.assertThat
import org.hamcrest.core.IsInstanceOf
import org.junit.Test

internal class TokenizerTest {
  companion object {
    const val SECRET = "🔑"
    const val EMAIL = "📨"
    const val SUBJECT = "🤙"
  }

  private val vertx = Vertx.vertx()

  private val tokenizer = Tokenizer(vertx, SECRET)

  private val provider = JWTProviderFactory.create(vertx, SECRET)

  @Test
  fun `'generateToken' generates token validated with secret`() {
    runBlocking {
      val token = tokenizer.generateToken(EMAIL, SUBJECT)

      awaitResult<User> { provider.authenticate(token.toJwt(), it) }
    }
  }

  @Test(expected = RuntimeException::class)
  fun `'generateToken' generates token, not validated with different secret`() {
    runBlocking {
      val token = tokenizer.generateToken(EMAIL, SUBJECT)

      val otherProvider = JWTProviderFactory.create(vertx, "fake secret")

      awaitResult<User> { otherProvider.authenticate(token.toJwt(), it) }
    }
  }

  @Test
  fun `'validateToken' will return success when valid token`() {
    val token = tokenizer.generateToken(EMAIL, SUBJECT)

    val result = tokenizer.validateToken(token!!)

    assertThat(result, IsInstanceOf(ValidateTokenResult.Success::class.java))
  }

  @Test
  fun `'validateToken' will return error when invalid token`() {
    val result = tokenizer.validateToken("fake token")

    assertThat(result, IsInstanceOf(ValidateTokenResult.Error::class.java))

    val err = result as ValidateTokenResult.Error

    assertThat(err.message, `is`("invalid token: fake token"))
    assertThat(err.cause, IsInstanceOf(RuntimeException::class.java))
  }

  @Test
  fun `'validateToken' will return error when valid token with different secret`() {
    val otherTokeniser = Tokenizer(vertx, "fake secret")
    val token = otherTokeniser.generateToken(EMAIL, SUBJECT)

    val result = tokenizer.validateToken(token!!)

    assertThat(result, IsInstanceOf(ValidateTokenResult.Error::class.java))

    val err = result as ValidateTokenResult.Error

    assertThat(err.message, `is`("invalid token: $token"))
    assertThat(err.cause, IsInstanceOf(RuntimeException::class.java))
  }

  @Test
  fun `'createSignature' generates token validated with secret`() {
    runBlocking {
      val token = tokenizer.createSignature(EMAIL)

      awaitResult<User> { provider.authenticate(token.toJwt(), it) }
    }
  }
}
