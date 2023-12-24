import CurrentProduct from './localstorage/current_product.js'
import ProductFormExtractor from './product_form_extractor.js'
import ProductController from './controller/product_controller.js'
import DetailRenderer from './detail_renderer.js'

/**
 *Showing details of json object
 *
 */
export default function Detail () {
  const product = {
    Product_id: 'id',
    Product_title: '名稱',
    Price: '價格',
    Barcode: '條碼',
    Publisher: '出版商',
    Publication_date: '發行日',
    Quantity: '數量',
    Description: '說明'
  }

  const setDatePicker = (id) => {
    const currentDate = new Date().toJSON().slice(0, 10)
    const publicationDate = document.querySelector(`#${id}`)

    publicationDate.type = 'date'
    publicationDate.min = '1900-01-01'
    publicationDate.max = `${currentDate}`
  }

  const addDetailValues = (obj = {}) => {
    const keys = Object.keys(obj)
    const current = CurrentProduct().json()
    const dateId = 'Publication_date'
    setDatePicker(dateId)

    for (const i in keys) {
      const k = keys[i]
      const el = document.querySelector(`#${k}`)

      if (k === dateId) {
        el.value = current[k].substring(0, 10)
      } else { el.value = current[k] }
    }
  }

  DetailRenderer('#detailDisplay').render(product)
  addDetailValues(product)

  document.querySelector('#formProduct_id').hidden = true

  UpdateControl()
}

function UpdateControl () {
  const updateBtn = document.querySelector('#updateBtn')
  const confirmBtn = document.querySelector('#confirmUpdateBtn')
  const cancelBtn = document.querySelector('#cancelUpdateBtn')

  const updateController = UpdateController()
  const viewMode = () => {
    cancelBtn.hidden = true
    confirmBtn.hidden = true
  }

  viewMode()

  const editMode = () => {
    cancelBtn.hidden = false
    confirmBtn.hidden = false
  }
  updateBtn.onclick = () => {
    editMode()
    updateBtn.hidden = true
    updateController.enableUpdate()
  }

  cancelBtn.onclick = () => {
    viewMode()
    updateBtn.hidden = false
    updateController.cancelUpdate()
  }

  confirmBtn.onclick = () => {
    updateController.confirmUpdate()
  }
}

function UpdateController () {
  const form = document.querySelectorAll('.form-control')

  const enableUpdate = () => {
    for (const f of form) { f.disabled = false }
  }

  const cancelUpdate = () => {
    for (const f of form) { f.disabled = true }
  }

  function updateSuccess () {
    const banner = document.querySelector('.alert')
    banner.hidden = false

    const alertText = document.querySelector('#alertText')
    alertText.innerHTML = '更新成功'
  }

  const confirmUpdate = async () => {
    const bodyStr = JSON.stringify([ProductFormExtractor().extractProduct()])
    ProductController().updateProduct(updateSuccess, bodyStr)
  }

  return {
    enableUpdate: () => enableUpdate(),
    cancelUpdate: () => cancelUpdate(),
    confirmUpdate: () => confirmUpdate()
  }
}
