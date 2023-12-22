import LocalStorage from './local_storage.js'

/**
 * Interacting with userinfo in localStorage
 * @constructor
 * @return {LocalStorage}
 */
export default function UserInfo () {
  return LocalStorage('userinfo')
}
