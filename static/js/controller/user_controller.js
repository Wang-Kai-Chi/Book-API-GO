import UserInfo from '../localstorage/user_info.js'
import JwtController from './jwt_controller.js'

export default function UserController () {
  const addUser = (bodyStr) => {
    return fetch('/api/v1/user/insert', {
      method: 'POST',
      body: bodyStr,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => {
      const d = res.json()
      if (res.status === 200) {
        return d
      } else {
        console.log('Register failed')
        return d.then(Promise.reject.bind(Promise))
      }
    }).then(() => location.reload())
      .catch(err => console.log(err))
  }

  const login = (bodyStr) => {
    return fetch('/api/v1/user/login', {
      method: 'POST',
      body: bodyStr,
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
      const noPswUser = () => {
        const temp = data
        temp.Password = ''
        return JSON.stringify(temp)
      }

      UserInfo().set(noPswUser())

      JwtController().getToken(JSON.stringify(data))
        .then(() => location.reload())
    }).catch(err => console.log(err))
  }

  const authurize = (bodyStr) => {
    return fetch('/api/v1/user/auth', {
      method: 'POST',
      body: bodyStr,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => {
      const d = res.json()
      if (res.status === 200) {
        return d
      } else {
        return d.then(Promise.reject.bind(Promise))
      }
    }).then(data => {
      console.log(data)
    })
      .catch(err => console.log(err))
  }

  return {
    addUser: (bodyStr = '') => addUser(bodyStr),
    login: (bodyStr = '') => login(bodyStr),
    authurize: (bodyStr = '') => authurize(bodyStr)
  }
}
