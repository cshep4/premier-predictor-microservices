package com.cshep4.premierpredictor.auth.result

sealed class MatchResult {
    object Match : MatchResult()
    object NoMatch : MatchResult()
    data class Error(val message: String, val cause: Exception? = null, val internal: Boolean = true) : MatchResult()
}
