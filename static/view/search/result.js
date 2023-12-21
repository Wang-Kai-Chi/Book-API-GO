import ProductController from '../product/product_controller.js'

export default function Result () {
  const filters = document.querySelector('#searchInput').value
  const controller = ProductController()

  if (filters.includes('=')) {
    if (!filters.includes('max')) {
      controller.getProductsByConditions(filters + 'max=500')
    } else {
      controller.getProductsByConditions(filters)
    }
  } else {
    controller.getProductsByBarcode(filters)
  }
}
