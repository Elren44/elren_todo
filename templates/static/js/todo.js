var btn = document.querySelector(".toggle_btn");
var divBody = document.querySelector(".showTasks");

btn.addEventListener("click", function (e) {
    divBody.classList.toggle("showTasks_show");
})