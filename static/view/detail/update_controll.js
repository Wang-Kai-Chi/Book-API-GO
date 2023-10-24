const controllArea = document.querySelector("#updateControll")
const confirmUpdateId = "confirmUpdateBtn"
const updateController = UpdateController("updateController")

UpdateControll()

function UpdateControll() {
    controllArea.innerHTML = UpdateBtnHTML(updateController.name)
}

function CancelUpdateBtnHTML(controllerName = "") {
    return /*html*/`
            <button type="button" class="btn btn-danger" 
            onclick="${controllerName}.cancelUpdate()">Cancel</button>
        `
}

function ConfirmUpdateBtnHTML(controllerName="") {
    return /*html*/`
            <button type="button" class="btn btn-primary" id="${confirmUpdateId}" 
            hx-put="/api/v1/product/update" hx-trigger="click" onclick="${controllerName}.confirmUpdate()">Confirm</button>
        `
}

function UpdateBtnHTML(controllerName = "") {
    return /*html*/`
            <button type="button" class="btn btn-primary" 
            onclick="${controllerName}.enableUpdate()">Update</button>
        `
}

function UpdateController(name = "") {
    const form = document.querySelectorAll('.form-control')

    const enableUpdate = () => {
        for (const f of form)
            f.disabled = false

        controllArea.innerHTML = CancelUpdateBtnHTML(name) + ConfirmUpdateBtnHTML(name)
    const confirmBtn = document.querySelector(`#${confirmUpdateId}`)
        htmx.process(confirmBtn)
    }

    const cancelUpdate = () => {
        for (const f of form)
            f.disabled = true
        controllArea.innerHTML = UpdateBtnHTML(name)
    }

    const confirmUpdate = () => {
    }

    return {
        enableUpdate: () => enableUpdate(),
        cancelUpdate: () => cancelUpdate(),
        confirmUpdate:()=>confirmUpdate(),
        name: name
    }
}