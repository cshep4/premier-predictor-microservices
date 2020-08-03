package com.cshep4.premierpredictor.auth.util

import io.grpc.Metadata
import io.grpc.Metadata.ASCII_STRING_MARSHALLER
import io.grpc.Metadata.Key.of
import io.grpc.stub.MetadataUtils

object GrpcUtils {
    fun <T : io.grpc.stub.AbstractStub<T>> T.withMetadata(key: String, value: String): T {
        val m = Metadata()
        m.put(of(key, ASCII_STRING_MARSHALLER), value)

        return MetadataUtils.attachHeaders(this, m)
    }
}