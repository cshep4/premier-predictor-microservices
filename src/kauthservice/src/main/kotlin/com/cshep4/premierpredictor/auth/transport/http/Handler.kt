package com.cshep4.premierpredictor.auth.transport.http

import com.cshep4.premierpredictor.auth.html.ResetPassword.buildResetPasswordForm
import com.cshep4.premierpredictor.auth.html.ResetPasswordFailed.buildResetPasswordFailedForm
import com.cshep4.premierpredictor.auth.html.ResetPasswordSuccess.buildResetPasswordSuccessForm
import com.cshep4.premierpredictor.auth.model.InitiatePasswordResetRequest
import com.cshep4.premierpredictor.auth.model.LoginRequest
import com.cshep4.premierpredictor.auth.model.RegisterRequest
import com.cshep4.premierpredictor.auth.model.ResetPasswordForm
import com.cshep4.premierpredictor.auth.result.InitiatePasswordResetResult
import com.cshep4.premierpredictor.auth.result.LoginResult
import com.cshep4.premierpredictor.auth.result.LoginResult.Companion.PASSWORD_DOES_NOT_MATCH_ERROR
import com.cshep4.premierpredictor.auth.result.LoginResult.Companion.USER_NOT_FOUND_ERROR
import com.cshep4.premierpredictor.auth.result.RegisterResult
import com.cshep4.premierpredictor.auth.result.RegisterResult.Companion.EMAIL_ALREADY_EXISTS_ERROR
import com.cshep4.premierpredictor.auth.result.ResetPasswordResult
import com.cshep4.premierpredictor.auth.service.AuthService
import com.cshep4.premierpredictor.auth.util.ResponseUtils.badRequest
import com.cshep4.premierpredictor.auth.util.ResponseUtils.conflict
import com.cshep4.premierpredictor.auth.util.ResponseUtils.internal
import com.cshep4.premierpredictor.auth.util.ResponseUtils.ok
import com.cshep4.premierpredictor.auth.util.ResponseUtils.unauthorized
import com.cshep4.premierpredictor.auth.util.StringUtils.isValidEmailAddress
import com.cshep4.premierpredictor.auth.util.StringUtils.isValidPassword
import org.jboss.logging.Logger
import org.jboss.resteasy.annotations.providers.multipart.MultipartForm
import javax.enterprise.inject.Default
import javax.inject.Inject
import javax.ws.rs.*
import javax.ws.rs.core.MediaType.*
import javax.ws.rs.core.Response

@Path("/")
@Produces(APPLICATION_JSON)
@Consumes(APPLICATION_JSON)
class Handler {
    val logger = Logger.getLogger(Handler::class.java)

    @Inject
    @field: Default
    lateinit var authService: AuthService

    @POST
    @Path("/login")
    fun login(req: LoginRequest): Response {
        when {
            req.email.isEmpty() -> return badRequest("email is empty")
            req.password.isEmpty() -> return badRequest("password is empty")
        }

        return when (val res = authService.login(req.email, req.password)) {
            is LoginResult.Success -> ok(res)
            PASSWORD_DOES_NOT_MATCH_ERROR -> return unauthorized("password does not match")
            USER_NOT_FOUND_ERROR -> return unauthorized("user not found")
            is LoginResult.Error -> {
                logger.errorf(res.cause, "login_error: %s", res.message)
                unauthorized("could not login")
            }
        }
    }

    @POST
    @Path("/sign-up")
    fun register(req: RegisterRequest): Response {
        when {
            req.firstName.isEmpty() -> return badRequest("first name is empty")
            req.surname.isEmpty() -> return badRequest("surname is empty")
            req.email.isEmpty() -> return badRequest("email is empty")
            req.password.isEmpty() -> return badRequest("password is empty")
            req.confirmation.isEmpty() -> return badRequest("confirmation is empty")
            req.predictedWinner.isEmpty() -> return badRequest("predicted winner is empty")
            !req.email.isValidEmailAddress() -> return badRequest("email address is invalid")
            !req.password.isValidPassword() -> return badRequest("password is invalid")
            req.password != req.confirmation -> return badRequest("password and confirmation do not match")
        }

        return when (val res = authService.register(req)) {
            is RegisterResult.Success -> ok(res)
            EMAIL_ALREADY_EXISTS_ERROR -> conflict(EMAIL_ALREADY_EXISTS_ERROR.message)
            is RegisterResult.Error -> {
                logger.errorf(res.cause, "register_error: %s", res.message)
                internal("could not register")
            }
        }
    }

    @POST
    @Path("/initiate-password-reset")
    fun initiatePasswordReset(req: InitiatePasswordResetRequest): Response {
        when {
            req.email.isEmpty() -> return badRequest("email is empty")
            !req.email.isValidEmailAddress() -> return badRequest("email address is invalid")
        }

        when (val res = authService.initiatePasswordReset(req.email)) {
            InitiatePasswordResetResult.USER_NOT_FOUND_ERROR -> return badRequest("user not found")
            is InitiatePasswordResetResult.Error -> {
                logger.errorf(res.cause, "initiate_password_reset_error: %s", res.message)
                return internal("could not initiate password reset")
            }
        }

        return ok()
    }

    @POST
    @Path("/reset-password")
    @Consumes(MULTIPART_FORM_DATA)
    @Produces(TEXT_HTML)
    fun resetPassword(@MultipartForm req: ResetPasswordForm): String {
        when {
            req.email.isEmpty() -> return buildResetPasswordFailedForm()
            req.signature.isEmpty() -> return buildResetPasswordFailedForm()
            req.password.isEmpty() -> return buildResetPasswordFailedForm("password cannot be blank")
            req.confirmation.isEmpty() -> return buildResetPasswordFailedForm("confirmation cannot be blank")
            !req.password.isValidPassword() -> return buildResetPasswordFailedForm("password is invalid")
            req.password != req.confirmation -> return buildResetPasswordFailedForm("password and confirmation do not match")
        }

        when (val res = authService.resetPassword(req.toResetPasswordRequest())) {
            is ResetPasswordResult.Error -> {
                logger.errorf(res.cause, "reset_password_error: %s", res.message)
                return buildResetPasswordFailedForm()
            }
        }

        return buildResetPasswordSuccessForm()
    }

    @GET
    @Path("/reset-password")
    @Produces(TEXT_HTML)
    fun resetPassword(): String = buildResetPasswordForm()
}