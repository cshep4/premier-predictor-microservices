package com.cshep4.premierpredictor.auth.extension

import io.vertx.core.json.JsonObject

fun String?.toJwt(): JsonObject {
  return JsonObject(mapOf(Pair("jwt", this)))
}
