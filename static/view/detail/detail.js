/**
 *Showing details of json object
 *
 */
function Detail() {
    const product = {
        Product_id: "id",
        Product_title: "名稱",
        Price: "價格",
        Barcode: "條碼",
        Publisher: "出版商",
        Publication_date: "發行日",
        Quantity: "數量",
        Description: "說明",
    }

    const addDetailValues = (obj = {}, storageKey = "") => {
        const keys = Object.keys(obj)
        const current = JSON.parse(localStorage.getItem(storageKey))
        for (const i in keys) {
            const k = keys[i]
            document.querySelector(`#${k}`).value = current[k]
        }
    }

    DetailRenderer("#detailDisplay").render(product)
    addDetailValues(product, "currentProduct")
    document.querySelector("#formProduct_id").hidden = true
}
