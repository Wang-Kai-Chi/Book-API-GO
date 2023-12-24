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

  const RequestArgs = () => {
    return {
      url: '',
      bodyStr: '',
      success: () => {}
    }
  }

  const ajax = async (met = '', args = RequestArgs()) => {
    return fetch(args.url, {
      method: met,
      body: args.bodyStr,
      headers: IknowHeaders().get()
    }).then(res => ResponseHandler().run(res, args.success))
      .catch(err => console.log(err))
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
