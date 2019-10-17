package com.cshep4.premierpredictor.auth.tokeniser

import com.cshep4.premierpredictor.auth.extension.toJwt
import com.cshep4.premierpredictor.auth.factory.JWTProviderFactory
import io.vertx.core.Vertx
import io.vertx.ext.auth.User
import io.vertx.kotlin.coroutines.awaitResult
import kotlinx.coroutines.runBlocking
import org.hamcrest.CoreMatchers.`is`
import org.hamcrest.MatcherAssert.assertThat
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
  fun `'generateToken' generates token not validated with different secret`() {
    runBlocking {
      val token = tokenizer.generateToken(EMAIL, SUBJECT)

      val otherProvider = JWTProviderFactory.create(vertx, "fake secret ")

      awaitResult<User> { otherProvider.authenticate(token.toJwt(), it) }
    }
  }

  @Test
  fun `'validateToken' will return true when valid token`() {
    val token = tokenizer.generateToken(EMAIL, SUBJECT)

    val valid = tokenizer.validateToken(token!!)

    assertThat(valid, `is`(true))
  }

  @Test
  fun `'validateToken' will return false when invalid token`() {
    val valid = tokenizer.validateToken("fake token")

    assertThat(valid, `is`(false))
  }

  @Test
  fun `'validateToken' will return false when valid token with different secret`() {
    val otherTokeniser = Tokenizer(vertx, "fake secret")
    val token = otherTokeniser.generateToken(EMAIL, SUBJECT)

    val valid = tokenizer.validateToken(token!!)

    assertThat(valid, `is`(false))
  }
}
