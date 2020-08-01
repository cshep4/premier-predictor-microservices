package com.cshep4.premierpredictor.auth.user

import com.cshep4.premierpredictor.auth.enum.Role.SERVICE
import com.cshep4.premierpredictor.auth.model.CreateUserRequest
import com.cshep4.premierpredictor.auth.model.User
import com.cshep4.premierpredictor.auth.result.CreateUserResult
import com.cshep4.premierpredictor.auth.result.GetByEmailResult
import com.cshep4.premierpredictor.auth.result.GetByEmailResult.Companion.USER_NOT_FOUND_ERROR
import com.cshep4.premierpredictor.auth.result.UpdatePasswordResult
import com.cshep4.premierpredictor.auth.result.UpdateSignatureResult
import com.cshep4.premierpredictor.auth.token.Tokenizer
import com.cshep4.premierpredictor.auth.util.GrpcUtils.withMetadata
import com.cshep4.premierpredictor.request.EmailRequest
import com.cshep4.premierpredictor.user.CreateRequest
import com.cshep4.premierpredictor.user.UpdatePasswordRequest
import com.cshep4.premierpredictor.user.UpdateSignatureRequest
import io.grpc.Status.Code.NOT_FOUND
import io.grpc.StatusRuntimeException
import io.grpc.stub.AbstractBlockingStub
import io.quarkus.grpc.runtime.annotations.GrpcService
import javax.enterprise.inject.Default
import javax.inject.Inject
import javax.inject.Singleton
import com.cshep4.premierpredictor.user.User as UserResponse
import com.cshep4.premierpredictor.user.UserServiceGrpc.UserServiceBlockingStub as UserClient

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

    fun getByEmail(email: String): GetByEmailResult {
        return try {
            val req = EmailRequest.newBuilder()
                    .setEmail(email)
                    .build()

            GetByEmailResult.Success(
                    user = client.withAuth()
                            .getUserByEmail(req)
                            .toUser()
            )
        } catch (e: Exception) {
            if (e is StatusRuntimeException && e.status.code == NOT_FOUND) {
                return USER_NOT_FOUND_ERROR
            }

            GetByEmailResult.Error(
                    message = "could not get user by email",
                    cause = e,
                    internal = true
            )
        }
    }

    fun updatePassword(id: String, password: String): UpdatePasswordResult = try {
        val req = UpdatePasswordRequest.newBuilder()
                .setId(id)
                .setPassword(password)
                .build()

        client.withAuth()
                .updatePassword(req)

        UpdatePasswordResult.Success
    } catch (e: StatusRuntimeException) {
        UpdatePasswordResult.Error(
                message = "could not update password",
                cause = e,
                internal = true
        )
    }

    fun updateSignature(id: String, signature: String): UpdateSignatureResult = try {
        val req = UpdateSignatureRequest.newBuilder()
                .setId(id)
                .setSignature(signature)
                .build()

        client.withAuth()
                .updateSignature(req)

        UpdateSignatureResult.Success
    } catch (e: StatusRuntimeException) {
        UpdateSignatureResult.Error(
                message = "could not update signature",
                cause = e,
                internal = true
        )
    }

    fun create(registerReq: CreateUserRequest): CreateUserResult = try {
        val req = CreateRequest.newBuilder()
                .setFirstName(registerReq.firstName)
                .setSurname(registerReq.surname)
                .setEmail(registerReq.email)
                .setPassword(registerReq.password)
                .setPredictedWinner(registerReq.predictedWinner)
                .build()

        val res = client.withAuth()
                .create(req)

        CreateUserResult.Success(
                id = res.id
        )
    } catch (e: StatusRuntimeException) {
        CreateUserResult.Error(
                message = "could not create user",
                cause = e,
                internal = true
        )
    }


    fun UserResponse.toUser(): User {
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

    fun <T : AbstractBlockingStub<T>> T.withAuth(): T {
        return this.withMetadata(
                key = AUTH_KEY,
                value = tokenizer.generateToken(SERVICE_NAME, SERVICE)
        )
    }
}