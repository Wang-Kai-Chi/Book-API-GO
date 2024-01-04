import HttpStatusHandler from '../request/http_status_handler.js'
import IknowHeaders from '../request/iknow_headers.js'
import ResponseHandler from '../request/response_handler.js'
import TokenManager from '../util/token_manager.js'

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
    const statusHandler = HttpStatusHandler()
    statusHandler.Unauthorized = () => TokenManager().handleAuthurizationExpired()

    return fetch(url, {
      method: 'GET',
      headers: IknowHeaders().get()
    }).then(res => ResponseHandler().run(res, statusHandler))
      .catch(err => console.log(err.Response))
  }

  const RequestArgs = () => {
    return {
      url: '',
      bodyStr: '',
      success: () => {}
    }
  }

  const ajax = async (met = '', args = RequestArgs()) => {
    const statusHandler = HttpStatusHandler()
    statusHandler.OK = () => args.success()
    statusHandler.Unauthorized = () => TokenManager().handleAuthurizationExpired()
    statusHandler.BadRequest = () => alert('not a product')

    return fetch(args.url, {
      method: met,
      body: args.bodyStr,
      headers: IknowHeaders().get()
    }).then(res => ResponseHandler().run(res, statusHandler))
      .catch(err => alert(err.Response))
  }

  const updateProduct = (args = RequestArgs()) => {
    return ajax('PUT', args)
  }

  const deleteProduct = (args = RequestArgs()) => {
    return ajax('DELETE', args)
  }

  const addProduct = (args = RequestArgs()) => {
    return ajax('POST', args)
  }

  return {
    getProduct: (url = '') => getProduct(url),
    updateProduct: (url = '', success = () => {}, bodyStr = '') => updateProduct({ url, bodyStr, success }),
    deleteProduct: (url = '', success = () => {}, bodyStr = '') => deleteProduct({ url, bodyStr, success }),
    addProduct: (url = '', success = () => {}, bodyStr = '') => addProduct({ url, bodyStr, success })
  }
}
