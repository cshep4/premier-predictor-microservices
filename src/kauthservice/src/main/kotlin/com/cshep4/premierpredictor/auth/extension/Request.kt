package com.cshep4.premierpredictor.auth.extension

import com.cshep4.premierpredictor.auth.RegisterRequest
import com.cshep4.premierpredictor.auth.ResetPasswordRequest
import com.cshep4.premierpredictor.auth.ValidateRequest
import com.cshep4.premierpredictor.auth.constant.Messages.COULD_NOT_RESET_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.PASSWORDS_DO_NOT_MATCH_ERROR
import com.cshep4.premierpredictor.auth.constant.Messages.PASSWORD_NOT_VALID_ERROR
import com.cshep4.premierpredictor.auth.model.ResetPasswordArgs
import com.cshep4.premierpredictor.auth.model.SignUpUser
import com.cshep4.premierpredictor.auth.result.ResetPasswordRequestValidationResult
import com.cshep4.premierpredictor.request.EmailRequest
import com.cshep4.premierpredictor.request.LoginRequest

fun RegisterRequest?.isValid() = !this?.firstName.isNullOrEmpty() &&
  !this!!.surname.isNullOrEmpty() &&
  !this.email.isNullOrEmpty() &&
  this.email.isValidEmailAddress() &&
  !this.password.isNullOrEmpty() &&
  this.password.isValidPassword() &&
  this.password == this.confirmation

fun RegisterRequest.toSignUpUser() = SignUpUser(
  firstName = this.firstName,
  surname = this.surname,
  email = this.email.toLowerCase(),
  password = this.password,
  confirmPassword = this.confirmation,
  predictedWinner = this.predictedWinner
)

fun LoginRequest?.isValid() = !this?.email.isNullOrEmpty() &&
  !this?.password.isNullOrEmpty()

fun ValidateRequest?.isValid() = !this?.token.isNullOrEmpty()

fun EmailRequest?.isValid() = !this?.email.isNullOrEmpty() && this!!.email.isValidEmailAddress()

fun ResetPasswordRequest?.validate() = when {
  this?.email.isNullOrEmpty() -> ResetPasswordRequestValidationResult.Error(COULD_NOT_RESET_ERROR)
  !this!!.email.isValidEmailAddress() -> ResetPasswordRequestValidationResult.Error(COULD_NOT_RESET_ERROR)
  this.password.isNullOrEmpty() -> ResetPasswordRequestValidationResult.Error(PASSWORD_NOT_VALID_ERROR)
  !this.password.isValidPassword() -> ResetPasswordRequestValidationResult.Error(PASSWORD_NOT_VALID_ERROR)
  this.password != this.confirmation -> ResetPasswordRequestValidationResult.Error(PASSWORDS_DO_NOT_MATCH_ERROR)
  else -> ResetPasswordRequestValidationResult.Success
}

fun ResetPasswordRequest.toResetPasswordArgs() = ResetPasswordArgs(
  email = this.email,
  signature = this.signature,
  password = this.password
)
