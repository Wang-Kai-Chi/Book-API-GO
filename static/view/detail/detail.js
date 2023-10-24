Detail()

/**
 *Showing details of json object
 *
 */
function Detail() {
    const product = {
        Product_title: "名稱",
        Price: "價格",
        Barcode: "條碼",
        Publisher: "出版商",
        Publication_date: "發行日",
        Quantity: "數量",
        Description: "說明",
    }

    const addDetailValues = (obj={}, storageKey="") => {
        const keys = Object.keys(obj)
        const current= JSON.parse(localStorage.getItem(storageKey))
        for (const i in keys) {
            const k = keys[i]
            document.querySelector(`#${k}`).value = current[k]
        }
    }

    DetailRenderer("#detailDisplay").render(product)
    addDetailValues(product, "currentProduct")
}
/**
 * Bootstrap input group
 *
 * @param {string} [name=""] name of input group
 * @param {string} [id=""] input id
 * @return {string} html element string 
 */
function DetailHTML(name = "", id = "") {
    return /*html*/`
        <div class="input-group mb-3">
            <span class="input-group-text">${name}</span>
            <input type="text" class="form-control" id="${id}" placeholder="none" disabled>
        </div>
    `
}

/**
 * Render list items of detail
 *
 * @param {string} [selector=""] selector of detail list
 * @return {object} 
 */
function DetailRenderer(selector = "") {
    const render = (details) => {
        const items = () => {
            let temp = ""
            const keys = Object.keys(details)
            for (const i in keys) {
                const k = keys[i]
                temp += DetailHTML(details[k], k)
            }

            return temp
        }
        const detailDisplay = document.querySelector(selector)
        detailDisplay.innerHTML = items()
    }
    return {
        render: (details = {}) => render(details)
    }
}