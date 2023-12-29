/**
 * Encapsulating localStorage API
 *
 * @export
 * @param {string} [key='']
 * @return {*}
 */
export default function LocalStorage (key = '') {
  return {
    string: () => { return localStorage.getItem(key) },
    json: () => {
      let temp
      try {
        temp = JSON.parse(localStorage.getItem(key))
      } catch (err) {
        temp = ''
      }
      return temp
    },
    remove: () => localStorage.removeItem(key),
    set: (info = '') => localStorage.setItem(key, info)
  }
}
