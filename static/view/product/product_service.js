import IknowToken from '../iknow_token.js'
import CardRenderer from '../card_renderer.js'
import UserInfo from '../user_info.js'
import ProductFormExtractor from '../product_form_extractor.js'
import NodeScriptReplace from '../node_script_replace.js'
import CurrentProduct from '../current_product.js'

export default function ProductService () {
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

  const handleAuthurizationExpired = () => {
    const reverify = confirm('驗證已過期，將重新驗證')

    if (reverify) {
      fetch('/static/view/auth/auth.html').then(res => res.text())
        .then(data => {
          document.body.innerHTML = data
          NodeScriptReplace(document.body)
        })
        .catch(err => console.log(err))
    }
  }

  /**
   *
   *
   * @param {*} [res=Promise()]
   * @param {*} [success=() => { }]
   * @return {Promise().json()}
   */
  const handleResponse = (res = Promise(), success = () => { }) => {
    const d = res.json()
    if (res.status === 200) {
      success()
      return d
    } else if (res.status === 401) {
      handleAuthurizationExpired()
      return d.then(Promise.reject.bind(Promise))
    } else {
      alert('驗證失敗, 請登入')
      return d.then(Promise.reject.bind(Promise))
    }
  }

  const getProduct = async (url = '') => {
    try {
      const res = await fetch(url, {
        method: 'GET',
        headers: getHeaders()
      })
      const value = handleResponse(res)
      return CardRenderer('#cardResult').render(value)
    } catch (err) {
      return console.log(err)
    }
  }

  const updateProduct = async (url) => {
    return fetch(url, {
      method: 'PUT',
      body: JSON.stringify([ProductFormExtractor().extractProduct()]),
      headers: getHeaders()
    }).then(res => handleResponse(res, () => {
      const banner = document.querySelector('.alert')
      banner.hidden = false

      const alertText = document.querySelector('#alertText')
      alertText.innerHTML = '更新成功'
    })).catch(err => console.log(err))
  }

  const deleteProduct = (url) => {
    return fetch(url, {
      method: 'DELETE',
      body: JSON.stringify([CurrentProduct().json()]),
      headers: getHeaders()
    }).then(res => handleResponse(res, () => console.log('Success', res)))
      .catch(err => console.log(err))
  }

  const addProduct = (url) => {
    return fetch(url, {
      method: 'POST',
      body: JSON.stringify([ProductFormExtractor().extractProduct()]),
      headers: getHeaders()
    }).then(res => handleResponse(res, alert('新增成功')))
      .catch(err => console.log(err))
  }

  return {
    getProduct: (url = '') => getProduct(url),
    updateProduct: (url = '') => updateProduct(url),
    deleteProduct: (url = '') => deleteProduct(url),
    addProduct: (url = '') => addProduct(url)
  }
}
