import IknowToken from '../iknow_token.js'
import CardRenderer from '../card_renderer.js'
import UserInfo from '../user_info.js'

export default function ProductController () {
  const service = ProductService()
  const getProductsByConditions = async (conditions) => {
    return service.getProduct(`/api/v1/product/query/?${conditions}`)
  }

  const getProductsByBarcode = async (barcode) => {
    return service.getProduct(`/api/v1/product/query/barcode/${barcode}`)
  }

  return {
    getProductsByConditions: (conditions) => getProductsByConditions(conditions),
    getProductsByBarcode: (barcode) => getProductsByBarcode(barcode)
  }
}

function ProductService () {
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

  const getProduct = (url = '') => {
    const handleResponse = (res) => {
      const d = res.json()
      if (res.status === 200) {
        return d
      } else {
        alert('驗證失敗, 請登入或重新取得驗證碼')
        return d.then(Promise.reject.bind(Promise))
      }
    }
    return fetch(url, {
      method: 'GET',
      headers: getHeaders()
    }).then(res => handleResponse(res))
      .then(value => CardRenderer('#cardResult').render(value))
      .catch(err => console.log(err))
  }

  return {
    getProduct: (url = '') => getProduct(url)
  }
}
