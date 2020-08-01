package com.cshep4.premierpredictor.auth.service

import com.cshep4.premierpredictor.auth.email.Emailer
import com.cshep4.premierpredictor.auth.enum.Role
import com.cshep4.premierpredictor.auth.enum.Role.SERVICE
import com.cshep4.premierpredictor.auth.enum.Role.USER
import com.cshep4.premierpredictor.auth.hash.Hasher
import com.cshep4.premierpredictor.auth.model.CreateUserRequest
import com.cshep4.premierpredictor.auth.model.RegisterRequest
import com.cshep4.premierpredictor.auth.model.ResetPasswordRequest
import com.cshep4.premierpredictor.auth.model.SendEmailRequest
import com.cshep4.premierpredictor.auth.result.*
import com.cshep4.premierpredictor.auth.result.GetByEmailResult.Companion.USER_NOT_FOUND_ERROR
import com.cshep4.premierpredictor.auth.result.LoginResult.Companion.PASSWORD_DOES_NOT_MATCH_ERROR
import com.cshep4.premierpredictor.auth.result.MatchResult.NoMatch
import com.cshep4.premierpredictor.auth.result.RegisterResult.Companion.EMAIL_ALREADY_EXISTS_ERROR
import com.cshep4.premierpredictor.auth.result.ResetPasswordResult.Companion.INVALID_SIGNATURE_ERROR
import com.cshep4.premierpredictor.auth.result.ResetPasswordResult.Companion.SIGNATURE_DOES_NOT_MATCH_ERROR
import com.cshep4.premierpredictor.auth.token.Tokenizer
import com.cshep4.premierpredictor.auth.user.UserService
import com.cshep4.premierpredictor.auth.util.EmailUtils.buildResetPasswordEmail
import javax.enterprise.inject.Default
import javax.inject.Inject
import javax.inject.Named
import javax.inject.Singleton

@Singleton
class AuthService {
    @Inject
    @field: Default
    lateinit var userService: UserService

    @Inject
    @field: Default
    lateinit var tokenizer: Tokenizer

    @Inject
    @field: Named("httpEmailer")
    lateinit var emailer: Emailer

    @Inject
    @field: Default
    lateinit var hasher: Hasher

    fun login(email: String, password: String): LoginResult {
        val user = when (val res = userService.getByEmail(email)) {
            is GetByEmailResult.Success -> res.user
            is GetByEmailResult.Error -> return LoginResult.Error(
                    message = res.message,
                    cause = res.cause,
                    internal = res.internal
            )
        }

        when (val res = hasher.match(password, user.password)) {
            is NoMatch -> return PASSWORD_DOES_NOT_MATCH_ERROR
            is MatchResult.Error -> return LoginResult.Error(
                    message = res.message,
                    cause = res.cause,
                    internal = res.internal
            )
        }
        return LoginResult.Success(
                id = user.id,
                token = tokenizer.generateToken(user.id, USER)
        )
    }

    fun register(req: RegisterRequest): RegisterResult {
        when (val res = userService.getByEmail(req.email)) {
            is GetByEmailResult.Success -> return EMAIL_ALREADY_EXISTS_ERROR
            USER_NOT_FOUND_ERROR -> {
            }
            is GetByEmailResult.Error -> return RegisterResult.Error(
                    message = res.message,
                    cause = res.cause,
                    internal = res.internal
            )
        }

        val hashedPassword = when (val res = hasher.hash(req.password)) {
            is HashResult.Success -> res.hash
            is HashResult.Error -> return RegisterResult.Error(
                    message = res.message,
                    cause = res.cause,
                    internal = res.internal
            )
        }

        val createReq = CreateUserRequest(
                firstName = req.firstName,
                surname = req.surname,
                email = req.email,
                password = hashedPassword,
                predictedWinner = req.predictedWinner
        )
        val id = when (val res = userService.create(createReq)) {
            is CreateUserResult.Success -> res.id
            is CreateUserResult.Error -> return RegisterResult.Error(
                    message = res.message,
                    cause = res.cause,
                    internal = res.internal
            )
        }

        return RegisterResult.Success(
                id = id,
                token = tokenizer.generateToken(id, USER)
        )
    }

    fun issueServiceToken(audience: String): String {
        return tokenizer.generateToken(audience, SERVICE)
    }

    fun validate(token: String, audience: String, role: Role): ValidateTokenResult {
        return tokenizer.validateToken(token, audience, role)
    }

    fun initiatePasswordReset(email: String): InitiatePasswordResetResult {
        val user = when (val res = userService.getByEmail(email)) {
            is GetByEmailResult.Success -> res.user
            is GetByEmailResult.Error -> return InitiatePasswordResetResult.Error(
                    message = res.message,
                    cause = res.cause,
                    internal = res.internal
            )
        }

        val signature = tokenizer.generateSignature()

        when (val res = userService.updateSignature(user.id, signature)) {
            is UpdateSignatureResult.Error -> return InitiatePasswordResetResult.Error(
                    message = res.message,
                    cause = res.cause,
                    internal = res.internal
            )
        }

        val emailReq = SendEmailRequest(
                sender = "Premier Predictor",
                recipient = "${user.firstName} ${user.surname}",
                senderEmail = "shepapps4@gmail.com",
                recipientEmail = email,
                subject = "Initiate Password Reset",
                content = buildResetPasswordEmail(email, user.firstName, signature)
        )
        when (val res = emailer.send(emailReq)) {
            is SendEmailResult.Error -> return InitiatePasswordResetResult.Error(
                    message = res.message,
                    cause = res.cause,
                    internal = res.internal
            )
        }

        return InitiatePasswordResetResult.Success
    }

    fun resetPassword(req: ResetPasswordRequest): ResetPasswordResult {
        when (tokenizer.validateSignature(req.signature)) {
            is ValidateSignatureResult.Error -> return INVALID_SIGNATURE_ERROR
        }

        val user = when (val res = userService.getByEmail(req.email)) {
            is GetByEmailResult.Success -> res.user
            is GetByEmailResult.Error -> return ResetPasswordResult.Error(
                    message = res.message,
                    cause = res.cause,
                    internal = res.internal
            )
        }

        if (!user.signature.equals(req.signature)) {
            return SIGNATURE_DOES_NOT_MATCH_ERROR
        }

        when (val res = userService.updatePassword(user.id, req.password)) {
            is UpdatePasswordResult.Error -> return ResetPasswordResult.Error(
                    message = res.message,
                    cause = res.cause,
                    internal = res.internal
            )
        }

        return ResetPasswordResult.Success
    }
}