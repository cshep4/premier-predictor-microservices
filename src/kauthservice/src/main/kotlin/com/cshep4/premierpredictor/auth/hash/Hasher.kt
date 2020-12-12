package com.cshep4.premierpredictor.auth.hash

import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import javax.enterprise.inject.Default
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class Hasher {
    @Inject
    @field: Default
    lateinit var bCryptPasswordEncoder: BCryptPasswordEncoder

    fun hash(plaintext: String): String = bCryptPasswordEncoder.encode(plaintext)

    fun match(plaintext: String, hash: String): Boolean = bCryptPasswordEncoder.matches(plaintext, hash)
}