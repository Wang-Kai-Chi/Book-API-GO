import IknowToken from '../localstorage/iknow_token.js'

export default function JwtController () {
  const getToken = (bodyStr) => {
    return fetch('/api/v1/jwt/token', {
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
        console.log('user info incorrect')
        return d.then(Promise.reject.bind(Promise))
      }
    }).then(data => {
      IknowToken().set(JSON.stringify(data))
      location.reload()
    }).catch(err => console.log(err))
  }

  return {
    getToken: (bodyStr = '') => getToken(bodyStr)
  }
}
