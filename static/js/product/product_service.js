import CardRenderer from '../card_renderer.js'
import ProductFormExtractor from '../product_form_extractor.js'
import IknowHeaders from '../request/IknowHeaders.js'
import ResponseHandler from '../request/ResponseHandler.js'
import CurrentProduct from '../localstorage/current_product.js'

export default function ProductService () {
  const getProduct = async (url = '') => {
    return fetch(url, {
      method: 'GET',
      headers: IknowHeaders().get()
    }).then(res => ResponseHandler().run(res))
      .then(data => CardRenderer('#cardResult').render(data))
      .catch(err => console.log(err))
  }

  const updateProduct = async (url) => {
    return fetch(url, {
      method: 'PUT',
      body: JSON.stringify([ProductFormExtractor().extractProduct()]),
      headers: IknowHeaders().get()
    }).then(res => ResponseHandler().run(res, () => {
      const banner = document.querySelector('.alert')
      banner.hidden = false

      const alertText = document.querySelector('#alertText')
      alertText.innerHTML = '更新成功'
    })).catch(err => console.log(err))
  }

  const deleteProduct = async (url) => {
    try {
      const res = await fetch(url, {
        method: 'DELETE',
        body: JSON.stringify([CurrentProduct().json()]),
        headers: IknowHeaders().get()
      })
      return await ResponseHandler().run(res, () => console.log('Success', res))
    } catch (err) {
      return console.log(err)
    }
  }

  const addProduct = async (url) => {
    try {
      const res = await fetch(url, {
        method: 'POST',
        body: JSON.stringify([ProductFormExtractor().extractProduct()]),
        headers: IknowHeaders().get()
      })
      return await ResponseHandler().run(res, alert('新增成功'))
    } catch (err) {
      return console.log(err)
    }
  }

  return {
    getProduct: (url = '') => getProduct(url),
    updateProduct: (url = '') => updateProduct(url),
    deleteProduct: (url = '') => deleteProduct(url),
    addProduct: (url = '') => addProduct(url)
  }
}
