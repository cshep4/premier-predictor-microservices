package com.cshep4.premierpredictor.auth.extension

import io.vertx.core.json.JsonObject.mapFrom

fun <T> T.jsonString(): String = mapFrom(this)
  .encode()
