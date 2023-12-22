import LocalStorage from './local_storage.js'
/**
 *Interacting with currentProduct in localStorage
 * @constructor
 * @return {LocalStorage}
 */
export default function CurrentProduct () {
  return LocalStorage('currentProduct')
}
