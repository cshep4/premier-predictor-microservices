package com.cshep4.premierpredictor.auth.config

import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.module.kotlin.KotlinModule
import javax.enterprise.context.ApplicationScoped
import javax.inject.Singleton
import javax.ws.rs.Produces

@ApplicationScoped
class CustomObjectMapperConfig {
    @Singleton
    @Produces
    fun objectMapper(): ObjectMapper {
        return ObjectMapper().registerModule(KotlinModule())
    }
}