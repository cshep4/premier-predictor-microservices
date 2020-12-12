package com.cshep4.premierpredictor.auth.config

import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import javax.enterprise.context.ApplicationScoped
import javax.enterprise.inject.Produces

@ApplicationScoped
class BcryptConfig {
    @Produces
    fun bCryptPasswordEncoder(): BCryptPasswordEncoder = BCryptPasswordEncoder(10)
}