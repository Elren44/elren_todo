let id = (id) => document.getElementById(id);

let classes = (classes) => document.getElementsByClassName(classes);
var valid = 0

let
    email = id("email"),
    password = id("password"),
    password2 = id("password2"),
    form = id("form"),

    errorMsg = classes("error"),
    successIcon = classes("success-icon"),
    failureIcon = classes("failure-icon");


form.addEventListener("submit", (e) => {
    e.preventDefault();

    engine(email, 0, "Это поле не может быть пустым");
    engine(password, 1, "Это поле не может быть пустым");
    engine(password2, 2, "Это поле не может быть пустым");
    checkPasswordMatch(2, "Пароли не совпадают!");
    if (valid == 4) {
        form.submit();
    }
});

let engine = (id, serial, message) => {

    if (id.value.trim() === "") {
        errorMsg[serial].innerHTML = message;
        id.style.border = "2px solid red";

        // icons
        failureIcon[serial].style.opacity = "1";
        successIcon[serial].style.opacity = "0";
        valid = 0
    }

    else {
        errorMsg[serial].innerHTML = "";
        id.style.border = "2px solid green";

        // icons
        failureIcon[serial].style.opacity = "0";
        successIcon[serial].style.opacity = "1";
        valid = valid +1;
    }
}

let checkPasswordMatch = function (serial, message) {
    if (password.value !== password2.value) {
        errorMsg[serial].innerHTML = message;
        password2.style.border = "2px solid red";

        // icons
        failureIcon[serial].style.opacity = "1";
        successIcon[serial].style.opacity = "0";
        valid = 0;
    }
    else {
        valid = valid+1;
    }
}