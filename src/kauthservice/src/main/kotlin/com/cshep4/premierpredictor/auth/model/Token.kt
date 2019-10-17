package com.cshep4.premierpredictor.auth.model

data class Token(val sub: String, val iss: String, val exp: Long)
