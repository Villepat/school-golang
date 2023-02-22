let createPostForm = document.getElementById("createPostForm");
createPostForm.addEventListener("submit", function (e) {
    e.preventDefault();
    createPost();
});

async function createPost() {
    const postTitle = document.getElementById("title").value;
    const postContent = document.getElementById("content").value;
    const postTags = document.getElementById("tags").value;
    const fd = new FormData();
    fd.append("title", postTitle);
    fd.append("content", postContent);
    fd.append("tags", postTags);
    const res = await fetch("/create", {
        method: "POST",
        body: fd,
    });
    const messageDiv = document.querySelector(".post");
    const message = document.createElement("p");
    const data = await res.json();
    if (res.status === 200) {
        message.innerHTML = data.message;
        message.className = "successNotification";
        message.style.display = "block";
        messageDiv.appendChild(message);
        console.log(data);
    } else {
        message.innerHTML = data.message;
        message.className = "errorNotification";
        message.style.display = "block";
        messageDiv.appendChild(message);
        console.log(data);
    }
}
