/**
 * Rendering bootstrap cards
 * @constructor
 * @param {string} [selector=""] css selector of html element that you want to display card
 * @return {object} return CardRenderer object 
 */
function CardRenderer(selector = "") {
    const renderCards = (value) => {
        const cards = () => {
            let temp = ""
            for (const i in value)
                temp += CardHTML(value[i], i)

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

/**
 *Bootstrap Card, with dropdown option and icon
 *@param {{ Product_title: string; Price: number; }} [product={Product_title:"",Price:0}] product object
 *@param {number} [index=0] index in list
 * @return {string} string of html card 
 */
function CardHTML(product = { Product_title: "", Price: 0 }, index = 0) {
    const VALUE_ID = `pValue${index}`
    const PRODUCT_DETAIL_TEMPLATE_URI = "/static/view/detail/detail.html"
    return /*html*/`
        <div class="card border-info" id="card${VALUE_ID}">
            <div class="card-body py-4 px-4">
                <div class="d-flex align-items-center">
                    <img id="productIcon" src="/static/assets/product32.png" alt="preview">
                    <div class="ms-3 name me-auto">
                        <div id="${VALUE_ID}" hidden>${JSON.stringify(product)}</div>
                        <h5 id="pTitle" class="font-bold">${product.Product_title}</h5>
                        <h6 id="pPrice" class="text-muted mb-0">${product.Price}</h6>
                    </div>
                    <div class="dropup">
                        <button class="btn" data-bs-toggle="dropdown" aria-expanded="false">
                            <img src="/static/assets/more32.png" alt="blank">
                        </button>
                        <ul class="dropdown-menu">
                            <li>
                                <a class="dropdown-item" hx-trigger="click" data-bs-dismiss="modal"
                                    hx-get="${PRODUCT_DETAIL_TEMPLATE_URI}"  hx-swap="innerHTML" 
                                    hx-target="#main" onclick="CurrentProduct().set(${VALUE_ID})">
                                    <img src="/static/assets/edit32.png" alt="blank">
                                </a>
                            </li>
                            <li>
                                <button id="deleteBtn" type="button" class="dropdown-item" onclick="handleDeleteProduct(${VALUE_ID})">
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
function handleDeleteProduct(cardId) {
    CurrentProduct().set(cardId)

    let body = `[${localStorage.getItem("currentProduct")}]`
    if (confirm('Confirm delete?')) {
        document.querySelector(`#card${cardId.id}`).hidden = true
        fetch(`/api/v1/product/delete`, {
            method: "DELETE",
            body: body,
            headers: new Headers({
                "Content-Type": "application/json",
            }),
        }).then(res => res.json())
            .catch(err => console.log(err))
            .then(response => console.log("Success", response))
    }
}

