Result()

function Result() {
    const filters = document.querySelector("#searchInput").value
    fetch(`/api/v1/product/query/?${filters}`)
        .then(data => data.json())
        .then(value => parseValue(value))
        .catch(err => console.log(err))
}
function parseValue(value) {
    renderCards(value)
}

function renderCards(value) {
    const cardTemplate = (title = "", price = "") => {
        return /*html*/`
    <div class="card border-info">
        <div class="card-body py-4 px-4">
            <div class="d-flex align-items-center">
                <img id="productIcon" src="/static/assets/product32.png" alt="preview">
                <div class="ms-3 name me-auto">
                    <h5 id="pTitle" class="font-bold">${title}</h5>
                    <h6 id="pPrice" class="text-muted mb-0">${price}</h6>
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
    const cards = () => {
        let temp = ""
        for (const v of value) {
            temp += cardTemplate(v.Product_title, v.Price)
        }
        return temp
    }
    const cardResult = document.querySelector("#cardResult")
    cardResult.innerHTML = cards()
    htmx.process(cardResult)
}
