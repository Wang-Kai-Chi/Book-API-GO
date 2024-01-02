export default function EmailController () {
  const sendVerificationMail = async (bodyStr) => {
    return fetch('/api/v1/email/send', {
      method: 'POST',
      body: bodyStr,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => {
      return res.json()
    })
  }

  return {
    sendVerificationMail: (bodyStr = '') => sendVerificationMail(bodyStr)
  }
}
