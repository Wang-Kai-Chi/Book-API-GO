import UserInfo from '../localstorage/user_info.js'
import JwtController from './jwt_controller.js'
import ResponseHandler from '../request/response_handler.js'

export default function UserController () {
  const addUser = async (bodyStr) => {
    return fetch('/api/v1/user/insert', {
      method: 'POST',
      body: bodyStr,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => {
      ResponseHandler().run(res)
    }).catch(err => console.log(err))
  }

  const handleLogin = async (data) => {
    const noPswUser = (data) => {
      const temp = data
      temp.Password = ''
      return JSON.stringify(temp)
    }
    UserInfo().set(noPswUser(data))

    JwtController().getToken(JSON.stringify(data))
      .then(() => location.reload())
  }

  const login = async (bodyStr) => {
    return fetch('/api/v1/user/login', {
      method: 'POST',
      body: bodyStr,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => ResponseHandler().run(res))
      .then(data => {
        handleLogin(data)
      }).catch(err => console.log(err))
  }

  const authurize = async (bodyStr, success) => {
    return fetch('/api/v1/user/auth', {
      method: 'POST',
      body: bodyStr,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => {
      ResponseHandler().run(res, success)
    })
      .catch(err => console.log(err))
  }

  return {
    addUser: (bodyStr = '') => addUser(bodyStr),
    login: (bodyStr = '') => login(bodyStr),
    authurize: (bodyStr = '', success = () => {}) => authurize(bodyStr, success)
  }
}
