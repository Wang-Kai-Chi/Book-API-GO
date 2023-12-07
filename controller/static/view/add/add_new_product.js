/**
 *Showing details of json object
 *
 */
function AddNewProduct() {
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

    DetailRenderer("#addDetail").render(product)
    document.querySelector("#formProduct_id").hidden = true
    const form = document.querySelectorAll('.form-control')

    const setDatePicker = (id) => {
        const currentDate = new Date().toJSON().slice(0, 10)
        const publicationDate = document.querySelector(`#${id}`)

        publicationDate.type = "date"
        publicationDate.min = "1900-01-01"
        publicationDate.max = `${currentDate}`
    }

    setDatePicker("Publication_date")

    for (const f of form)
        f.disabled = false
}
