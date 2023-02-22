let registrationForm = document.getElementById("registrationForm");
registrationForm.addEventListener("submit", function (e) {
    e.preventDefault();
    clickme();
});

async function clickme() {
    let url = "/register";
    // get data from each input by id
    let email = document.getElementById("email").value;
    let username = document.getElementById("username").value;
    let pass = document.getElementById("password").value;
    // create payload as FormData
    const formData = new FormData();
    // add fields to that payload
    formData.append("email", email);
    formData.append("username", username);
    formData.append("password", pass);

    // make POST request
    res = await fetch(url, {
        method: "POST",
        body: formData,
    });
    if (res.status == 200) {
        document.querySelector(".successNotification").style.display = "block";
    } else {
        document.querySelector(".errorNotification").style.display = "block";
    }
}
