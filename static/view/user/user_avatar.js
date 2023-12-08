function UserAvatar (userInfo = UserInfo(), iknowToken = IknowToken()) {
  const user = userInfo.json()

  const logout = () => {
    userInfo.remove()
    iknowToken.remove()
    location.reload()
  }
  const userDropList = document.querySelector('#userDroplist')
  if (user !== null) {
    userDropList.innerHTML = LoggedinListHTML(user.Name)
    document.querySelector('#logout').onclick = () => logout()
  } else {
    userDropList.innerHTML = DefaultListHTML()
  }

  htmx.process(userDropList)
}

function LoggedinListHTML (name = '') {
  return /* html */`
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

function DefaultListHTML () {
  return /* html */`<ul class="dropdown-menu text-small shadow" aria-labelledby="dropdownUser">
        <li><button class="dropdown-item" hx-trigger="click" hx-swap="innerHTML" hx-get="/static/view/login/login.html"
                hx-target="body">登入</button></li>
        <li>
            <hr class="dropdown-divider">
        </li>
        <li><button class="dropdown-item" hx-trigger="click" hx-swap="innerHTML" hx-get="/static/view/register/register.html"
                hx-target="body">註冊</button></li>
    </ul>`
}
