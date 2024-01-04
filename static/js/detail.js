import CurrentProduct from './localstorage/current_product.js'

import ProductFormExtractor from './util/product_form_extractor.js'
import DetailRenderer from './util/detail_renderer.js'
import ProductController from './controller/product_controller.js'
import DatePicker from './util/date_picker.js'
import { product } from './data/product.js'

/**
 *Showing details of json object
 *
 */
export default function Detail () {
  const addDetailValues = (obj = {}) => {
    const dateId = 'Publication_date'
    DatePicker().set(dateId)

    const keys = Object.keys(obj)
    const current = CurrentProduct().json()

    for (const i in keys) {
      const k = keys[i]
      const el = document.querySelector(`#${k}`)

      if (k === dateId) {
        el.value = current[k].substring(0, 10)
      } else {
        el.value = current[k]
      }
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
    const bodyStr = () => {
      let temp = ''
      try {
        const ps = [ProductFormExtractor().extractProduct()]
        temp = JSON.stringify(ps)
      } catch (err) {
        console.log(err)
      }
      return temp
    }
    ProductController().updateProduct(updateSuccess, bodyStr())
  }

  return {
    enableUpdate: () => enableUpdate(),
    cancelUpdate: () => cancelUpdate(),
    confirmUpdate: () => confirmUpdate()
  }
}
