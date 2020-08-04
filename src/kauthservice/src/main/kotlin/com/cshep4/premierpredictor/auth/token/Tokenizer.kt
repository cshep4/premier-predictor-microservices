package com.cshep4.premierpredictor.auth.token

import com.cshep4.premierpredictor.auth.enum.Role
import com.cshep4.premierpredictor.auth.enum.Role.SERVICE
import com.cshep4.premierpredictor.auth.exception.InvalidTokenException
import com.cshep4.premierpredictor.auth.util.ServiceUtils.getEnv
import io.jsonwebtoken.JwtBuilder
import io.jsonwebtoken.Jwts
import io.jsonwebtoken.SignatureAlgorithm.HS512
import java.util.*
import javax.inject.Singleton

@Singleton
class Tokenizer {
    val secret = Base64.getEncoder().encodeToString(getEnv("JWT_SECRET").toByteArray())

    fun generateToken(audience: String, role: Role): String {
        return Jwts.builder()
                .setAudience(audience)
                .addClaims(mapOf(Pair("role", role)))
                .setExpiration(role)
                .signWith(HS512, secret)
                .compact()
    }

    fun validateToken(token: String, audience: String, role: Role) {
        val jwt = Jwts.parser()
                .setSigningKey(secret)
                .parseClaimsJws(token)
                .body

        if (audience != "" && jwt.audience != audience) {
            throw InvalidTokenException("audience does not match: expected audience: $audience, token audience: ${jwt.audience}")

        }

        if (jwt["role"] != role.toString()) {
            throw InvalidTokenException("role does not match: expected role: $role, token role: ${jwt["role"]}")
        }
    }

    fun generateSignature(): String {
        val calendar = Calendar.getInstance()
        calendar.add(Calendar.HOUR, 24)

        return Jwts.builder()
                .setExpiration(calendar.time)
                .signWith(HS512, secret)
                .compact()
    }

    fun validateSignature(signature: String) {
        Jwts.parser()
                .setSigningKey(secret)
                .parseClaimsJws(signature)
                .body
    }

    fun JwtBuilder.setExpiration(role: Role): JwtBuilder {
        if (role == SERVICE) {
            val calendar = Calendar.getInstance()
            calendar.add(Calendar.SECOND, 5)
            return this.setExpiration(calendar.time)
        }

        return this
    }
}