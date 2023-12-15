const iknow_token = () => IknowToken()
/**
 * Interacting with iknow bearer token in localStorage
 * @constructor
 * @return {UserInfo}
 * json(): json of IknowToken
 * remove(): remove IknowToken from localStorage
 * set(): set string of IknowToken to localStorage
 */
export default function IknowToken () {
  const id = 'iknowtoken'

  return {
    json: () => { return JSON.parse(localStorage.getItem(id)) },
    remove: () => localStorage.removeItem(id),
    set: (info = '') => localStorage.setItem(id, info)
  }
}
