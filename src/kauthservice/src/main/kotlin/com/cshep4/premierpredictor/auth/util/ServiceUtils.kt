package com.cshep4.premierpredictor.auth.util

import io.quarkus.runtime.Quarkus.asyncExit
import java.lang.System.getenv

object ServiceUtils {
    fun shutdown(key: String): String {
        println("environment variable $key not set")
        asyncExit()
        return ""
    }

    fun getEnv(key: String): String {
        return getenv(key) ?: shutdown(key)
    }

    fun getEnv(key: String, default: String): String {
        return getenv(key) ?: default
    }
}
