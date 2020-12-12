package com.cshep4.premierpredictor.config

import com.mongodb.ConnectionString
import com.mongodb.MongoClientSettings
import com.mongodb.client.MongoClient
import com.mongodb.client.MongoClients
import org.springframework.beans.factory.annotation.Value
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.bson.codecs.configuration.CodecRegistries.fromProviders
import org.bson.codecs.configuration.CodecRegistries.fromRegistries
import org.bson.codecs.pojo.PojoCodecProvider
import org.springframework.util.StringUtils.isEmpty

@Configuration
class MongoConfig {
    companion object {
        const val ID = "_id"
    }

    @Value("\${MONGO_SCHEME}")
    private val mongoScheme: String? = null

    @Value("\${MONGO_HOST}")
    private val mongoHost: String? = null

    @Value("\${MONGO_USERNAME}")
    private val mongoUsername: String? = null

    @Value("\${MONGO_PASSWORD}")
    private val mongoPassword: String? = null

    @Value("\${MONGO_PORT}")
    private val mongoPort: String? = null

    @Bean
    fun mongo(): MongoClient {
        var mongoUri = if (mongoUsername == "" || mongoPassword == "") {
            "$mongoScheme://$mongoHost"
        } else {
            "$mongoScheme://$mongoUsername:$mongoPassword@$mongoHost"
        }

        if (!isEmpty(mongoPort)) {
            mongoUri += ":$mongoPort"
        }

        mongoUri += "/?retryWrites=true"

        val pojoCodecRegistry = fromRegistries(
                MongoClientSettings.getDefaultCodecRegistry(),
                fromProviders(
                        PojoCodecProvider.builder()
                                .automatic(true)
                                .build()
                )
        )

        val settings = MongoClientSettings.builder()
                .codecRegistry(pojoCodecRegistry)
                .applyConnectionString(ConnectionString(mongoUri))
                .build()

        return MongoClients.create(settings)
    }
}