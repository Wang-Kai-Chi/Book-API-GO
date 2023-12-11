setTimeout(() => Login(UserInfo(), IknowToken()), 50)
function Login (userInfo = UserInfo(), iknowtoken = IknowToken()) {
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
        console.log('user info incorrect')
        return d.then(Promise.reject.bind(Promise))
      }
    }).then(data => {
      iknowtoken.set(JSON.stringify(data))
    }).catch(err => console.log(err))
      .then(() => window.location.href = '/')
  }

  const submit = () => {
    const email = document.querySelector('#email')
    const password = document.querySelector('#password')
    const user = User()

    user.this().Email = email.value
    user.this().Password = password.value

    fetch('/api/v1/user/login', {
      method: 'POST',
      body: JSON.stringify(user.this()),
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => {
      const d = res.json()
      if (res.status === 200) {
        return d
      } else {
        alert('電子郵件或密碼錯誤')
        return d.then(Promise.reject.bind(Promise))
      }
    }).then(data => {
      data.Password = ''
      userInfo.set(JSON.stringify(data))

      data.Password = user.this().Password
      getToken(JSON.stringify(data))
    }).catch(err => console.log(err))
  }

  document.querySelector('#submit').onclick = () => submit()
}
