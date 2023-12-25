import ProductController from './controller/product_controller.js'
import CardRenderer from './card_renderer.js'

export default function SearchDialog () {
  const filters = document.querySelector('#searchInput').value
  const controller = ProductController()

  if (filters.includes('=')) {
    if (!filters.includes('max')) {
      controller.getProductsByConditions(filters + 'max=500')
        .then(data => CardRenderer('#resultView').render(data))
    } else {
      controller.getProductsByConditions(filters)
        .then(data => CardRenderer('#resultView').render(data))
    }
  } else {
    controller.getProductsByBarcode(filters)
      .then(data => CardRenderer('#resultView').render(data))
  }
}
