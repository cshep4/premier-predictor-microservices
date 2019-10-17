package com.cshep4.premierpredictor.auth.transport.http

import com.cshep4.premierpredictor.auth.service.Service
import io.vertx.core.Vertx
import io.vertx.ext.web.Router
import io.vertx.ext.web.RoutingContext

class HttpHandler(private val service: Service) {
  fun route(vertx: Vertx): Router {
    val router = Router.router(vertx)

    router.get("/health")
      .handler(this::health)

    return router
  }

  private fun health(routingContext: RoutingContext) {
    routingContext.response()
      .putHeader("content-type", "application/json")
      .setStatusCode(200)
      .end("Service is healthy!")
  }
}
