package com.cshep4.premierpredictor.auth.repository

import io.vertx.core.AsyncResult
import io.vertx.core.Future
import io.vertx.core.Handler
import io.vertx.core.Vertx
import io.vertx.core.http.HttpServer
import io.vertx.core.json.JsonObject
import io.vertx.ext.mongo.IndexOptions
import io.vertx.ext.mongo.MongoClient

class Repository {
  companion object {
    const val DATABASE = "user"
    const val COLLECTION = "user"
  }

  private lateinit var client: MongoClient

  fun init(vertx: Vertx, handler: Handler<AsyncResult<Unit>>): Repository {
    val mongoScheme: String = System.getenv("MONGO_SCHEME") ?: ""
    val mongoUsername: String = System.getenv("MONGO_USERNAME") ?: ""
    val mongoPassword: String = System.getenv("MONGO_PASSWORD") ?: ""
    val mongoHost: String = System.getenv("MONGO_HOST") ?: ""
    val mongoPort: String = System.getenv("MONGO_PORT") ?: ""

    val mongoUri = when {
        mongoUsername.isEmpty() -> "$mongoScheme://$mongoHost:$mongoPort"
        mongoPassword.isEmpty() -> "$mongoScheme://$mongoHost:$mongoPort"
        else -> "$mongoScheme://$mongoUsername:$mongoPassword@$mongoHost"
    }

    val config = JsonObject(
      mapOf(
        Pair("connection_string", mongoUri),
        Pair("db_name", DATABASE)
      )
    )
    client = MongoClient.createShared(vertx, config)

    val field = JsonObject(mapOf(Pair("email", 1)))

    val opts = IndexOptions()
      .name("email_idx")
      .unique(true)
      .sparse(false)

    client.createIndexWithOptions(COLLECTION, field, opts) {
      if (!it.succeeded()) {
        handler.handle(Future.failedFuture(it.cause()))
        return@createIndexWithOptions
      }

      handler.handle(Future.succeededFuture())
    }

    return this
  }
}
