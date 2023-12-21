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

  const updateProduct = async () => {
    return service.updateProduct('/api/v1/product/update')
  }

  const deleteProduct = async () => {
    return service.deleteProduct('/api/v1/product/delete')
  }

  return {
    getProductsByConditions: (conditions = '') => getProductsByConditions(conditions),
    getProductsByBarcode: (barcode = '') => getProductsByBarcode(barcode),
    updateProduct: () => updateProduct(),
    deleteProduct: () => deleteProduct()
  }
}
