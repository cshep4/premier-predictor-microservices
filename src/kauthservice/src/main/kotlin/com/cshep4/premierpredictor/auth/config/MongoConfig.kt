package com.cshep4.premierpredictor.auth.config

import io.vertx.core.json.JsonObject
import io.vertx.ext.mongo.IndexOptions

data class MongoConfig(
  val uri: String
) {
  class Index private constructor(val fields: JsonObject, val opts: IndexOptions = IndexOptions()) {
    companion object {
      val indexes = listOf(
        Index(
          fields = JsonObject(
            mapOf(
              Pair("email", 1)
            )
          ),
          opts = IndexOptions()
            .name("email_idx")
            .unique(true)
            .sparse(false)
        )
      )
    }
  }

  companion object {
    val indexes = Index.indexes
  }
}
