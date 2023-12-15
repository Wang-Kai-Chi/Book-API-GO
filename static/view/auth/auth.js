import IknowToken from '../iknow_token.js'
import UserInfo from '../user_info.js'

Auth()

function Auth () {
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
      IknowToken().set(JSON.stringify(data))
      location.reload()
    }).catch(err => console.log(err))
  }

  const submit = () => {
    const password = document.querySelector('#password')
    let user = User().this()

    user = UserInfo().json()
    user.Password = password.value

    getToken(JSON.stringify(user))
  }

  document.querySelector('#submit').onclick = () => submit()
}
