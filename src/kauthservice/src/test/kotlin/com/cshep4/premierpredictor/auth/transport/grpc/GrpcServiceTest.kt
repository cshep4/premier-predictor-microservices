package com.cshep4.premierpredictor.auth.transport.grpc

import com.cshep4.premierpredictor.auth.ValidateRequest
import com.cshep4.premierpredictor.auth.service.Service
import com.google.protobuf.Empty
import com.nhaarman.mockitokotlin2.any
import com.nhaarman.mockitokotlin2.times
import com.nhaarman.mockitokotlin2.verify
import com.nhaarman.mockitokotlin2.whenever
import io.grpc.Status.UNAUTHENTICATED
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
    const val TOKEN = "🔑"
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
  fun `'validate' will return UNAUTHENTICATED if token is null`(ctx: TestContext) {
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

      assertThat(status.status.code, `is`(UNAUTHENTICATED.code))
      assertThat(status.status.description, `is`("Token must be provided."))

      async.complete()
    })
  }

  @Test
  fun `'validate' will return UNAUTHENTICATED if there is a problem validating token`(ctx: TestContext) {
    val async = ctx.async()

    val req = ValidateRequest.newBuilder().setToken(TOKEN).build()
    val future = FutureFactoryImpl().future<Empty>()

    whenever(service.validate(TOKEN)).thenReturn(false)

    grpcService.validate(req, future.setHandler {
      if (it.succeeded()) {
        ctx.fail("future should not have succeeded")
      }

      verify(service).validate(TOKEN)

      assertThat(it.cause(), IsInstanceOf(StatusException::class.java))

      val status = it.cause() as StatusException

      assertThat(status.status, `is`(UNAUTHENTICATED))

      async.complete()
    })
  }

  @Test
  fun `'validate' will return ok if token is valid`(ctx: TestContext) {
    val async = ctx.async()

    val req = ValidateRequest.newBuilder().setToken(TOKEN).build()
    val future = FutureFactoryImpl().future<Empty>()

    whenever(service.validate(TOKEN)).thenReturn(true)

    grpcService.validate(req, future.setHandler {
      if (it.failed()) {
        ctx.fail("future should have succeeded")
      }

      verify(service).validate(TOKEN)

      async.complete()
    })
  }
}
