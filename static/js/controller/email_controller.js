import HttpStatusHandler from '../request/http_status_handler.js'
import ResponseHandler from '../request/response_handler.js'

export default function EmailController () {
  const sendVerificationMail = async (bodyStr) => {
    const statusHandler = HttpStatusHandler()
    statusHandler.BadRequest = () => console.log('user email incorrect')

    return fetch('/api/v1/email/send', {
      method: 'POST',
      body: bodyStr,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => ResponseHandler().run(res, statusHandler))
  }

  return {
    sendVerificationMail: (bodyStr = '') => sendVerificationMail(bodyStr)
  }
}
