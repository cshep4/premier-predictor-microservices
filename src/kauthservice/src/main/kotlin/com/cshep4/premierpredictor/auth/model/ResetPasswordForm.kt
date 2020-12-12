package com.cshep4.premierpredictor.auth.model

import org.jboss.resteasy.annotations.providers.multipart.PartType
import javax.ws.rs.FormParam
import javax.ws.rs.core.MediaType.TEXT_PLAIN

class ResetPasswordForm() {
    @FormParam("email")
    @PartType(TEXT_PLAIN)
    var email: String = ""

    @FormParam("signature")
    @PartType(TEXT_PLAIN)
    var signature: String = ""

    @FormParam("password")
    @PartType(TEXT_PLAIN)
    var password: String = ""

    @FormParam("confirmation")
    @PartType(TEXT_PLAIN)
    var confirmation: String = ""

    fun toResetPasswordRequest(): ResetPasswordRequest {
        return ResetPasswordRequest(
                email = email,
                signature = signature,
                password = password,
                confirmation = confirmation
        )
    }
}
