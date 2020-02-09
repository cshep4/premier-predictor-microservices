package com.cshep4.premierpredictor.auth.model

data class SignUpUser(var firstName: String,
                      var surname: String,
                      var email: String,
                      var password: String,
                      var confirmPassword: String,
                      var predictedWinner: String)
