import ProductService from './product_service.js'

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
