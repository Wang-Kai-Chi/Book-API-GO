import UserInfo from '../localstorage/user_info.js'
import IknowToken from '../localstorage/iknow_token.js'

export default function IknowHeaders () {
  const getHeaders = () => {
    const token = (IknowToken().json() === null)
      ? ''
      : 'Bearer ' + IknowToken().json().Token

    const auth = (IknowToken().json() === null)
      ? ''
      : UserInfo().json().Auth

    const iknowHeaders = new Headers({
      'Content-Type': 'application/json',
      Authorization: token,
      'Auth-Key': auth
    })

    return iknowHeaders
  }

  return {
    get: () => getHeaders()
  }
}
