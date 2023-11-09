function UpdateControl() {
    const updateBtn = document.querySelector(`#updateBtn`)
    const confirmBtn = document.querySelector(`#confirmUpdateBtn`)
    const cancelBtn = document.querySelector(`#cancelUpdateBtn`)

    const updateController = UpdateController()
    const viewMode = () => {
        cancelBtn.hidden = true
        confirmBtn.hidden = true
    }

    viewMode()

    const editMode = () => {
        cancelBtn.hidden = false
        confirmBtn.hidden = false
    }
    updateBtn.onclick = () => {
        editMode()
        updateBtn.hidden = true
        updateController.enableUpdate()
    }

    cancelBtn.onclick = () => {
        viewMode()
        updateBtn.hidden = false
        updateController.cancelUpdate()
    }

    confirmBtn.onclick = () => {
        updateController.confirmUpdate()
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
        const update = async (body) => {
            return fetch(`/api/v1/product/update`, {
                method: "PUT",
                body: JSON.stringify(body),
                headers: new Headers({
                    "Content-Type": "application/json",
                }),
            }).then(res => res.json())
        }
        const alert=document.querySelector(".alert")
        const alertText = document.querySelector("#alertText")

        update([ProductFormExtractor().extractProduct()])
            .catch(err => {
                alert.hidden = false
                alertText.innerHTML = err
            })
            .then(response => {
                alert.hidden = false
                alertText.innerHTML = "更新成功" 
            })
    }

    return {
        enableUpdate: () => enableUpdate(),
        cancelUpdate: () => cancelUpdate(),
        confirmUpdate: () => confirmUpdate(),
    }
}

