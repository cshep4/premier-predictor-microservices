package com.cshep4.premierpredictor.auth.transport.grpc

import com.cshep4.premierpredictor.auth.*
import com.cshep4.premierpredictor.auth.MutinyAuthServiceGrpc.AuthServiceImplBase
import com.cshep4.premierpredictor.auth.Role.ROLE_INVALID
import com.cshep4.premierpredictor.auth.enum.toRole
import com.cshep4.premierpredictor.auth.result.*
import com.cshep4.premierpredictor.auth.result.RegisterResult.Companion.EMAIL_ALREADY_EXISTS_ERROR
import com.cshep4.premierpredictor.auth.service.AuthService
import com.cshep4.premierpredictor.auth.util.StringUtils.isValidEmailAddress
import com.cshep4.premierpredictor.auth.util.StringUtils.isValidPassword
import com.cshep4.premierpredictor.request.LoginRequest
import com.google.protobuf.GeneratedMessageV3
import io.grpc.Status
import io.grpc.Status.*
import io.grpc.StatusException
import io.smallrye.mutiny.Uni
import org.jboss.logging.Logger
import javax.enterprise.inject.Default
import javax.inject.Inject
import javax.inject.Singleton
import com.cshep4.premierpredictor.auth.model.RegisterRequest as RegisterUser
import com.cshep4.premierpredictor.auth.model.ResetPasswordRequest as ResetRequest

@Singleton
class Server : AuthServiceImplBase() {
    val logger = Logger.getLogger(Server::class.java)

    @Inject
    @field: Default
    lateinit var authService: AuthService

    override fun login(req: LoginRequest): Uni<LoginResponse> {
        when {
            req.email.isEmpty() -> throw INVALID_ARGUMENT.error("email is empty")
            req.password.isEmpty() -> throw INVALID_ARGUMENT.error("password is empty")
        }

        val res = when (val res = authService.login(req.email, req.password)) {
            is LoginResult.Success -> res
            is LoginResult.Error -> {
                if (res.internal) {
                    logger.errorf(res.cause, "login_error: %s", res.message)
                }
                throw UNAUTHENTICATED.error(res.message, res.cause)
            }
        }


        return LoginResponse.newBuilder()
                .setId(res.id)
                .setToken(res.token)
                .build()
                .toUni()
    }

    override fun register(req: RegisterRequest): Uni<RegisterResponse> {
        when {
            req.firstName.isEmpty() -> throw INVALID_ARGUMENT.error("first name is empty")
            req.surname.isEmpty() -> throw INVALID_ARGUMENT.error("surname is empty")
            req.email.isEmpty() -> throw INVALID_ARGUMENT.error("email is empty")
            req.password.isEmpty() -> throw INVALID_ARGUMENT.error("password is empty")
            req.confirmation.isEmpty() -> throw INVALID_ARGUMENT.error("confirmation is empty")
            req.predictedWinner.isEmpty() -> throw INVALID_ARGUMENT.error("predicted winner is empty")
            !req.email.isValidEmailAddress() -> throw INVALID_ARGUMENT.error("email address is invalid")
            !req.password.isValidPassword() -> throw INVALID_ARGUMENT.error("password is invalid")
            req.password != req.confirmation -> throw INVALID_ARGUMENT.error("password and confirmation do not match")
        }

        val ru = RegisterUser(
                firstName = req.firstName,
                surname = req.surname,
                email = req.email,
                password = req.password,
                predictedWinner = req.predictedWinner
        )

        val res = when (val res = authService.register(ru)) {
            is RegisterResult.Success -> res
            EMAIL_ALREADY_EXISTS_ERROR -> throw ALREADY_EXISTS.error(EMAIL_ALREADY_EXISTS_ERROR.message)
            is RegisterResult.Error -> {
                logger.errorf(res.cause, "register_error: %s", res.message)
                throw INTERNAL.error(res.message, res.cause)
            }
        }

        return RegisterResponse.newBuilder()
                .setId(res.id)
                .setToken(res.token)
                .build()
                .toUni()
    }

    override fun validate(req: ValidateRequest): Uni<ValidateResponse> {
        when {
            req.token.isEmpty() -> throw INVALID_ARGUMENT.error("token is empty")
            req.role == ROLE_INVALID -> throw INVALID_ARGUMENT.error("role is invalid")
        }

        when (val res = authService.validate(req.token, req.audience, req.role.toRole())) {
            is ValidateTokenResult.Error -> {
                if (res.internal) {
                    logger.errorf(res.cause, "validate_error: %s", res.message)
                }
                throw UNAUTHENTICATED.error(res.message, res.cause)
            }
        }

        return ValidateResponse.getDefaultInstance().toUni()
    }

    override fun issueServiceToken(req: IssueServiceTokenRequest): Uni<IssueServiceTokenResponse> {
        val token = authService.issueServiceToken(req.audience)
        return IssueServiceTokenResponse.newBuilder()
                .setToken(token)
                .build()
                .toUni()
    }

    override fun initiatePasswordReset(req: InitiatePasswordResetRequest): Uni<InitiatePasswordResetResponse> {
        when {
            req.email.isEmpty() -> throw INVALID_ARGUMENT.error("email is empty")
            !req.email.isValidEmailAddress() -> throw INVALID_ARGUMENT.error("email address is invalid")
        }

        when (val res = authService.initiatePasswordReset(req.email)) {
            is InitiatePasswordResetResult.Error -> {
                if (res.internal) {
                    logger.errorf(res.cause, "initiate_password_reset_error: %s", res.message)
                }
                throw INTERNAL.error(res.message, res.cause)
            }
        }

        return InitiatePasswordResetResponse.getDefaultInstance().toUni()
    }

    override fun resetPassword(req: ResetPasswordRequest): Uni<ResetPasswordResponse> {
        when {
            req.email.isEmpty() -> throw INVALID_ARGUMENT.error("email is empty")
            req.signature.isEmpty() -> throw INVALID_ARGUMENT.error("signature is empty")
            req.password.isEmpty() -> throw INVALID_ARGUMENT.error("password is empty")
            req.confirmation.isEmpty() -> throw INVALID_ARGUMENT.error("confirmation is empty")
            !req.password.isValidPassword() -> throw INVALID_ARGUMENT.error("password is invalid")
            req.password != req.confirmation -> throw INVALID_ARGUMENT.error("password and confirmation do not match")
        }

        val rr = ResetRequest(
                email = req.email,
                signature = req.signature,
                password = req.password
        )

        when (val res = authService.resetPassword(rr)) {
            is ResetPasswordResult.Error -> {
                if (res.internal) {
                    logger.errorf(res.cause, "reset_password_error: %s", res.message)
                }
                throw INTERNAL.error(res.message, res.cause)
            }
        }

        return ResetPasswordResponse.getDefaultInstance().toUni()
    }

    fun <T : GeneratedMessageV3> T.toUni(): Uni<T> = Uni.createFrom().item { this }

    fun Status.error(description: String, cause: Throwable?): StatusException = this
            .withDescription(description)
            .withCause(cause)
            .asException()

    fun Status.error(description: String): StatusException = this
            .withDescription(description)
            .asException()
}