import CurrentProduct from './current_product.js'
import Detail from './detail/detail.js'
import UpdateControl from './detail/update_control.js'

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
      document.querySelector(`#deleteBtn${VALUE_ID(i)}`).onclick = () => handleDeleteProduct(`#${VALUE_ID(i)}`)
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
  CurrentProduct().set(cardId)

  const body = `[${localStorage.getItem('currentProduct')}]`
  if (confirm('Confirm delete?')) {
    document.querySelector(`#card${cardId.id}`).hidden = true
    fetch('/api/v1/product/delete', {
      method: 'DELETE',
      body,
      headers: new Headers({
        'Content-Type': 'application/json'
      })
    }).then(res => res.json())
      .catch(err => console.log(err))
      .then(response => console.log('Success', response))
  }
}
