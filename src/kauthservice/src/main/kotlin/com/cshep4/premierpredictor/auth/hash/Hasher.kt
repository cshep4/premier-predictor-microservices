package com.cshep4.premierpredictor.auth.hash

import com.cshep4.premierpredictor.auth.result.HashResult
import com.cshep4.premierpredictor.auth.result.MatchResult
import com.cshep4.premierpredictor.auth.result.MatchResult.Match
import com.cshep4.premierpredictor.auth.result.MatchResult.NoMatch
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import javax.enterprise.inject.Default
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class Hasher {
    @Inject
    @field: Default
    lateinit var bCryptPasswordEncoder: BCryptPasswordEncoder

    fun hash(plaintext: String): HashResult = try {
        HashResult.Success(
                hash = bCryptPasswordEncoder.encode(plaintext)
        )
    } catch (e: Exception) {
        HashResult.Error(
                message = "could not generate hash",
                cause = e
        )
    }

    fun match(plaintext: String, hash: String): MatchResult = try {
        when (bCryptPasswordEncoder.matches(plaintext, hash)) {
            true -> Match
            false -> NoMatch
        }
    } catch (e: Exception) {
        MatchResult.Error(
                message = "could not check hash match",
                cause = e
        )
    }
}