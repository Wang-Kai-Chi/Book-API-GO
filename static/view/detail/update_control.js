import * as p from "../service/product_service.js"
const productService = p.ProductService()
const controllArea = document.querySelector("#updateControl")

UpdateControl()

function UpdateControl() {
    const updateController = UpdateController()

    const updateBtn = UpdateBtnHTML()
    const confirmBtn = ConfirmUpdateBtnHTML()
    const cancelBtn = CancelUpdateBtnHTML()

    controllArea.innerHTML = updateBtn.str + cancelBtn.str + confirmBtn.str

    const viewMode = () => {
        cancelBtn.q().hidden = true
        confirmBtn.q().hidden = true
    }

    viewMode()

    const editMode = () => {
        cancelBtn.q().hidden = false
        confirmBtn.q().hidden = false
    }
    updateBtn.q().onclick = () => {
        editMode()
        extractProduct()
        updateBtn.q().hidden = true
        updateController.enableUpdate()
    }

    cancelBtn.q().onclick = () => {
        viewMode()
        updateBtn.q().hidden = false
        updateController.cancelUpdate()
    }

    confirmBtn.q().onclick = () => {
        updateController.confirmUpdate()
    }
}

function CancelUpdateBtnHTML() {
    const cancelUpdateId = "cancelUpdateBtn"
    return {
        str:/*html*/`
            <button type="button" class="btn btn-danger" 
            id="${cancelUpdateId}">Cancel</button>
        `,
        q: () => { return document.querySelector(`#${cancelUpdateId}`) },
    }
}

function ConfirmUpdateBtnHTML() {
    const confirmUpdateId = "confirmUpdateBtn"
    return {
        str:/*html*/`
            <button type="button" class="btn btn-primary" 
            id="${confirmUpdateId}">Confirm</button>
        `,
        q: () => { return document.querySelector(`#${confirmUpdateId}`) },
    }
}

function UpdateBtnHTML() {
    const updateId = "updateBtn"
    return {
        str:/*html*/`
            <button type="button" class="btn btn-primary" 
            id="${updateId}">Update</button>
        `,
        q: () => { return document.querySelector(`#${updateId}`) },
    }
}

function UpdateController() {
    const form = document.querySelectorAll('.form-control')

    const enableUpdate = () => {
        for (const f of form)
            f.disabled = false
    }

    const cancelUpdate = () => {
        for (const f of form)
            f.disabled = true
    }

    const confirmUpdate = () => {
        const p = () => extractProduct()
        productService.update([p()])
            .catch(err => console.log(err))
            .then(response => console.log("Success", response))
    }

    return {
        enableUpdate: () => enableUpdate(),
        cancelUpdate: () => cancelUpdate(),
        confirmUpdate: () => confirmUpdate(),
    }
}

/**
 *Product entity
 *@constructor
 * @return {object} 
 */
const Product = () => {
    let product = {
        Product_id: 0,
        Product_title: "名稱",
        Price: "價格",
        Barcode: "條碼",
        Publisher: "出版商",
        Publication_date: "發行日",
        Quantity: 0,
        Description: "說明",
    }

    return {
        this: () => { return product },
        keys: () => { return Object.keys(product) },
    }
}

/**
 *Extracting Product() from detail list
 *
 * @return {Product().this()} 
 */
function extractProduct() {
    const product = Product()
    const setValueMatchDataType = (data, value) => {
        if (data == Number.isInteger())
            data = parseInt(value)
        else
            data = value
        return data
    }
    for (const i in product.keys()) {
        const current = product.keys()[i]
        const value = document.querySelector(`#${current}`).value

        let data = setValueMatchDataType(product.this()[current], value)
        product.this()[current] = data
    }
    return product.this()
}