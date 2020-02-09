package com.cshep4.premierpredictor.auth.service

import com.cshep4.premierpredictor.auth.constant.Messages.COULD_NOT_RESET_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.SIGNATURE_EXPIRED_ERROR
import com.cshep4.premierpredictor.auth.email.Emailer
import com.cshep4.premierpredictor.auth.email.PasswordResetEmailBuilder
import com.cshep4.premierpredictor.auth.exception.SignatureNotValidException
import com.cshep4.premierpredictor.auth.model.ResetPasswordArgs
import com.cshep4.premierpredictor.auth.model.SignUpUser
import com.cshep4.premierpredictor.auth.model.User
import com.cshep4.premierpredictor.auth.repository.Repository
import com.cshep4.premierpredictor.auth.result.*
import com.cshep4.premierpredictor.auth.tokeniser.Tokenizer
import org.mindrot.jbcrypt.BCrypt
import org.mindrot.jbcrypt.BCrypt.gensalt

class Service(
  private val repository: Repository,
  private val emailer: Emailer,
  private val tokenizer: Tokenizer
) {
  companion object {
    const val LOGIN_SUBJECT = "/login"
    const val REGISTER_SUBJECT = "/sign-up"
  }

  fun login(email: String, password: String): LoginResult {
    return when (val result = repository.getByEmail(email.toLowerCase())) {
      is GetByEmailResult.Error -> LoginResult.Error(result.message, result.cause)
      is GetByEmailResult.Success -> checkLogin(result.user, password)
    }
  }

  private fun checkLogin(user: User, password: String): LoginResult {
    try {
      if (BCrypt.checkpw(password, user.password)) {
        val token = tokenizer.generateToken(user.email, LOGIN_SUBJECT)
        return LoginResult.Success(user.id, token)
      }
    } catch (e: Throwable) {
    }

    return LoginResult.Error(
      message = "invalid password",
      cause = Exception("invalid password")
    )
  }

  fun register(signUpUser: SignUpUser): RegisterResult {
    if (repository.getByEmail(signUpUser.email) is GetByEmailResult.Success) {
      return EMAIL_EXISTS_ERROR
    }

    signUpUser.password = BCrypt.hashpw(signUpUser.password, gensalt())

    return when (val result = repository.storeUser(signUpUser)) {
      is StoreUserResult.Error -> RegisterResult.Error(result.message, result.cause)
      is StoreUserResult.Success -> RegisterResult.Success(id = result.id, token = tokenizer.generateToken(signUpUser.email, REGISTER_SUBJECT))
    }
  }

  fun validate(token: String): ValidateTokenResult = tokenizer.validateToken(token)

  fun initiatePasswordReset(email: String): InitiatePasswordResetResult {
    val user = when (val result = repository.getByEmail(email.toLowerCase())) {
      is GetByEmailResult.Error -> return InitiatePasswordResetResult.Error(
        message = result.message,
        cause = result.cause
      )
      is GetByEmailResult.Success -> result.user
    }

    val signature = tokenizer.createSignature(email)

    when (val result = repository.updateUser(user.id, Pair("signature", signature))) {
      is UpdateUserResult.Error -> return InitiatePasswordResetResult.Error(
        message = result.message,
        cause = result.cause
      )
    }

    val emailArgs = PasswordResetEmailBuilder()
      .withSender()
      .withSubject()
      .withRecipientEmail(email)
      .withRecipient("${user.firstName} ${user.surname}")
      .withMessage(email, signature)
      .build()

    return when (val result = emailer.send(emailArgs)) {
      is SendEmailResult.Error -> InitiatePasswordResetResult.Error(result.message, result.cause)
      is SendEmailResult.Success -> InitiatePasswordResetResult.Success
    }
  }

  fun resetPassword(req: ResetPasswordArgs): ResetPasswordResult {
    when (val result = tokenizer.validateToken(req.signature)) {
      is ValidateTokenResult.Error -> return ResetPasswordResult.Error(
        message = SIGNATURE_EXPIRED_ERROR,
        cause = result.cause
      )
    }

    val user = when (val result = repository.getByEmail(req.email.toLowerCase())) {
      is GetByEmailResult.Error -> return ResetPasswordResult.Error(
        message = COULD_NOT_RESET_ERROR,
        cause = result.cause
      )
      is GetByEmailResult.Success -> result.user
    }

    if (req.signature != user.signature) {
      return ResetPasswordResult.Error(
        message = COULD_NOT_RESET_ERROR,
        cause = SignatureNotValidException("Signature does not match: ${req.signature}")
      )
    }

    val password = BCrypt.hashpw(req.password, gensalt())

    return when (val result = repository.updateUser(user.id, Pair("password", password))) {
      is UpdateUserResult.Error -> ResetPasswordResult.Error(message = COULD_NOT_RESET_ERROR, cause = result.cause)
      is UpdateUserResult.Success -> ResetPasswordResult.Success
    }
  }
}
