import IknowToken from '../localstorage/iknow_token.js'
import ResponseHandler from '../request/response_handler.js'
import HttpStatusHandler from '../request/http_status_handler.js'

/**
 *
 *
 * @export
 * @return {{getToken:(string)=>{}}}
 */
export default function JwtController () {
  const getToken = (bodyStr) => {
    const statusHandler = HttpStatusHandler()
    statusHandler.BadRequest = () => alert('找不到使用者')

    return fetch('/api/v1/jwt/token', {
      method: 'POST',
      body: bodyStr,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => ResponseHandler().run(res, statusHandler))
      .then(data => {
        IknowToken().set(JSON.stringify(data))
      }).catch(err => console.log(err))
  }

  return {
    getToken: (bodyStr = '') => getToken(bodyStr)
  }
}
