import LocalStorage from './local_storage.js'
/**
 * Interacting with iknow bearer token in localStorage
 * @constructor
 * @return {LocalStorage}
 */
export default function IknowToken () {
  return LocalStorage('iknowtoken')
}
