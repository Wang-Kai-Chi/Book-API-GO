Result()

function Result() {
    let filters = document.querySelector("#searchInput").value
    if (filters.includes("=")) {
        if (!filters.includes("max"))
            ProductService().getByConditions(filters + "max=500")
        else
            ProductService().getByConditions(filters)
    }
    else
        ProductService().getByBarcode(filters)
}
/**
 * Getting Product json data from api
 *
 * @return {object} return ProductService 
 */
function ProductService() {
    const getByBarcode = (barcode) => {
        fetch(`/api/v1/product/query/barcode/${barcode}`)
            .then(data => data.json())
            .then(value => CardRenderer("#cardResult").render(value))
            .catch(err => console.log(err))
    }
    const getByConditions = (conditions) => {
        fetch(`/api/v1/product/query/?${conditions}`)
            .then(data => data.json())
            .then(value => CardRenderer("#cardResult").render(value))
            .catch(err => console.log(err))
    }
    return {
        getByBarcode: (barcode) => getByBarcode(barcode),
        getByConditions: (conditions) => getByConditions(conditions),
    }
}
/**
 *Bootstrap Card, with dropdown option and icon
 *
 * @param {string} [title=""] title of card
 * @param {string} [sub=""] subtitle of card 
 * @return {string} string of html card 
 */
function CardHTML(title = "", sub = "") {
    return /*html*/`
        <div class="card border-info">
            <div class="card-body py-4 px-4">
                <div class="d-flex align-items-center">
                    <img id="productIcon" src="/static/assets/product32.png" alt="preview">
                    <div class="ms-3 name me-auto">
                        <h5 id="pTitle" class="font-bold">${title}</h5>
                        <h6 id="pPrice" class="text-muted mb-0">${sub}</h6>
                    </div>
                    <div class="dropup">
                        <button class="btn" data-bs-toggle="dropdown" aria-expanded="false">
                            <img src="/static/assets/more32.png" alt="blank">
                        </button>
                        <ul class="dropdown-menu">
                            <li>
                                <a class="dropdown-item" hx-trigger="click" data-bs-dismiss="modal"
                                    hx-get="/static/view/detail.html" hx-swap="innerHTML" hx-target="#main">
                                    <img src="/static/assets/edit32.png" alt="blank">
                                </a>
                            </li>
                            <li>
                                <button id="deleteBtn" type="button" class="dropdown-item" onclick="confirm('Confirm delete?')">
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
 * Rendering bootstrap cards
 * @constructor
 * @param {string} [selector=""] css selector of html element that you want to display card
 * @return {object} return CardRenderer object 
 */
function CardRenderer(selector = "") {
    function renderCards(value) {
        const cards = () => {

            let temp = ""
            for (const v of value)
                temp += CardHTML(v.Product_title, v.Price)

            return temp
        }
        const cardResult = document.querySelector(selector)
        cardResult.innerHTML = cards()
        htmx.process(cardResult)
    }
    return {
        render: (value) => renderCards(value),
    }
}

