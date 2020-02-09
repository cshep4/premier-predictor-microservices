package com.cshep4.premierpredictor.auth.repository

import com.cshep4.premierpredictor.auth.config.MongoConfig
import com.cshep4.premierpredictor.auth.config.MongoConfig.Index
import com.cshep4.premierpredictor.auth.entity.toUserEntity
import com.cshep4.premierpredictor.auth.exception.InternalException
import com.cshep4.premierpredictor.auth.model.SignUpUser
import com.cshep4.premierpredictor.auth.result.GetByEmailResult
import com.cshep4.premierpredictor.auth.result.StoreUserResult
import com.cshep4.premierpredictor.auth.result.UpdateUserResult
import com.google.common.collect.ImmutableMap
import io.vertx.core.AsyncResult
import io.vertx.core.Future.failedFuture
import io.vertx.core.Future.succeededFuture
import io.vertx.core.Handler
import io.vertx.core.Vertx
import io.vertx.core.json.JsonObject
import io.vertx.ext.mongo.MongoClient
import io.vertx.kotlin.coroutines.awaitResult
import kotlinx.coroutines.runBlocking
import org.bson.types.ObjectId
import java.util.Collections.singletonMap

class Repository(private val mongoCfg: MongoConfig) {
  companion object {
    const val DATABASE = "user"
    const val COLLECTION = "user"
  }

  private lateinit var client: MongoClient

  fun init(vertx: Vertx, handler: Handler<AsyncResult<Unit>>): Repository {
    val config = JsonObject(
      mapOf(
        Pair("connection_string", mongoCfg.uri),
        Pair("db_name", DATABASE),
        Pair("useObjectId", true)
      )
    )
    client = MongoClient.createShared(vertx, config)

    MongoConfig.indexes.forEach { createIndex(it, handler) }

    return this
  }

  private fun createIndex(idx: Index, handler: Handler<AsyncResult<Unit>>) {
    client.createIndexWithOptions(COLLECTION, idx.fields, idx.opts) {
      if (!it.succeeded()) {
        handler.handle(failedFuture(it.cause()))
        return@createIndexWithOptions
      }

      handler.handle(succeededFuture())
    }
  }

  fun getByEmail(email: String): GetByEmailResult = runBlocking {
    val query = JsonObject(
      mapOf(
        Pair("email", email)
      )
    )

    try {
      val resp = awaitResult<JsonObject?> {
        client.findOne(
          COLLECTION,
          query,
          null,
          it
        )
      } ?: return@runBlocking GetByEmailResult.Error(
        message = "user not found for email: $email",
        cause = Exception("user not found for email: $email")
      )

      if (resp.isEmpty) {
        return@runBlocking GetByEmailResult.Error(
          message = "valid user not found for email: $email",
          cause = Exception("valid user not found for email: $email")
        )
      }

      return@runBlocking GetByEmailResult.Success(
        user = resp.toUserEntity()
          .toUser()
      )
    } catch (e: Exception) {
      return@runBlocking GetByEmailResult.Error(
        message = "error executing getByEmail query for email: $email",
        cause = InternalException(e)
      )
    }
  }

  fun storeUser(signUpUser: SignUpUser): StoreUserResult = runBlocking {
    val userEntity = signUpUser.toUserEntity().toJson()
    userEntity.remove("_id")

    try {
      val id = awaitResult<String> {
        client.insert(
          COLLECTION,
          userEntity,
          it
        )
      }

      return@runBlocking StoreUserResult.Success(
        id = id
      )
    } catch (e: Exception) {
      return@runBlocking StoreUserResult.Error(
        message = "error executing storeUser query",
        cause = InternalException(e)
      )
    }
  }

  fun updateUser(id: String, vararg args: Pair<String, Any>): UpdateUserResult = runBlocking {
    if (!ObjectId.isValid(id)) {
      return@runBlocking UpdateUserResult.Error(
        message = "invalid user id",
        cause = Exception("invalid user id")
      )
    }

    try {
      val res = awaitResult<JsonObject> {
        client.findOneAndUpdate(
          COLLECTION,
          JsonObject(mapOf(Pair("\$oid", id))),
          JsonObject(mapOf(*args)),
          it
        )
      }

      return@runBlocking UpdateUserResult.Success
    } catch (e: Exception) {
      return@runBlocking UpdateUserResult.Error(
        message = "error executing storeUser query",
        cause = InternalException(e)
      )
    }
  }

  fun close() = client.close()
}
