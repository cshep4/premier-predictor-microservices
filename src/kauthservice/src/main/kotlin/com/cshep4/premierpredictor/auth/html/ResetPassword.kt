package com.cshep4.premierpredictor.auth.html

import com.cshep4.premierpredictor.auth.util.ServiceUtils

object ResetPassword {
    const val RESET_PASSWORD_PATH = "/reset-password"

    val url = ServiceUtils.getEnv(
            key = "SERVICE_ADDR",
            default = "http://localhost:8080"
    )

    fun buildResetPasswordForm(): String {
        return """
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Reset Password</title>
    <!-- Latest compiled and minified CSS -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css"
          integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
</head>
<body>
<style>
    * {
        box-sizing: border-box
    }

    /* Full-width input fields */
    input[type=text], input[type=password] {
        width: 100%;
        padding: 15px;
        margin: 5px 0 22px 0;
        display: inline-block;
        border: none;
        background: #f1f1f1;
    }

    input[type=text]:focus, input[type=password]:focus {
        background-color: #ddd;
        outline: none;
    }

    hr {
        border: 1px solid #f1f1f1;
        margin-bottom: 25px;
    }

    /* Set a style for all buttons */
    button {
        background-color: #af3434;
        color: white;
        padding: 14px 20px;
        margin: 8px 0;
        border: none;
        cursor: pointer;
        width: 100%;
        opacity: 0.9;
    }

    button:hover {
        opacity: 1;
    }

    /* Add padding to container elements */
    .container {
        padding: 16px;
    }

    /* Clear floats */
    .clearfix::after {
        content: "";
        clear: both;
        display: table;
    }

    /* Change styles for cancel button and signup button on extra small screens */
    @media screen and (max-width: 300px) {
        .cancelbtn, .signupbtn {
            width: 100%;
        }
    }
</style>

<form action="$url$RESET_PASSWORD_PATH" enctype="multipart/form-data" method="post" name="myForm"
      style="border:1px solid #ccc">
    <div class="container">
        <h1>Reset Password</h1>
        <hr>
        <div id="errorMessage"></div>
        <input type="hidden" id="signature" name="signature">
        <input type="hidden" id="email" name="email">

        <label for="password"><b>New Password</b></label>
        <input type="password" placeholder="Enter Password" id="password" name="password" required>

        <ul>
            <li id="criteria1">Between 6 and 20 characters.</li>
            <li id="criteria2">Contain an uppercase letter.</li>
            <li id="criteria3">Contain a lowercase letter.</li>
            <li id="criteria4">Contain a number.</li>
            <li id="criteria5">Passwords must match.</li>
        </ul>

        <label for="confirmation"><b>Confirm Password</b></label>
        <input type="password" placeholder="Confirm Password" id="confirmation" name="confirmation" required>


        <div class="clearfix">
            <button onclick="resetpassword(this.form, this.form.password, this.form.confirmation);" type="button">
                Reset Password
            </button>
        </div>
    </div>
</form>
<script>
    var queryString = window.location.search.slice(1);
    if (queryString) {
        // stuff after # is not part of query string, so get rid of it
        queryString = queryString.split('#')[0];

        // split our query string into its component parts
        var arr = queryString.split('&');

        for (var i = 0; i < arr.length; i++) {
            // separate the keys and the values
            var a = arr[i].split('=');

            if (a[0] === "email") {
                document.getElementById("email").value = a[1];
            }

            if (a[0] === "signature") {
                document.getElementById("signature").value = a[1];
            }
        }
    }

    function doesPasswordContainUppercaseLetters(password) {
        return (/[A-Z]/.test(password));
    }

    function doesPasswordContainLowercaseLetters(password) {
        return (/[a-z]/.test(password));
    }

    function doesPasswordContainNumbers(password) {
        return (/[0-9]/.test(password));
    }

    function isPasswordBetween6And20Characters(password) {
        return password.length >= 6 && password.length <= 20;
    }

    function doPasswordsMatch(password, conf) {
        return password == conf;
    }

    function updateColour(id, criteriaMet) {
        var colour = "red";
        if (criteriaMet) {
            var colour = "green";
        }
        document.getElementById(id).style.color = colour;
    }

    var inputBox = document.getElementById('password');
    inputBox.onkeyup = function () {
        checkCriteria();
    };

    function checkCriteria() {
        password = document.getElementById("password").value;
        conf = document.getElementById("confirmation").value;

        updateColour("criteria1", isPasswordBetween6And20Characters(password));
        updateColour("criteria2", doesPasswordContainUppercaseLetters(password));
        updateColour("criteria3", doesPasswordContainLowercaseLetters(password));
        updateColour("criteria4", doesPasswordContainNumbers(password));
        updateColour("criteria5", isPasswordBetween6And20Characters(password));
    }

    updateColour("criteria1", false);
    updateColour("criteria2", false);
    updateColour("criteria3", false);
    updateColour("criteria4", false);
    updateColour("criteria5", false);

    function resetpassword(form, password, conf) {
        var error = 0;
        // Make sure the password is long enough (at least 6 characters)
        if (password.value.length < 6 || password.value.length > 20) {
            document.getElementById("errorMessage").innerHTML = '<p style="color:red">Passwords must be between 6 and 20 characters.  Please try again</p>';
            error = 1;
            return;
        }

        var re = /(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{6,}/;
        if (!re.test(password.value)) {
            document.getElementById("errorMessage").innerHTML = '<p style="color:red">Passwords must contain at least one number, one lowercase and one uppercase letter.  Please try again</p>';
            error = 1;
            return;
        }

        if (password.value != conf.value) {
            document.getElementById("errorMessage").innerHTML = '<p style="color:red">Your password and confirmation do not match. Please try again</p>';
            error = 1;
            return;
        }

        if (error == 0) {
            // Finally submit the form.
            form.submit();
        }
    }
</script>
</body>
</html>
        """.trimIndent()
    }
}