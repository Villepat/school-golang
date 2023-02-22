let filterSelect = document.getElementById("filterByCreated");
filterSelect.addEventListener("change", function () {
    filterByCreationTime(parseInt(filterSelect.value));
});

async function filterByCreationTime(timeBracket) {
    let filterTime;
    switch (timeBracket) {
        case 0:
            filterTime = new Date("1970-01-01");
            break;
        case 1:
            // Show last 7 days
            filterTime = new Date() - 7 * 24 * 60 * 60 * 1000;
            break;
        case 2:
            // Show last 30 days
            filterTime = new Date() - 30 * 24 * 60 * 60 * 1000;
            break;
        case 3:
            // Show last 90 days
            filterTime = new Date() - 90 * 24 * 60 * 60 * 1000;
            break;
        case 4:
            // Show last minute
            filterTime = new Date() - 1 * 60 * 1000;
            break;
    }
    let url = "/creationdates";
    const res = await fetch(url);
    const data = await res.json();
    let filteredData = data.filter((item) => {
        return new Date(item.Creation) > filterTime;
    });
    document.querySelectorAll(".subject").forEach((subject) => {
        if (
            filteredData
                .map((post) => post.PostID)
                .includes(parseInt(subject.id))
        ) {
            subject.style.display = "block";
        } else {
            subject.style.display = "none";
        }
    });
}
