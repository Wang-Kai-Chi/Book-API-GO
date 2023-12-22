export default function LocalStorage (key = '') {
  return {
    string: () => { return localStorage.getItem(key) },
    json: () => { return JSON.parse(localStorage.getItem(key)) },
    remove: () => localStorage.removeItem(key),
    set: (info = '') => localStorage.setItem(key, info)
  }
}
