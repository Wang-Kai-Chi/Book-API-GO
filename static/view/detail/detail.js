/**
 *Showing details of json object
 *
 */
function Detail(currentProduct=CurrentProduct()) {
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

    const setDatePicker = (id) => {
        const currentDate = new Date().toJSON().slice(0, 10)
        const publicationDate = document.querySelector(`#${id}`)

        publicationDate.type = "date"
        publicationDate.min = "1900-01-01"
        publicationDate.max = `${currentDate}`
    }

    const addDetailValues = (obj = {}) => {
        const keys = Object.keys(obj)
        const current = currentProduct.json() 
        const dateId = "Publication_date"
        setDatePicker(dateId)

        for (const i in keys) {
            const k = keys[i]
            const el = document.querySelector(`#${k}`)

            if (k === dateId) {
                el.value = current[k].substring(0, 10)
            } else
                el.value = current[k]
        }
    }

    DetailRenderer("#detailDisplay").render(product)
    addDetailValues(product)

    document.querySelector("#formProduct_id").hidden = true
}
