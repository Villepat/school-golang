let commentForm = document.getElementById("post-comment");
commentForm.addEventListener("submit", function (e) {
    e.preventDefault();
    postComment();
});

async function postComment() {
    let comment = document.getElementById("comment-text").value;
    let fd = new FormData();
    fd.append("content", comment);

    let url = window.location.href;
    let res = await fetch(url, {
        method: "POST",
        body: fd,
    });
    if (res.status == 200) {
        location.reload();
    } else {
        alert("Failed to post comment");
    }
}

let delPost = document.getElementById("delete-post");
delPost.addEventListener("click", function (e) {
    e.preventDefault();
    deletePost();
});
async function deletePost() {
    let url = window.location.href;
    let res = await fetch(url, {
        method: "DELETE",
    });
    if (res.status == 200) {
        alert("Post deleted successfully");
        //wait for 1 second and redirect to home page:
        setTimeout(function () {
            window.location.href = "/";
        }, 1000);
        
    } else {
        alert("Failed to delete post");
    }
}


