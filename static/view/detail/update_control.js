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
    cancelBtn.q().hidden = true
    confirmBtn.q().hidden = true

    updateBtn.q().onclick = () => {
        cancelBtn.q().hidden = false
        confirmBtn.q().hidden = false
        updateBtn.q().hidden = true
        updateController.enableUpdate()
    }

    cancelBtn.q().onclick = () => {
        cancelBtn.q().hidden = true
        confirmBtn.q().hidden = true
        updateBtn.q().hidden = false
        updateController.cancelUpdate()
    }

    confirmBtn.q().onclick=()=>{
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
        productService.update([{
            "Product_id": 1629,
            "Barcode": "9789861052151",
            "Product_title": "妖怪少爺 (9)",
            "Publisher": "東立",
            "Publication_date": "2022-03-19T00:00:00Z",
            "Price": "85元",
            "Quantity": 1,
            "Description": "none"
        }])
            .catch(err => console.log(err))
            .then(response=>console.log("Success",response))
    }

    return {
        enableUpdate: () => enableUpdate(),
        cancelUpdate: () => cancelUpdate(),
        confirmUpdate: () => confirmUpdate(),
    }
}