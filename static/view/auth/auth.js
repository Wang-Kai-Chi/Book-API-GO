/* eslint-disable no-undef */
Auth(IknowToken(), UserInfo())

function Auth (iknowtoken = IknowToken(), userInfo = UserInfo()) {
  const User = () => {
    const user = {
      Id: '',
      Name: '',
      Email: '',
      Phone: '',
      Password: ''
    }
    return {
      this: () => { return user },
      keys: () => { return Object.keys(user) }
    }
  }

  const getToken = async (user) => {
    fetch('/api/v1/jwt/token', {
      method: 'POST',
      body: user,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => {
      const d = res.json()
      if (res.status === 200) {
        return d
      } else {
        alert('密碼錯誤')
        return d.then(Promise.reject.bind(Promise))
      }
    }).then(data => {
      iknowtoken.set(JSON.stringify(data))
      location.reload()
    }).catch(err => console.log(err))
  }

  const submit = () => {
    const password = document.querySelector('#password')
    let user = User().this()

    user = userInfo.json()
    user.Password = password.value

    getToken(JSON.stringify(user))
  }

  document.querySelector('#submit').onclick = () => submit()
}
