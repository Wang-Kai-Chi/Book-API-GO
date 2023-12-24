import IknowHeaders from '../request/iknow_headers.js'
import ResponseHandler from '../request/response_handler.js'

/**
 *
 *
 * @export
 * @return {*}
 */
export default function ProductController () {
  const service = ProductService()

  const getProductsByConditions = (conditions) => {
    return service.getProduct(`/api/v1/product/query/?${conditions}`)
  }

  const getProductsByBarcode = (barcode) => {
    return service.getProduct(`/api/v1/product/query/barcode/${barcode}`)
  }

  const updateProduct = (success = () => {}, bodyStr = '') => {
    return service.updateProduct('/api/v1/product/update', success, bodyStr)
  }

  const deleteProduct = (success = () => {}, bodyStr = '') => {
    return service.deleteProduct('/api/v1/product/delete', success, bodyStr)
  }

  const addProduct = (success = () => {}, bodyStr = '') => {
    return service.addProduct('/api/v1/product/insert', success, bodyStr)
  }

  return {
    getProductsByConditions: (conditions = '') => getProductsByConditions(conditions),
    getProductsByBarcode: (barcode = '') => getProductsByBarcode(barcode),
    updateProduct: (success = () => {}, bodyStr = '') => updateProduct(success, bodyStr),
    deleteProduct: (success = () => {}, bodyStr = '') => deleteProduct(success, bodyStr),
    addProduct: (success = () => {}, bodyStr = '') => addProduct(success, bodyStr)
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

  const ajax = async (url = '', met = '', bodyStr = '', success = () => {}) => {
    return fetch(url, {
      method: met,
      body: bodyStr,
      headers: IknowHeaders().get()
    }).then(res => ResponseHandler().run(res, success))
      .catch(err => console.log(err))
  }

  const updateProduct = (url, success = () => {}, bodyStr = '') => {
    return ajax(url, 'PUT', bodyStr, success)
  }

  const deleteProduct = (url, success = () => {}, bodyStr = '') => {
    return ajax(url, 'DELETE', bodyStr, success)
  }

  const addProduct = (url, success = () => {}, bodyStr = '') => {
    return ajax(url, 'POST', bodyStr, success)
  }

  return {
    getProduct: (url = '') => getProduct(url),
    updateProduct: (url = '', success = () => {}, bodyStr = '') => updateProduct(url, success, bodyStr),
    deleteProduct: (url = '', success = () => {}, bodyStr = '') => deleteProduct(url, success, bodyStr),
    addProduct: (url = '', success = () => {}, bodyStr = '') => addProduct(url, success, bodyStr)
  }
}
