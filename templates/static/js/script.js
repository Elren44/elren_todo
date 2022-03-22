var menu_home = document.querySelector(".navbar_home")
if (window.location.href.indexOf("todo") > -1) {
    var menu_todo = document.querySelector(".navbar_todo");
    menu_todo.classList.add("active");
    menu_home.classList.remove("active");
}

if (window.location.href.indexOf("signup") > -1) {
    var menu_signup = document.querySelector(".navbar_signup");
    menu_signup.classList.add("active");
    menu_home.classList.remove("active");
}

if (window.location.href.indexOf("login") > -1) {
    var menu_login = document.querySelector(".navbar_login");
    menu_login.classList.add("active");
    menu_home.classList.remove("active");
}
