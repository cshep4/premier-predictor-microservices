package com.cshep4.premierpredictor.auth.entity

import com.cshep4.premierpredictor.auth.extension.getLocalDateTime
import com.cshep4.premierpredictor.auth.extension.putIfNotNull
import com.cshep4.premierpredictor.auth.model.SignUpUser
import com.cshep4.premierpredictor.auth.model.User
import com.fasterxml.jackson.annotation.JsonFormat
import com.fasterxml.jackson.annotation.JsonIgnoreProperties
import com.fasterxml.jackson.annotation.JsonProperty
import com.fasterxml.jackson.core.JsonGenerator
import com.fasterxml.jackson.core.JsonParser
import com.fasterxml.jackson.databind.DeserializationContext
import com.fasterxml.jackson.databind.JsonDeserializer
import com.fasterxml.jackson.databind.JsonSerializer
import com.fasterxml.jackson.databind.SerializerProvider
import com.fasterxml.jackson.databind.annotation.JsonDeserialize
import com.fasterxml.jackson.databind.annotation.JsonSerialize
import io.vertx.core.json.JsonObject
import org.bson.types.ObjectId
import java.io.IOException
import java.time.LocalDateTime
import java.time.format.DateTimeFormatter
import java.time.LocalDate
import java.time.format.DateTimeFormatter.ISO_DATE_TIME
import com.fasterxml.jackson.databind.deser.std.StdDeserializer
import com.fasterxml.jackson.databind.ser.std.StdSerializer
import java.time.ZoneOffset
import java.time.ZoneOffset.UTC
import java.time.format.DateTimeFormatter.ISO_ZONED_DATE_TIME

class UserEntity(
  val id: ObjectId,
  var firstName: String = "",
  var surname: String = "",
  var email: String,
  var password: String,
  var predictedWinner: String = "",
  var score: Int = 0,
  var admin: Boolean = false,
  var adFree: Boolean = false,
  val joined: LocalDateTime? = null,
  var signature: String? = null
) {
  fun toJson() = JsonObject()
    .putIfNotNull("_id", id.toHexString())
    .putIfNotNull("firstName", firstName)
    .putIfNotNull("surname", surname)
    .putIfNotNull("email", email)
    .putIfNotNull("password", password)
    .putIfNotNull("predictedWinner", predictedWinner)
    .putIfNotNull("score", score)
    .putIfNotNull("admin", admin)
    .putIfNotNull("joined", mapOf(Pair("\$date", joined?.toInstant(UTC).toString())))
    .putIfNotNull("adFree", adFree)
    .putIfNotNull("signature", signature)

  fun toUser() = User(
    id = this.id.toHexString(),
    firstName = this.firstName,
    surname = this.surname,
    email = this.email,
    password = this.password,
    predictedWinner = this.predictedWinner,
    score = this.score,
    joined = this.joined,
    admin = this.admin,
    adFree = this.adFree
  )
}

fun SignUpUser.toUserEntity(): UserEntity {
  return UserEntity(
    id = ObjectId(),
    firstName = this.firstName,
    surname = this.surname,
    email = this.email,
    password = this.password,
    predictedWinner = this.predictedWinner,
    score = 0,
    joined = LocalDateTime.now())
}

fun JsonObject.toUserEntity(): UserEntity {
  if (!ObjectId.isValid(this.getString("_id"))) {
    throw IllegalArgumentException("Invalid id")
  }

  return UserEntity(
    id = ObjectId(this.getString("_id")),
    firstName = this.getString("firstName"),
    surname = this.getString("surname"),
    email = this.getString("email"),
    password = this.getString("password"),
    predictedWinner = this.getString("predictedWinner"),
    score = this.getInteger("score"),
    admin = this.getBoolean("admin"),
    adFree = this.getBoolean("adFree"),
    joined = this.getLocalDateTime("joined"),
    signature = this.getString("signature")
  )
}

fun User.toUserEntity(): UserEntity {
  if (!ObjectId.isValid(this.id)) {
    throw IllegalArgumentException("Invalid id")
  }

  return UserEntity(
    id = ObjectId(this.id),
    firstName = this.firstName,
    surname = this.surname,
    email = this.email,
    password = this.password,
    predictedWinner = this.predictedWinner,
    score = this.score,
    joined = this.joined,
    admin = this.admin,
    adFree = this.adFree
  )
}
