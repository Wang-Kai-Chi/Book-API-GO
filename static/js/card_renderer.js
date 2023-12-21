import CurrentProduct from './localstorage/current_product.js'
import ProductController from './product/product_controller.js'
import DetailRenderer from './detail_renderer.js'

const VALUE_ID = (index) => `v${index}`
/**
 * Rendering bootstrap cards
 * @constructor
 * @param {string} [selector=""] css selector of html element that you want to display card
 * @return {object} return CardRenderer object
 */
export default function CardRenderer (selector = '') {
  const renderCards = (value) => {
    const cards = () => {
      let temp = ''
      for (const i in value) { temp += CardHTML(value[i], i) }

      return temp
    }
    const cardResult = document.querySelector(selector)
    cardResult.innerHTML = cards()

    for (const i in value) {
      document.querySelector(`#editBtn${VALUE_ID(i)}`).onclick = () => {
        CurrentProduct().set(document.querySelector(`#${VALUE_ID(i)}`))

        Detail()
        UpdateControl()

        document.querySelector('#recentProduct').hidden = true
        document.querySelector('#productDetail').hidden = false
      }
      document.querySelector(`#deleteBtn${VALUE_ID(i)}`).onclick = () => handleDeleteProduct(`${VALUE_ID(i)}`)
    }
  }
  return {
    render: (value) => renderCards(value)
  }
}

/**
 *Bootstrap Card, with dropdown option and icon
 *@param {{ Product_title: string; Price: number; }} [product={Product_title:"",Price:0}] product object
 *@param {number} [index=0] index in list
 * @return {string} string of html card
 */
function CardHTML (product = { Product_title: '', Price: 0 }, index = 0) {
  const valueId = VALUE_ID(index)
  return /* html */`
        <div class="card border-info" id="card${valueId}">
            <div class="card-body py-4 px-4">
                <div class="d-flex align-items-center">
                    <img id="productIcon" src="/static/assets/product32.png" alt="preview">
                    <div class="ms-3 name me-auto">
                        <div id="${valueId}" hidden>${JSON.stringify(product)}</div>
                        <h5 id="pTitle" class="font-bold">${product.Product_title}</h5>
                        <h6 id="pPrice" class="text-muted mb-0">${product.Price}</h6>
                    </div>
                    <div class="dropup">
                        <button class="btn" data-bs-toggle="dropdown" aria-expanded="false">
                            <img src="/static/assets/more32.png" alt="blank">
                        </button>
                        <ul class="dropdown-menu">
                            <li>
                                <a class="dropdown-item" id="editBtn${valueId}" data-bs-dismiss="modal">
                                    <img src="/static/assets/edit32.png" alt="blank">
                                </a>
                            </li>
                            <li>
                                <button id="deleteBtn${valueId}" type="button" class="dropdown-item">
                                    <img src="/static/assets/garbage32.png" alt="blank">
                                </button>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
        </div>
        <div class="p-2 g-col-6"></div>
    `
}

/**
 *Removing product card and delete data from DB
 *
 * @param {*} cardId
 */
function handleDeleteProduct (cardId) {
  CurrentProduct().set(document.querySelector(`#${cardId}`))

  if (confirm('Confirm delete?')) {
    ProductController().deleteProduct(() => alert('刪除成功'))
    document.querySelector(`#card${cardId}`).hidden = true
  }
}

/**
 *Showing details of json object
 *
 */
function Detail () {
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

  const confirmUpdate = async () => ProductController().updateProduct(() => {
    const banner = document.querySelector('.alert')
    banner.hidden = false

    const alertText = document.querySelector('#alertText')
    alertText.innerHTML = '更新成功'
  })

  return {
    enableUpdate: () => enableUpdate(),
    cancelUpdate: () => cancelUpdate(),
    confirmUpdate: () => confirmUpdate()
  }
}
