import UserInfo from '../localstorage/user_info.js'
import JwtController from './jwt_controller.js'
import ResponseHandler from '../request/response_handler.js'
import HttpStatusHandler from '../request/http_status_handler.js'

/**
 *
 *
 * @export
 * @return {{
 * addUser: (string) => {},
 * login: (string) => {},
 * authurize: (string, function) => {}
 * }}
 */
export default function UserController () {
  const addUser = async (bodyStr) => {
    const statusHandler = HttpStatusHandler()
    statusHandler.BadRequest = () => alert('帳戶已存在，請使用其他名稱或電子郵件')

    return fetch('/api/v1/user/insert', {
      method: 'POST',
      body: bodyStr,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => ResponseHandler().run(res, statusHandler))
      .then(data => {
        try {
          UserInfo().set(JSON.stringify(data))
        } catch (err) {
          console.log('store userId failed. error: ' + err)
        }
      })
      .catch(err => console.log(err))
  }

  const handleLogin = async (data) => {
    const noPswUser = (data) => {
      const temp = data
      temp.Password = ''
      return JSON.stringify(temp)
    }
    try {
      UserInfo().set(noPswUser(data))
    } catch (err) {
      console.log(err)
    }
    JwtController().getToken(JSON.stringify(data))
  }

  const login = async (bodyStr) => {
    const statusHandler = HttpStatusHandler()
    statusHandler.BadRequest = () => alert('電子郵件錯誤')
    statusHandler.Unauthorized = () => alert('密碼錯誤')

    return fetch('/api/v1/user/login', {
      method: 'POST',
      body: bodyStr,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => ResponseHandler().run(res, statusHandler))
      .then(data => {
        handleLogin(data)
      }).catch(err => console.log(err))
  }

  const authurize = async (bodyStr, success) => {
    const statusHandler = HttpStatusHandler()
    statusHandler.OK = success

    return fetch('/api/v1/user/auth', {
      method: 'POST',
      body: bodyStr,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => {
      ResponseHandler().run(res, statusHandler)
    }).catch(err => console.log(err))
  }

  return {
    addUser: (bodyStr = '') => addUser(bodyStr),
    login: (bodyStr = '') => login(bodyStr),
    authurize: (bodyStr = '', success = () => { }) => authurize(bodyStr, success)
  }
}
