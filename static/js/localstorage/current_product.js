/**
 *Interacting with currentProduct in localStorage
 * @constructor
 * @return {*}
 */
export default function CurrentProduct () {
  const key = 'currentProduct'

  return {
    json: () => { return JSON.parse(localStorage.getItem(key)) },
    set: (cardId = '') => localStorage.setItem(key, cardId.innerHTML)
  }
}
