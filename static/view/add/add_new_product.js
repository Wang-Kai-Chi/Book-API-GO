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

    for (const f of form)
        f.disabled = false
}
