package com.cshep4.premierpredictor.auth.extension

import com.google.protobuf.Empty
import io.grpc.Status
import io.grpc.Status.UNAUTHENTICATED
import io.grpc.StatusException
import io.vertx.core.Future

fun <T> Future<T>.ok(resp: T) {
  this.handle(Future.succeededFuture(resp))
}

fun Future<Empty>.ok() {
  this.handle(Future.succeededFuture(Empty.newBuilder().build()))
}

fun <T> Future<T>.unauthenticated() {
  this.handle(Future.failedFuture(StatusException(UNAUTHENTICATED)))
}

fun <T> Future<T>.error(e: Exception) {
  this.handle(Future.failedFuture(e))
}

fun <T> Future<T>.error(s: Status) {
  this.handle(Future.failedFuture(StatusException(s)))
}
