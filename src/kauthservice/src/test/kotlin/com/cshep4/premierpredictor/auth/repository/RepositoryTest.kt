package com.cshep4.premierpredictor.auth.repository

import com.cshep4.premierpredictor.auth.config.MongoConfig
import com.cshep4.premierpredictor.auth.exception.InternalException
import com.cshep4.premierpredictor.auth.repository.Repository.Companion.COLLECTION
import com.cshep4.premierpredictor.auth.repository.Repository.Companion.DATABASE
import com.cshep4.premierpredictor.auth.result.GetByEmailResult
import com.cshep4.premierpredictor.auth.result.UpdateUserResult
import com.fasterxml.jackson.module.kotlin.KotlinModule
import io.vertx.core.Vertx
import io.vertx.core.VertxOptions
import io.vertx.core.json.Json
import io.vertx.core.json.JsonObject
import io.vertx.ext.mongo.MongoClient
import io.vertx.ext.mongo.MongoClientDeleteResult
import io.vertx.kotlin.coroutines.awaitResult
import kotlinx.coroutines.runBlocking
import org.hamcrest.CoreMatchers.`is`
import org.hamcrest.MatcherAssert.assertThat
import org.hamcrest.core.IsInstanceOf
import org.junit.After
import org.junit.Before
import org.junit.Test


internal class RepositoryTest {
  companion object {
    const val EMAIL = "📨"
    const val PASSWORD = "🔑"
  }

  private val mongoCfg = MongoConfig(uri = "mongodb://localhost:27017")

  private val vertx = Vertx.vertx(VertxOptions().setBlockedThreadCheckInterval(1000 * 60 * 60))

  private lateinit var repository: Repository

  private lateinit var client: MongoClient

  @Before
  fun init() = runBlocking {
//    Json.mapper.registerModule(KotlinModule())

    awaitResult<Unit> {
      repository = Repository(mongoCfg)
        .init(vertx, it)
    }

    val config = JsonObject(
      mapOf(
        Pair("connection_string", "mongodb://localhost:27017"),
        Pair("db_name", DATABASE),
        Pair("useObjectId", true)
      )
    )
    client = MongoClient.createShared(vertx, config)
  }

  @Test
  fun `'getByEmail' returns error if email doesn't exist`() {
    val result = repository.getByEmail(EMAIL)

    assertThat(result, IsInstanceOf(GetByEmailResult.Error::class.java))

    val err = result as GetByEmailResult.Error

    assertThat(err.message, `is`("user not found for email: $EMAIL"))
    assertThat(err.cause, IsInstanceOf(Exception::class.java))
    assertThat(err.cause!!.message, `is`("user not found for email: $EMAIL"))
  }

  @Test
  fun `'getByEmail' returns error if user is in wrong format`() = runBlocking {
    val doc = JsonObject(
      mapOf(
        Pair("email", EMAIL),
        Pair("score", "invalid"),
        Pair("joined", "invalid"),
        Pair("admin", "invalid"),
        Pair("adFree", "invalid")
      )
    )

    awaitResult<String> { client.insert(COLLECTION, doc, it) }

    val result = repository.getByEmail(EMAIL)

    assertThat(result, IsInstanceOf(GetByEmailResult.Error::class.java))

    val err = result as GetByEmailResult.Error

    assertThat(err.message, `is`("error executing getByEmail query for email: $EMAIL"))
    assertThat(err.cause, IsInstanceOf(InternalException::class.java))
  }

  @Test
  fun `'getByEmail' returns user`() = runBlocking {
    val doc = JsonObject(
      mapOf(
        Pair("firstName", "first"),
        Pair("surname", "last"),
        Pair("email", EMAIL),
        Pair("password", PASSWORD),
        Pair("predictedWinner", "winner"),
        Pair("admin", false),
        Pair("adFree", false),
        Pair("score", 1)
      )
    )

    awaitResult<String> { client.insert(COLLECTION, doc, it) }

    val result = repository.getByEmail(EMAIL)

    assertThat(result, IsInstanceOf(GetByEmailResult.Success::class.java))

    val success = result as GetByEmailResult.Success

    assertThat(success.user.firstName, `is`("first"))
    assertThat(success.user.surname, `is`("last"))
    assertThat(success.user.email, `is`(EMAIL))
    assertThat(success.user.password, `is`(PASSWORD))
    assertThat(success.user.predictedWinner, `is`("winner"))
    assertThat(success.user.score, `is`(1))
  }

  @Test
  fun `'updateUser' updates specified properties for user`() = runBlocking {
    val doc = JsonObject(
      mapOf(
        Pair("firstName", "first"),
        Pair("surname", "last"),
        Pair("email", EMAIL),
        Pair("password", PASSWORD),
        Pair("predictedWinner", "winner"),
        Pair("admin", false),
        Pair("adFree", false),
        Pair("score", 1)
      )
    )

    val newName = "new name"

    val id = awaitResult<String> { client.insert(COLLECTION, doc, it) }

    val updateResult = repository.updateUser(id, Pair("firstName", newName))

    assertThat(updateResult, IsInstanceOf(UpdateUserResult.Success::class.java))

    val result = repository.getByEmail(EMAIL)

    assertThat(result, IsInstanceOf(GetByEmailResult.Success::class.java))

    val success = result as GetByEmailResult.Success

    assertThat(success.user.firstName, `is`(newName))
  }

  @After
  fun tearDown() = runBlocking {
    awaitResult<MongoClientDeleteResult> { client.removeDocuments(COLLECTION, JsonObject(), it) }

    client.close()
    repository.close()
  }
}
