package com.cshep4.premierpredictor.auth.service

import com.cshep4.premierpredictor.auth.constant.Messages.COULD_NOT_RESET_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.SIGNATURE_EXPIRED_ERROR
import com.cshep4.premierpredictor.auth.email.Emailer
import com.cshep4.premierpredictor.auth.email.PasswordResetEmailBuilder
import com.cshep4.premierpredictor.auth.exception.SignatureNotValidException
import com.cshep4.premierpredictor.auth.model.ResetPasswordArgs
import com.cshep4.premierpredictor.auth.model.SignUpUser
import com.cshep4.premierpredictor.auth.model.User
import com.cshep4.premierpredictor.auth.repository.Repository
import com.cshep4.premierpredictor.auth.result.*
import com.cshep4.premierpredictor.auth.service.Service.Companion.LOGIN_SUBJECT
import com.cshep4.premierpredictor.auth.service.Service.Companion.REGISTER_SUBJECT
import com.cshep4.premierpredictor.auth.tokeniser.Tokenizer
import com.nhaarman.mockitokotlin2.any
import com.nhaarman.mockitokotlin2.times
import com.nhaarman.mockitokotlin2.verify
import com.nhaarman.mockitokotlin2.whenever
import org.hamcrest.CoreMatchers.`is`
import org.hamcrest.CoreMatchers.instanceOf
import org.hamcrest.MatcherAssert.assertThat
import org.hamcrest.core.IsInstanceOf
import org.junit.Test
import org.junit.runner.RunWith
import org.mindrot.jbcrypt.BCrypt
import org.mindrot.jbcrypt.BCrypt.gensalt
import org.mockito.InjectMocks
import org.mockito.Mock
import org.mockito.junit.MockitoJUnitRunner

@RunWith(MockitoJUnitRunner::class)
internal class ServiceTest {
  companion object {
    const val ID = "👤"
    const val EMAIL = "📨"
    const val FIRST_NAME = "first name"
    const val LAST_NAME = "last name"
    const val PASSWORD = "🔑"
    const val TOKEN = "🎫"
    const val SIGNATURE = "✒️"
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
  fun `'login' returns error when user cannot be found in store`() {
    val err = GetByEmailResult.Error(message = "error")

    whenever(repository.getByEmail(EMAIL)).thenReturn(err)

    val result = service.login(EMAIL, PASSWORD)

    assertThat(result, instanceOf(LoginResult.Error::class.java))

    val error = result as LoginResult.Error

    assertThat(error.message, `is`(err.message))
    verify(tokenizer, times(0)).generateToken(EMAIL, LOGIN_SUBJECT)
  }

  @Test
  fun `'login' returns error when password doesn't match hashed password in db`() {
    val repoSuccess = GetByEmailResult.Success(user = User(id = ID, email = EMAIL, password = PASSWORD))

    whenever(repository.getByEmail(EMAIL)).thenReturn(repoSuccess)

    val result = service.login(EMAIL, PASSWORD)

    assertThat(result, instanceOf(LoginResult.Error::class.java))

    val error = result as LoginResult.Error

    assertThat(error.message, `is`("invalid password"))
    verify(tokenizer, times(0)).generateToken(EMAIL, LOGIN_SUBJECT)
  }

  @Test
  fun `'login' finds user in store, checks password and returns id and generated token`() {
    val password = BCrypt.hashpw(PASSWORD, gensalt())
    val repoSuccess = GetByEmailResult.Success(user = User(id = ID, email = EMAIL, password = password))

    whenever(repository.getByEmail(EMAIL)).thenReturn(repoSuccess)
    whenever(tokenizer.generateToken(EMAIL, LOGIN_SUBJECT)).thenReturn(TOKEN)

    val result = service.login(EMAIL, PASSWORD)

    assertThat(result, IsInstanceOf(LoginResult.Success::class.java))

    val success = result as LoginResult.Success

    assertThat(success.id, `is`(ID))
    assertThat(success.token, `is`(TOKEN))
  }

  @Test
  fun `'register' returns error if email already exists`() {
    whenever(repository.getByEmail(EMAIL)).thenReturn(GetByEmailResult.Success(user = User(id = ID, email = EMAIL, password = PASSWORD)))

    val signUpUser = SignUpUser(
      email = EMAIL,
      firstName = FIRST_NAME,
      surname = LAST_NAME,
      password = PASSWORD,
      confirmPassword = PASSWORD,
      predictedWinner = "winner"
    )

    val result = service.register(signUpUser)

    verify(tokenizer, times(0)).generateToken(EMAIL, REGISTER_SUBJECT)

    assertThat(result, instanceOf(RegisterResult.Error::class.java))

    val error = result as RegisterResult.Error

    assertThat(error, `is`(EMAIL_EXISTS_ERROR))
  }

  @Test
  fun `'register' returns error if there is an error storing user`() {
    whenever(repository.getByEmail(EMAIL)).thenReturn(GetByEmailResult.Error(message = "error"))

    val signUpUser = SignUpUser(
      email = EMAIL,
      firstName = FIRST_NAME,
      surname = LAST_NAME,
      password = PASSWORD,
      confirmPassword = PASSWORD,
      predictedWinner = "winner"
    )

    whenever(repository.storeUser(any())).thenReturn(StoreUserResult.Error(message = "error"))

    val result = service.register(signUpUser)

    verify(tokenizer, times(0)).generateToken(EMAIL, REGISTER_SUBJECT)

    assertThat(result, instanceOf(RegisterResult.Error::class.java))

    val error = result as RegisterResult.Error

    assertThat(error.message, `is`("error"))
  }

  @Test
  fun `'register' stores user and returns id with token`() {
    whenever(repository.getByEmail(EMAIL)).thenReturn(GetByEmailResult.Error(message = "error"))

    val signUpUser = SignUpUser(
      email = EMAIL,
      firstName = FIRST_NAME,
      surname = LAST_NAME,
      password = PASSWORD,
      confirmPassword = PASSWORD,
      predictedWinner = "winner"
    )

    whenever(repository.storeUser(any())).thenReturn(StoreUserResult.Success(id = ID))
    whenever(tokenizer.generateToken(EMAIL, REGISTER_SUBJECT)).thenReturn(TOKEN)

    val result = service.register(signUpUser)

    assertThat(result, IsInstanceOf(RegisterResult.Success::class.java))

    val success = result as RegisterResult.Success

    assertThat(success.id, `is`(ID))
    assertThat(success.token, `is`(TOKEN))
  }

  @Test
  fun `'validate' returns error when token is not valid`() {
    val err: ValidateTokenResult = ValidateTokenResult.Error(message = "error")

    whenever(tokenizer.validateToken(TOKEN)).thenReturn(err)

    val result = service.validate(TOKEN)

    assertThat(result, `is`(err))
  }

  @Test
  fun `'validate' returns success when token is valid`() {
    val success: ValidateTokenResult = ValidateTokenResult.Success

    whenever(tokenizer.validateToken(TOKEN)).thenReturn(success)

    val result = service.validate(TOKEN)

    assertThat(result, `is`(success))
  }

  @Test
  fun `'initiatePasswordReset' returns error when user cannot be found in store`() {
    val err = GetByEmailResult.Error(message = "error")

    whenever(repository.getByEmail(EMAIL)).thenReturn(err)

    val result = service.initiatePasswordReset(EMAIL)

    assertThat(result, instanceOf(InitiatePasswordResetResult.Error::class.java))

    val error = result as InitiatePasswordResetResult.Error

    assertThat(error.message, `is`(err.message))

    verify(tokenizer, times(0)).createSignature(EMAIL)
    verify(repository, times(0)).updateUser(any(), any())
    verify(emailer, times(0)).send(any())
  }

  @Test
  fun `'initiatePasswordReset' returns error when there is an error updating user`() {
    val repoSuccess = GetByEmailResult.Success(user = User(id = ID, email = EMAIL, password = PASSWORD, firstName = FIRST_NAME, surname = LAST_NAME))
    val err = UpdateUserResult.Error(message = "error")

    whenever(repository.getByEmail(EMAIL)).thenReturn(repoSuccess)
    whenever(tokenizer.createSignature(EMAIL)).thenReturn(SIGNATURE)
    whenever(repository.updateUser(ID, Pair("signature", SIGNATURE))).thenReturn(err)

    val result = service.initiatePasswordReset(EMAIL)

    assertThat(result, instanceOf(InitiatePasswordResetResult.Error::class.java))

    val error = result as InitiatePasswordResetResult.Error

    assertThat(error.message, `is`(err.message))

    verify(tokenizer).createSignature(EMAIL)
    verify(repository).updateUser(ID, Pair("signature", SIGNATURE))
    verify(emailer, times(0)).send(any())
  }

  @Test
  fun `'initiatePasswordReset' returns error when there is an error sending reset email`() {
    val repoSuccess = GetByEmailResult.Success(user = User(id = ID, email = EMAIL, password = PASSWORD, firstName = FIRST_NAME, surname = LAST_NAME))
    val err = SendEmailResult.Error(message = "error")

    val emailArgs = PasswordResetEmailBuilder()
      .withSender()
      .withSubject()
      .withRecipientEmail(EMAIL)
      .withRecipient("$FIRST_NAME $LAST_NAME")
      .withMessage(EMAIL, SIGNATURE)
      .build()

    whenever(repository.getByEmail(EMAIL)).thenReturn(repoSuccess)
    whenever(tokenizer.createSignature(EMAIL)).thenReturn(SIGNATURE)
    whenever(repository.updateUser(ID, Pair("signature", SIGNATURE))).thenReturn(UpdateUserResult.Success)
    whenever(emailer.send(emailArgs)).thenReturn(err)

    val result = service.initiatePasswordReset(EMAIL)

    assertThat(result, instanceOf(InitiatePasswordResetResult.Error::class.java))

    val error = result as InitiatePasswordResetResult.Error

    assertThat(error.message, `is`(err.message))

    verify(tokenizer).createSignature(EMAIL)
    verify(repository).updateUser(ID, Pair("signature", SIGNATURE))
    verify(emailer).send(emailArgs)
  }

  @Test
  fun `'initiatePasswordReset' returns success when password reset initiated`() {
    val repoSuccess = GetByEmailResult.Success(user = User(id = ID, email = EMAIL, password = PASSWORD, firstName = FIRST_NAME, surname = LAST_NAME))

    val emailArgs = PasswordResetEmailBuilder()
      .withSender()
      .withSubject()
      .withRecipientEmail(EMAIL)
      .withRecipient("$FIRST_NAME $LAST_NAME")
      .withMessage(EMAIL, SIGNATURE)
      .build()

    whenever(repository.getByEmail(EMAIL)).thenReturn(repoSuccess)
    whenever(tokenizer.createSignature(EMAIL)).thenReturn(SIGNATURE)
    whenever(repository.updateUser(ID, Pair("signature", SIGNATURE))).thenReturn(UpdateUserResult.Success)
    whenever(emailer.send(emailArgs)).thenReturn(SendEmailResult.Success)

    val result = service.initiatePasswordReset(EMAIL)

    val success: InitiatePasswordResetResult = InitiatePasswordResetResult.Success
    assertThat(result, `is`(success))

    verify(tokenizer).createSignature(EMAIL)
    verify(repository).updateUser(ID, Pair("signature", SIGNATURE))
    verify(emailer).send(emailArgs)
  }

  @Test
  fun `'resetPassword' returns error when signature is not valid`() {
    val err = ValidateTokenResult.Error(message = "error")

    whenever(tokenizer.validateToken(SIGNATURE)).thenReturn(err)

    val args = ResetPasswordArgs(
      email = EMAIL,
      signature = SIGNATURE,
      password = PASSWORD
    )

    val result = service.resetPassword(args)

    assertThat(result, instanceOf(ResetPasswordResult.Error::class.java))

    val error = result as ResetPasswordResult.Error

    assertThat(error.message, `is`(SIGNATURE_EXPIRED_ERROR))

    verify(repository, times(0)).getByEmail(any())
    verify(repository, times(0)).updateUser(any(), any())
  }

  @Test
  fun `'resetPassword' returns error when user cannot be retrieved`() {
    val err = GetByEmailResult.Error(message = "error")

    whenever(tokenizer.validateToken(SIGNATURE)).thenReturn(ValidateTokenResult.Success)
    whenever(repository.getByEmail(EMAIL)).thenReturn(err)

    val args = ResetPasswordArgs(
      email = EMAIL,
      signature = SIGNATURE,
      password = PASSWORD
    )

    val result = service.resetPassword(args)

    assertThat(result, instanceOf(ResetPasswordResult.Error::class.java))

    val error = result as ResetPasswordResult.Error

    assertThat(error.message, `is`(COULD_NOT_RESET_ERROR))

    verify(repository).getByEmail(EMAIL)
    verify(repository, times(0)).updateUser(any(), any())
  }

  @Test
  fun `'resetPassword' returns error when signatures don't match`() {
    val repoSuccess = GetByEmailResult.Success(user = User(id = ID, email = EMAIL, password = PASSWORD, firstName = FIRST_NAME, surname = LAST_NAME, signature = SIGNATURE))

    whenever(tokenizer.validateToken("🖋️")).thenReturn(ValidateTokenResult.Success)
    whenever(repository.getByEmail(EMAIL)).thenReturn(repoSuccess)

    val args = ResetPasswordArgs(
      email = EMAIL,
      signature = "🖋️",
      password = PASSWORD
    )

    val result = service.resetPassword(args)

    assertThat(result, instanceOf(ResetPasswordResult.Error::class.java))

    val error = result as ResetPasswordResult.Error

    assertThat(error.message, `is`(COULD_NOT_RESET_ERROR))
    assertThat(error.cause, instanceOf(SignatureNotValidException::class.java))

    verify(repository).getByEmail(EMAIL)
    verify(repository, times(0)).updateUser(any(), any())
  }

  @Test
  fun `'resetPassword' returns error when user cannot be updated`() {
    val repoSuccess = GetByEmailResult.Success(user = User(id = ID, email = EMAIL, password = PASSWORD, firstName = FIRST_NAME, surname = LAST_NAME, signature = SIGNATURE))
    val err = UpdateUserResult.Error(message = "error")

    whenever(tokenizer.validateToken(SIGNATURE)).thenReturn(ValidateTokenResult.Success)
    whenever(repository.getByEmail(EMAIL)).thenReturn(repoSuccess)
    whenever(repository.updateUser(any(), any())).thenReturn(err)

    val args = ResetPasswordArgs(
      email = EMAIL,
      signature = SIGNATURE,
      password = PASSWORD
    )

    val result = service.resetPassword(args)

    assertThat(result, instanceOf(ResetPasswordResult.Error::class.java))

    val error = result as ResetPasswordResult.Error

    assertThat(error.message, `is`(COULD_NOT_RESET_ERROR))

    verify(repository).getByEmail(EMAIL)
    verify(repository).updateUser(any(), any())
  }

  @Test
  fun `'resetPassword' returns success when password is reset`() {
    val repoSuccess = GetByEmailResult.Success(user = User(id = ID, email = EMAIL, password = PASSWORD, firstName = FIRST_NAME, surname = LAST_NAME, signature = SIGNATURE))

    whenever(tokenizer.validateToken(SIGNATURE)).thenReturn(ValidateTokenResult.Success)
    whenever(repository.getByEmail(EMAIL)).thenReturn(repoSuccess)
    whenever(repository.updateUser(any(), any())).thenReturn(UpdateUserResult.Success)

    val args = ResetPasswordArgs(
      email = EMAIL,
      signature = SIGNATURE,
      password = PASSWORD
    )

    val result = service.resetPassword(args)

    val success: ResetPasswordResult = ResetPasswordResult.Success
    assertThat(result, `is`(success))

    verify(repository).getByEmail(EMAIL)
    verify(repository).updateUser(any(), any())
  }
}
