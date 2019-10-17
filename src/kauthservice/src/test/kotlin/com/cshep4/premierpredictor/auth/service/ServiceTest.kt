package com.cshep4.premierpredictor.auth.service

import com.cshep4.premierpredictor.auth.email.Emailer
import com.cshep4.premierpredictor.auth.repository.Repository
import com.cshep4.premierpredictor.auth.tokeniser.Tokenizer
import com.nhaarman.mockitokotlin2.whenever
import org.hamcrest.CoreMatchers.`is`
import org.hamcrest.MatcherAssert.assertThat
import org.junit.Test
import org.junit.runner.RunWith
import org.mockito.InjectMocks
import org.mockito.Mock
import org.mockito.junit.MockitoJUnitRunner

@RunWith(MockitoJUnitRunner::class)
internal class ServiceTest {
  companion object {
    const val TOKEN = "🎫"
  }

  @Mock
  private lateinit var tokenizer: Tokenizer

  @Mock
  private lateinit var repository: Repository

  @Mock
  private lateinit var emailer: Emailer

  @InjectMocks
  private lateinit var service: Service

  @Test
  fun `'validate' returns true when token is valid`() {
    whenever(tokenizer.validateToken(TOKEN)).thenReturn(true)

    val result = service.validate(TOKEN)

    assertThat(result, `is`(true))
  }

  @Test
  fun `'validate' returns false when token is not valid`() {
    whenever(tokenizer.validateToken(TOKEN)).thenReturn(false)

    val result = service.validate(TOKEN)

    assertThat(result, `is`(false))
  }
}
