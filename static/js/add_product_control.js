import ProductController from './product/product_controller.js'

export default function AddProductControl () {
  document.querySelector('#confirmAdd').onclick = () => ProductController().addProduct()
}
