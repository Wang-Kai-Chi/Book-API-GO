import ProductController from './controller/product_controller.js'
import CardRenderer from './card_renderer.js'

SearchDialog()

function SearchDialog () {
  const showResult = () => {
    const filters = document.querySelector('#searchInput').value
    const controller = ProductController()

    if (filters.includes('=')) {
      if (!filters.includes('max')) {
        controller.getProductsByConditions(filters + 'max=500')
          .then(data => CardRenderer('#cardResult').render(data))
      } else {
        controller.getProductsByConditions(filters)
          .then(data => CardRenderer('#cardResult').render(data))
      }
    } else {
      controller.getProductsByBarcode(filters)
        .then(data => CardRenderer('#cardResult').render(data))
    }
  }

  document.querySelector('#confirmSearch').onclick = () => showResult()
}