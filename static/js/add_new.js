import ProductFormExtractor from './util/product_form_extractor.js'
import DetailRenderer from './util/detail_renderer.js'
import ProductController from './controller/product_controller.js'
import DatePicker from './util/date_picker.js'
import { product } from './data/product.js'

/**
 *Showing details of json object
 *
 */
export default function AddNew () {
  DetailRenderer('#addDetail').render(product)
  document.querySelector('#formProduct_id').hidden = true
  const form = document.querySelectorAll('.form-control')

  DatePicker().set('#Publication_date')

  for (const f of form) { f.disabled = false }

  document.querySelector('#confirmAdd').onclick = () => ProductController().addProduct(
    () => alert('新增成功'),
    JSON.stringify([ProductFormExtractor().extractProduct()]))
}
