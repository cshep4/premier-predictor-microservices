package com.cshep4.premierpredictor.auth.extension

import io.vertx.core.json.JsonObject
import java.time.LocalDateTime
import java.time.format.DateTimeFormatter

fun JsonObject.getLocalDateTime(key: String): LocalDateTime? {
  try {
    val obj = this.getJsonObject(key) ?: return null

    val dateStr = obj.getString("\$date") ?: return null

    return LocalDateTime.parse(dateStr, DateTimeFormatter.ISO_ZONED_DATE_TIME)
  } catch (e: Exception) {
    return null
  }
}

fun <T> JsonObject.putIfNotNull(key: String, value: T): JsonObject {
  value ?: return this

  return this.put(key, value)
}
