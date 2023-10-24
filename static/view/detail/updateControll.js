const controllArea = document.querySelector("#updateControll")

UpdateControll()

function UpdateControll() {
    controllArea.innerHTML = UpdateBtnHTML()
}

function CancelUpdateBtnHTML() {
    return /*html*/`
            <button type="button" class="btn btn-danger" 
            onclick="UpdateController().cancelUpdate()">Cancel</button>
        `
}

function ConfirmUpdateBtnHTML() {
    return /*html*/`
            <button type="button" class="btn btn-primary" 
            onclick="">Confirm</button>
        `
}

function UpdateBtnHTML() {
    return /*html*/`
            <button type="button" class="btn btn-primary" 
            onclick="UpdateController().enableUpdate()">Update</button>
        `
}

function UpdateController() {
    const enableUpdate = () => {
        const form = document.querySelectorAll('.form-control')
        for (const f of form)
            f.disabled = false
        controllArea.innerHTML = CancelUpdateBtnHTML() + ConfirmUpdateBtnHTML()
    }

    const cancelUpdate = () => {
        const form = document.querySelectorAll('.form-control')
        for (const f of form)
            f.disabled =true 
        controllArea.innerHTML = UpdateBtnHTML() 

    }

    return {
        enableUpdate: () => enableUpdate(),
        cancelUpdate:()=>cancelUpdate(),
    }
}