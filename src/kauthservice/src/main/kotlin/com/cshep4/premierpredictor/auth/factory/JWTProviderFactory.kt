package com.cshep4.premierpredictor.auth.factory

import io.vertx.core.Vertx
import io.vertx.ext.auth.PubSecKeyOptions
import io.vertx.ext.auth.jwt.JWTAuth
import io.vertx.ext.auth.jwt.JWTAuthOptions

class JWTProviderFactory {
  companion object {
    fun create(vertx: Vertx, secret: String): JWTAuth {
      val pubSecKey = PubSecKeyOptions()
        .setAlgorithm("HS512")
        .setPublicKey(secret)
        .setSymmetric(true)

      val opts = JWTAuthOptions()
        .addPubSecKey(pubSecKey)

      return JWTAuth.create(vertx, opts)
    }
  }
}
