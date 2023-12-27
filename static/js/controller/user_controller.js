import UserInfo from '../localstorage/user_info.js'
import JwtController from './jwt_controller.js'
import ResponseHandler from '../request/response_handler.js'

export default function UserController () {
  const addUser = (bodyStr) => {
    return fetch('/api/v1/user/insert', {
      method: 'POST',
      body: bodyStr,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => {
      ResponseHandler.run(res)
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
      ResponseHandler.run(res)
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

  const authurize = (bodyStr, success = () => {}) => {
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

  const getUserId = (bodyStr) => {
    return fetch('/api/v1/user/auth', {
      method: 'POST',
      body: bodyStr,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => {
      ResponseHandler.run(res)
    }).then(data => {
      UserInfo().set(data)
    })
      .catch(err => console.log(err))
  }

  return {
    addUser: (bodyStr = '') => addUser(bodyStr),
    login: (bodyStr = '') => login(bodyStr),
    authurize: (bodyStr = '', success = () => {}) => authurize(bodyStr, success),
    getUserId: (bodyStr = '') => getUserId(bodyStr)
  }
}
