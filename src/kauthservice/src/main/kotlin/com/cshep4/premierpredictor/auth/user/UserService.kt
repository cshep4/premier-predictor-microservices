package com.cshep4.premierpredictor.auth.user

import com.cshep4.premierpredictor.auth.enum.Role.SERVICE
import com.cshep4.premierpredictor.auth.exception.UserNotFoundException
import com.cshep4.premierpredictor.auth.model.CreateUserRequest
import com.cshep4.premierpredictor.auth.model.User
import com.cshep4.premierpredictor.auth.token.Tokenizer
import com.cshep4.premierpredictor.auth.util.GrpcUtils.withMetadata
import com.cshep4.premierpredictor.user.CreateRequest
import com.cshep4.premierpredictor.user.GetUserByEmailRequest
import com.cshep4.premierpredictor.user.UpdatePasswordRequest
import com.cshep4.premierpredictor.user.UpdateSignatureRequest
import io.grpc.Status.Code.NOT_FOUND
import io.grpc.StatusRuntimeException
import io.grpc.stub.AbstractStub
import io.quarkus.grpc.runtime.annotations.GrpcService
import java.time.Duration.ofSeconds
import javax.enterprise.inject.Default
import javax.inject.Inject
import javax.inject.Singleton
import com.cshep4.premierpredictor.user.MutinyUserServiceGrpc.MutinyUserServiceStub as UserClient
import com.cshep4.premierpredictor.user.User as GrpcUser

@Singleton
class UserService {
    companion object {
        const val SERVICE_NAME = "user"
        const val AUTH_KEY = "token"
    }

    @Inject
    @GrpcService(SERVICE_NAME)
    lateinit var client: UserClient

    @Inject
    @field: Default
    lateinit var tokenizer: Tokenizer

    fun getByEmail(email: String): User {
        return try
            val req = GetUserByEmailRequest.newBuilder()
                    .setEmail(email)
                    .build()

            client.withAuth()
                    .getUserByEmail(req)
                    .await()
                    .atMost(ofSeconds(1))
                    .user
                    .toUser()
        } catch (e: Exception) {
            if (e is StatusRuntimeException && e.status.code == NOT_FOUND) {
                throw UserNotFoundException()
            }

            throw e
        }
    }

    fun updatePassword(id: String, password: String) {
        val req = UpdatePasswordRequest.newBuilder()
                .setId(id)
                .setPassword(password)
                .build()

        client.withAuth()
                .updatePassword(req)
                .await()
                .atMost(ofSeconds(1))
    }

    fun updateSignature(id: String, signature: String) {
        val req = UpdateSignatureRequest.newBuilder()
                .setId(id)
                .setSignature(signature)
                .build()

        client.withAuth()
                .updateSignature(req)
                .await()
                .atMost(ofSeconds(1))
    }

    fun create(registerReq: CreateUserRequest): String {
        val req = CreateRequest.newBuilder()
                .setFirstName(registerReq.firstName)
                .setSurname(registerReq.surname)
                .setEmail(registerReq.email)
                .setPassword(registerReq.password)
                .setPredictedWinner(registerReq.predictedWinner)
                .build()

        return client.withAuth()
                .create(req)
                .await()
                .atMost(ofSeconds(1))
                .id
    }


    fun GrpcUser.toUser(): User {
        return User(
                id = this.id,
                firstName = this.firstName,
                surname = this.surname,
                predictedWinner = this.predictedWinner,
                score = this.score,
                email = this.email,
                password = this.password,
                signature = this.signature
        )
    }

    fun <T : AbstractStub<T>> T.withAuth(): T {
        return this.withMetadata(
                key = AUTH_KEY,
                value = tokenizer.generateToken(SERVICE_NAME, SERVICE)
        )
    }
}