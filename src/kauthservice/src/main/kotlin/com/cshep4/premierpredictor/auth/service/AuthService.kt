package com.cshep4.premierpredictor.auth.service

import com.cshep4.premierpredictor.auth.email.Emailer
import com.cshep4.premierpredictor.auth.enum.Role
import com.cshep4.premierpredictor.auth.enum.Role.SERVICE
import com.cshep4.premierpredictor.auth.enum.Role.USER
import com.cshep4.premierpredictor.auth.exception.UserNotFoundException
import com.cshep4.premierpredictor.auth.hash.Hasher
import com.cshep4.premierpredictor.auth.model.CreateUserRequest
import com.cshep4.premierpredictor.auth.model.RegisterRequest
import com.cshep4.premierpredictor.auth.model.ResetPasswordRequest
import com.cshep4.premierpredictor.auth.model.SendEmailRequest
import com.cshep4.premierpredictor.auth.result.*
import com.cshep4.premierpredictor.auth.result.LoginResult.Companion.PASSWORD_DOES_NOT_MATCH_ERROR
import com.cshep4.premierpredictor.auth.result.RegisterResult.Companion.EMAIL_ALREADY_EXISTS_ERROR
import com.cshep4.premierpredictor.auth.result.ResetPasswordResult.Companion.SIGNATURE_DOES_NOT_MATCH_ERROR
import com.cshep4.premierpredictor.auth.token.Tokenizer
import com.cshep4.premierpredictor.auth.user.UserService
import com.cshep4.premierpredictor.auth.util.EmailUtils
import com.cshep4.premierpredictor.auth.html.ResetPasswordEmail
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
        return try {
            val user = userService.getByEmail(email)

            if (!hasher.match(password, user.password)) {
                return PASSWORD_DOES_NOT_MATCH_ERROR
            }

            LoginResult.Success(
                    id = user.id,
                    token = tokenizer.generateToken(user.id, USER)
            )
        } catch (e: Exception) {
            if (e is UserNotFoundException) {
                return LoginResult.USER_NOT_FOUND_ERROR
            }
            LoginResult.Error(
                    message = "could not login",
                    cause = e
            )
        }
    }

    fun register(req: RegisterRequest): RegisterResult {
        try {
            userService.getByEmail(req.email)
            return EMAIL_ALREADY_EXISTS_ERROR
        } catch (e: Exception) {
            if (e !is UserNotFoundException) {
                return RegisterResult.Error(
                        message = "could not get user",
                        cause = e
                )
            }
        }

        return try {
            val createReq = CreateUserRequest(
                    firstName = req.firstName,
                    surname = req.surname,
                    email = req.email,
                    password = hasher.hash(req.password),
                    predictedWinner = req.predictedWinner
            )
            val id = userService.create(createReq)

            RegisterResult.Success(
                    id = id,
                    token = tokenizer.generateToken(id, USER)
            )
        } catch (e: Exception) {
            RegisterResult.Error(
                    message = "could not register",
                    cause = e
            )
        }
    }

    fun validate(token: String, audience: String, role: Role): ValidateTokenResult = try {
        tokenizer.validateToken(token, audience, role)

        ValidateTokenResult.Success
    } catch (e: Exception) {
        ValidateTokenResult.Error(
                message = "could not verify token",
                cause = e
        )
    }

    fun issueServiceToken(audience: String): String {
        return tokenizer.generateToken(audience, SERVICE)
    }

    fun initiatePasswordReset(email: String): InitiatePasswordResetResult {
        return try {
            val user = userService.getByEmail(email)

            val signature = tokenizer.generateSignature()

            userService.updateSignature(user.id, signature)

            val emailReq = SendEmailRequest(
                    sender = "Premier Predictor",
                    recipient = "${user.firstName} ${user.surname}",
                    senderEmail = "shepapps4@gmail.com",
                    recipientEmail = email,
                    subject = "Premier Predictor Password Reset",
                    content = EmailUtils.buildResetPasswordEmail(email, user.firstName, signature),
                    htmlContent = ResetPasswordEmail.buildResetPasswordEmail(email, user.firstName, signature)
            )
            emailer.send(emailReq)

            InitiatePasswordResetResult.Success
        } catch (e: Exception) {
            if (e is UserNotFoundException) {
                return InitiatePasswordResetResult.USER_NOT_FOUND_ERROR
            }

            InitiatePasswordResetResult.Error(
                    message = "could not initiate password reset",
                    cause = e
            )
        }
    }

    fun resetPassword(req: ResetPasswordRequest): ResetPasswordResult {
        return try {
            tokenizer.validateSignature(req.signature)

            val user = userService.getByEmail(req.email)

            if (!user.signature.equals(req.signature)) {
                return SIGNATURE_DOES_NOT_MATCH_ERROR
            }

            val password = hasher.hash(req.password)

            userService.updatePassword(user.id, password)

            ResetPasswordResult.Success
        } catch (e: Exception) {
            if (e is UserNotFoundException) {
                return ResetPasswordResult.USER_NOT_FOUND_ERROR
            }

            ResetPasswordResult.Error(
                    message = "could not reset password",
                    cause = e
            )
        }
    }
}