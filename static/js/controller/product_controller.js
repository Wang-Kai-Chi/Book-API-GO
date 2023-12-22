import ProductFormExtractor from '../product_form_extractor.js'
import IknowHeaders from '../request/IknowHeaders.js'
import ResponseHandler from '../request/ResponseHandler.js'
import CurrentProduct from '../localstorage/current_product.js'

/**
 *
 *
 * @export
 * @return {*}
 */
export default function ProductController () {
  const service = ProductService()

  const getProductsByConditions = async (conditions) => {
    return service.getProduct(`/api/v1/product/query/?${conditions}`)
  }

  const getProductsByBarcode = async (barcode) => {
    return service.getProduct(`/api/v1/product/query/barcode/${barcode}`)
  }

  const updateProduct = async (success = () => {}) => {
    return service.updateProduct('/api/v1/product/update', success())
  }

  const deleteProduct = async (success = () => {}) => {
    return service.deleteProduct('/api/v1/product/delete', success())
  }

  const addProduct = async (success = () => {}) => {
    return service.addProduct('/api/v1/product/insert', success())
  }

  return {
    getProductsByConditions: (conditions = '') => getProductsByConditions(conditions),
    getProductsByBarcode: (barcode = '') => getProductsByBarcode(barcode),
    updateProduct: (success = () => {}) => updateProduct(success()),
    deleteProduct: (success = () => {}) => deleteProduct(success()),
    addProduct: (success = () => {}) => addProduct(success())
  }
}

function ProductService () {
  const getProduct = async (url = '') => {
    return fetch(url, {
      method: 'GET',
      headers: IknowHeaders().get()
    }).then(res => ResponseHandler().run(res))
      .catch(err => console.log(err))
  }

  const updateProduct = async (url, success = () => {}) => {
    return fetch(url, {
      method: 'PUT',
      body: JSON.stringify([ProductFormExtractor().extractProduct()]),
      headers: IknowHeaders().get()
    }).then(res => ResponseHandler().run(res, success())).catch(err => console.log(err))
  }

  const deleteProduct = async (url, success = () => {}) => {
    return fetch(url, {
      method: 'DELETE',
      body: JSON.stringify([CurrentProduct().json()]),
      headers: IknowHeaders().get()
    }).then(res => ResponseHandler().run(res, success(res)))
      .catch(err => console.log(err))
  }

  const addProduct = async (url, success = () => {}) => {
    return fetch(url, {
      method: 'POST',
      body: JSON.stringify([ProductFormExtractor().extractProduct()]),
      headers: IknowHeaders().get()
    }).then(res => ResponseHandler().run(res, success()))
      .catch(err => console.log(err))
  }

  return {
    getProduct: (url = '') => getProduct(url),
    updateProduct: (url = '', success = () => {}) => updateProduct(url, success()),
    deleteProduct: (url = '', success = () => {}) => deleteProduct(url, success()),
    addProduct: (url = '', success = () => {}) => addProduct(url, success())
  }
}
