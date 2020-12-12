package com.cshep4.premierpredictor.auth.transport.http.health

import javax.ws.rs.GET
import javax.ws.rs.Path
import javax.ws.rs.Produces
import javax.ws.rs.core.MediaType
import javax.ws.rs.core.Response

@Path("/health")
class Handler {
    @GET
    @Produces(MediaType.APPLICATION_JSON)
    fun health(): Response = Response.ok()
            .build()
}