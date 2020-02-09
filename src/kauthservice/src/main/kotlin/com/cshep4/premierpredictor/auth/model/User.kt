package com.cshep4.premierpredictor.auth.model

import com.fasterxml.jackson.annotation.JsonProperty
import java.time.LocalDateTime

data class User(
  @JsonProperty("_id")
  val id: String,
  var firstName: String = "",
  var surname: String = "",
  var email: String,
  var password: String,
  var predictedWinner: String = "",
  var score: Int = 0,
  val joined: LocalDateTime? = null,
  var admin: Boolean = false,
  var adFree: Boolean = false,
  var signature: String? = null
)
