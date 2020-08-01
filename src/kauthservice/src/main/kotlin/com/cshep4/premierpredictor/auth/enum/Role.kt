package com.cshep4.premierpredictor.auth.enum

import com.cshep4.premierpredictor.auth.exception.InvalidRoleException
import com.cshep4.premierpredictor.auth.Role as ProtoRole

enum class Role {
    USER,
    SERVICE,
    ADMIN;

    fun toProto(): com.cshep4.premierpredictor.auth.Role {
        return when (this) {
            USER -> ProtoRole.ROLE_USER
            SERVICE -> ProtoRole.ROLE_SERVICE
            ADMIN -> ProtoRole.ROLE_ADMIN
        }
    }
}

fun ProtoRole.toRole(): Role {
    return when (this) {
        ProtoRole.ROLE_USER -> Role.USER
        ProtoRole.ROLE_SERVICE -> Role.SERVICE
        ProtoRole.ROLE_ADMIN -> Role.ADMIN
        else -> throw InvalidRoleException()
    }
}

