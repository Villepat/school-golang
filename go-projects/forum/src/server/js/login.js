let loginForm = document.getElementById("loginForm");
loginForm.addEventListener("submit", function (e) {
    e.preventDefault();
    clickme();
});

async function clickme() {
    let url = "/login";
    // get data from each input by id

    let username = document.getElementById("username").value;
    let pass = document.getElementById("password").value;
    // create payload as FormData
    const formData = new FormData();
    // add fields to that payload

    formData.append("username", username);
    formData.append("password", pass);

    // make POST request
    res = await fetch(url, {
        method: "POST",
        body: formData,
    });
    console.log(res);
    if (res.status == 200) {
        window.location.replace("/");
    } else {
        document.querySelector(".errorNotification").style.display = "block";
    }
}
