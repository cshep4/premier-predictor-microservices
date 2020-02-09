package com.cshep4.premierpredictor.auth.transport.grpc

import com.cshep4.premierpredictor.auth.LoginResponse
import com.cshep4.premierpredictor.auth.RegisterRequest
import com.cshep4.premierpredictor.auth.ResetPasswordRequest
import com.cshep4.premierpredictor.auth.ValidateRequest
import com.cshep4.premierpredictor.auth.constant.Messages.COULD_NOT_RESET_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.EMAIL_ALREADY_EXISTS_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.INITIATE_PASSWORD_RESET_VALIDATION_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.LOGIN_INVALID_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.LOGIN_VALIDATION_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.PASSWORDS_DO_NOT_MATCH_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.PASSWORD_NOT_VALID_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.REGISTER_VALIDATION_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.TOKEN_VALIDATION_ERROR
import com.cshep4.premierpredictor.auth.extension.toResetPasswordArgs
import com.cshep4.premierpredictor.auth.extension.toSignUpUser
import com.cshep4.premierpredictor.auth.result.*
import com.cshep4.premierpredictor.auth.service.Service
import com.cshep4.premierpredictor.request.EmailRequest
import com.cshep4.premierpredictor.request.LoginRequest
import com.google.protobuf.Empty
import com.nhaarman.mockitokotlin2.any
import com.nhaarman.mockitokotlin2.times
import com.nhaarman.mockitokotlin2.verify
import com.nhaarman.mockitokotlin2.whenever
import io.grpc.Status.*
import io.grpc.StatusException
import io.vertx.core.impl.FutureFactoryImpl
import io.vertx.ext.unit.TestContext
import io.vertx.ext.unit.junit.VertxUnitRunner
import org.hamcrest.CoreMatchers.`is`
import org.hamcrest.MatcherAssert.assertThat
import org.hamcrest.core.IsInstanceOf
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith
import org.mockito.InjectMocks
import org.mockito.Mock
import org.mockito.MockitoAnnotations

@RunWith(VertxUnitRunner::class)
internal class GrpcServiceTest {
  companion object {
    const val ID = "👤"
    const val FIRST_NAME = "first name"
    const val LAST_NAME = "last name"
    const val EMAIL = "test@test.com"
    const val PASSWORD = "Qwerty123"
    const val TOKEN = "🎫"
  }

  @Mock
  private lateinit var service: Service

  @InjectMocks
  private lateinit var grpcService: GrpcService

  @Before
  fun init() {
    MockitoAnnotations.initMocks(this)
  }

  @Test
  fun `'login' will return INVALID_ARGUMENT if email is null`(ctx: TestContext) {
    val async = ctx.async()

    val req = LoginRequest.newBuilder()
      .setPassword(PASSWORD)
      .build()
    val future = FutureFactoryImpl().future<LoginResponse>()

    grpcService.login(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).login(any(), any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(LOGIN_VALIDATION_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'login' will return INVALID_ARGUMENT if password is null`(ctx: TestContext) {
    val async = ctx.async()

    val req = LoginRequest.newBuilder()
      .setEmail(EMAIL)
      .build()
    val future = FutureFactoryImpl().future<LoginResponse>()

    grpcService.login(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).login(any(), any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(LOGIN_VALIDATION_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'login' will return UNAUTHENTICATED if login details are invalid`(ctx: TestContext) {
    val async = ctx.async()

    val req = LoginRequest.newBuilder()
      .setEmail(EMAIL)
      .setPassword(PASSWORD)
      .build()
    val future = FutureFactoryImpl().future<LoginResponse>()

    whenever(service.login(EMAIL, PASSWORD)).thenReturn(LoginResult.Error("error", Exception("error")))

    grpcService.login(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))
      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(UNAUTHENTICATED.code))
      assertThat(status.status.description, `is`(LOGIN_INVALID_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'login' will return user id and token`(ctx: TestContext) {
    val async = ctx.async()

    val req = LoginRequest.newBuilder()
      .setEmail(EMAIL)
      .setPassword(PASSWORD)
      .build()
    val future = FutureFactoryImpl().future<LoginResponse>()

    whenever(service.login(EMAIL, PASSWORD)).thenReturn(LoginResult.Success(id = ID, token = TOKEN))

    grpcService.login(req, future.setHandler {
      if (it.failed()) {
        ctx.fail("future should have succeeded")
      }

      val result = it.result()

      assertThat(result.id, `is`(ID))
      assertThat(result.token, `is`(TOKEN))

      async.complete()
    })
  }

  @Test
  fun `'register' will return INVALID_ARGUMENT if first name is null`(ctx: TestContext) {
    val async = ctx.async()

    val req = RegisterRequest.newBuilder()
      .setSurname(LAST_NAME)
      .setEmail(EMAIL)
      .setPassword(PASSWORD)
      .setConfirmation(PASSWORD)
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.register(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).register(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(REGISTER_VALIDATION_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'register' will return INVALID_ARGUMENT if surname is null`(ctx: TestContext) {
    val async = ctx.async()

    val req = RegisterRequest.newBuilder()
      .setFirstName(FIRST_NAME)
      .setEmail(EMAIL)
      .setPassword(PASSWORD)
      .setConfirmation(PASSWORD)
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.register(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).register(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(REGISTER_VALIDATION_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'register' will return INVALID_ARGUMENT if email is null`(ctx: TestContext) {
    val async = ctx.async()

    val req = RegisterRequest.newBuilder()
      .setFirstName(FIRST_NAME)
      .setSurname(LAST_NAME)
      .setPassword(PASSWORD)
      .setConfirmation(PASSWORD)
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.register(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).register(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(REGISTER_VALIDATION_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'register' will return INVALID_ARGUMENT if email is invalid`(ctx: TestContext) {
    val async = ctx.async()

    val req = RegisterRequest.newBuilder()
      .setFirstName(FIRST_NAME)
      .setSurname(LAST_NAME)
      .setEmail("📨")
      .setPassword(PASSWORD)
      .setConfirmation(PASSWORD)
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.register(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).register(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(REGISTER_VALIDATION_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'register' will return INVALID_ARGUMENT if password is null`(ctx: TestContext) {
    val async = ctx.async()

    val req = RegisterRequest.newBuilder()
      .setFirstName(FIRST_NAME)
      .setSurname(LAST_NAME)
      .setEmail(EMAIL)
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.register(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).register(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(REGISTER_VALIDATION_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'register' will return INVALID_ARGUMENT if password is invalid`(ctx: TestContext) {
    val async = ctx.async()

    val req = RegisterRequest.newBuilder()
      .setFirstName(FIRST_NAME)
      .setSurname(LAST_NAME)
      .setEmail(EMAIL)
      .setPassword("🔑")
      .setConfirmation("🔑")
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.register(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).register(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(REGISTER_VALIDATION_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'register' will return INVALID_ARGUMENT if password is not equal to confirmation`(ctx: TestContext) {
    val async = ctx.async()

    val req = RegisterRequest.newBuilder()
      .setFirstName(FIRST_NAME)
      .setSurname(LAST_NAME)
      .setEmail(EMAIL)
      .setPassword(PASSWORD)
      .setConfirmation("🔑")
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.register(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).register(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(REGISTER_VALIDATION_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'register' will return INVALID_ARGUMENT if email already exists`(ctx: TestContext) {
    val async = ctx.async()

    val req = RegisterRequest.newBuilder()
      .setFirstName(FIRST_NAME)
      .setSurname(LAST_NAME)
      .setEmail(EMAIL)
      .setPassword(PASSWORD)
      .setConfirmation(PASSWORD)
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    val signUpUser = req.toSignUpUser()


    whenever(service.register(signUpUser)).thenReturn(com.cshep4.premierpredictor.auth.result.EMAIL_EXISTS_ERROR)

    grpcService.register(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(EMAIL_ALREADY_EXISTS_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'register' will return INTERNAL if error registering`(ctx: TestContext) {
    val async = ctx.async()

    val req = RegisterRequest.newBuilder()
      .setFirstName(FIRST_NAME)
      .setSurname(LAST_NAME)
      .setEmail(EMAIL)
      .setPassword(PASSWORD)
      .setConfirmation(PASSWORD)
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    val signUpUser = req.toSignUpUser()

    whenever(service.register(signUpUser)).thenReturn(RegisterResult.Error("error", Exception("error")))

    grpcService.register(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))
      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INTERNAL.code))
      assertThat(status.status.description, `is`("error, please try again"))

      async.complete()
    })
  }

  //TODO - create register happy path test

  @Test
  fun `'validate' will return INVALID_ARGUMENT if token is null`(ctx: TestContext) {
    val async = ctx.async()

    val req = ValidateRequest.newBuilder().build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.validate(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).validate(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(TOKEN_VALIDATION_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'validate' will return UNAUTHENTICATED if there is a problem validating token`(ctx: TestContext) {
    val async = ctx.async()

    val req = ValidateRequest.newBuilder().setToken(TOKEN).build()
    val future = FutureFactoryImpl().future<Empty>()

    whenever(service.validate(TOKEN)).thenReturn(ValidateTokenResult.Error("error", Exception("error")))

    grpcService.validate(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service).validate(TOKEN)

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))
      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(UNAUTHENTICATED.code))
      assertThat(status.status.description, `is`("error"))

      async.complete()
    })
  }

  @Test
  fun `'validate' will return ok if token is valid`(ctx: TestContext) {
    val async = ctx.async()

    val req = ValidateRequest.newBuilder().setToken(TOKEN).build()
    val future = FutureFactoryImpl().future<Empty>()

    whenever(service.validate(TOKEN)).thenReturn(ValidateTokenResult.Success)

    grpcService.validate(req, future.setHandler {
      if (it.failed()) {
        ctx.fail("future should have succeeded")
      }

      verify(service).validate(TOKEN)

      async.complete()
    })
  }

  @Test
  fun `'initiatePasswordReset' will return INVALID_ARGUMENT if email is null`(ctx: TestContext) {
    val async = ctx.async()

    val req = EmailRequest.newBuilder().build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.initiatePasswordReset(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).initiatePasswordReset(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(INITIATE_PASSWORD_RESET_VALIDATION_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'initiatePasswordReset' will return INVALID_ARGUMENT if email is invalid`(ctx: TestContext) {
    val async = ctx.async()

    val req = EmailRequest.newBuilder()
      .setEmail("📨")
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.initiatePasswordReset(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).initiatePasswordReset(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(INITIATE_PASSWORD_RESET_VALIDATION_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'initiatePasswordReset' will return INTERNAL if error initiating password reset`(ctx: TestContext) {
    val async = ctx.async()

    val req = EmailRequest.newBuilder()
      .setEmail(EMAIL)
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    whenever(service.initiatePasswordReset(EMAIL)).thenReturn(InitiatePasswordResetResult.Error("error", Exception("error")))

    grpcService.initiatePasswordReset(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service).initiatePasswordReset(EMAIL)

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))
      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INTERNAL.code))
      assertThat(status.status.description, `is`("error"))

      async.complete()
    })
  }

  @Test
  fun `'initiatePasswordReset' will return ok when password reset initiated`(ctx: TestContext) {
    val async = ctx.async()

    val req = EmailRequest.newBuilder()
      .setEmail(EMAIL)
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    whenever(service.initiatePasswordReset(EMAIL)).thenReturn(InitiatePasswordResetResult.Success)

    grpcService.initiatePasswordReset(req, future.setHandler {
      if (it.failed()) {
        ctx.fail("future should have succeeded")
      }

      verify(service).initiatePasswordReset(EMAIL)

      async.complete()
    })
  }

  @Test
  fun `'resetPassword' will return INVALID_ARGUMENT if email is null`(ctx: TestContext) {
    val async = ctx.async()

    val req = ResetPasswordRequest.newBuilder().build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.resetPassword(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).resetPassword(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(COULD_NOT_RESET_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'resetPassword' will return INVALID_ARGUMENT if email is invalid`(ctx: TestContext) {
    val async = ctx.async()

    val req = ResetPasswordRequest.newBuilder()
      .setEmail("📨")
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.resetPassword(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).resetPassword(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(COULD_NOT_RESET_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'resetPassword' will return INVALID_ARGUMENT if password is null`(ctx: TestContext) {
    val async = ctx.async()

    val req = ResetPasswordRequest.newBuilder()
      .setEmail(EMAIL)
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.resetPassword(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).resetPassword(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(PASSWORD_NOT_VALID_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'resetPassword' will return INVALID_ARGUMENT if password is invalid`(ctx: TestContext) {
    val async = ctx.async()

    val req = ResetPasswordRequest.newBuilder()
      .setEmail(EMAIL)
      .setPassword("🔑")
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.resetPassword(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).resetPassword(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(PASSWORD_NOT_VALID_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'resetPassword' will return INVALID_ARGUMENT if passwords do not match`(ctx: TestContext) {
    val async = ctx.async()

    val req = ResetPasswordRequest.newBuilder()
      .setEmail(EMAIL)
      .setPassword(PASSWORD)
      .setConfirmation("🔑")
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    grpcService.resetPassword(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service, times(0)).resetPassword(any())

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INVALID_ARGUMENT.code))
      assertThat(status.status.description, `is`(PASSWORDS_DO_NOT_MATCH_ERROR))

      async.complete()
    })
  }

  @Test
  fun `'resetPassword' will return INTERNAL if error resetting password`(ctx: TestContext) {
    val async = ctx.async()

    val req = ResetPasswordRequest.newBuilder()
      .setEmail(EMAIL)
      .setPassword(PASSWORD)
      .setConfirmation(PASSWORD)
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    val resetPasswordArgs = req.toResetPasswordArgs()

    whenever(service.resetPassword(resetPasswordArgs)).thenReturn(ResetPasswordResult.Error("error", Exception("error")))

    grpcService.resetPassword(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service).resetPassword(resetPasswordArgs)

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))
      val status = it.cause() as StatusException

      assertThat(status.status.code, `is`(INTERNAL.code))
      assertThat(status.status.description, `is`("error"))

      async.complete()
    })
  }

  @Test
  fun `'resetPassword' will return ok when password is reset`(ctx: TestContext) {
    val async = ctx.async()

    val req = ResetPasswordRequest.newBuilder()
      .setEmail(EMAIL)
      .setPassword(PASSWORD)
      .setConfirmation(PASSWORD)
      .build()
    val future = FutureFactoryImpl().future<Empty>()

    val resetPasswordArgs = req.toResetPasswordArgs()

    whenever(service.resetPassword(resetPasswordArgs)).thenReturn(ResetPasswordResult.Success)

    grpcService.resetPassword(req, future.setHandler {
      if (it.failed()) {
        ctx.fail("future should have succeeded")
      }

      verify(service).resetPassword(resetPasswordArgs)

      async.complete()
    })
  }
}
