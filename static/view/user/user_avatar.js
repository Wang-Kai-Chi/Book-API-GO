
function UserAvatar() {
    const logout = () => {
        localStorage.removeItem(USER_INFO_ID)
        location.reload()
    }

    const USER_INFO_ID = "userinfo"
    const userDropList = document.querySelector("#userDroplist")
    const user = JSON.parse(localStorage.getItem(USER_INFO_ID))

    if (user !== null) {
        userDropList.innerHTML = LoggedinListHTML(user.Name)
        document.querySelector("#logout").onclick = () => logout()
    } else {
        userDropList.innerHTML = DefaultListHTML()
    }

    htmx.process(userDropList)
}

function LoggedinListHTML(name = "") {
    return /*html*/`
    <ul class="dropdown-menu text-small shadow" aria-labelledby="dropdownUser">
        <li><button class="dropdown-item">${name}</button></li>
        <li><button class="dropdown-item">設定</button></li>
        <li>
            <hr class="dropdown-divider">
        </li>
        <li><a class="dropdown-item" id="logout">登出</a></li>
    </ul>
    `
}

function DefaultListHTML() {
    return /*html*/`
    <ul class="dropdown-menu text-small shadow" aria-labelledby="dropdownUser">
        <li><button class="dropdown-item" hx-trigger="click" hx-swap="innerHTML" hx-get="/static/view/login/login.html"
                hx-target="body">登入</button></li>
        <li>
            <hr class="dropdown-divider">
        </li>
        <li><a class="dropdown-item" href="#">註冊</a></li>
    </ul>
    `
}