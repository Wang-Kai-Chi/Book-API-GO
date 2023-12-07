function UpdateControl(iknowToken = IknowToken()) {
    const updateBtn = document.querySelector(`#updateBtn`)
    const confirmBtn = document.querySelector(`#confirmUpdateBtn`)
    const cancelBtn = document.querySelector(`#cancelUpdateBtn`)

    const updateController = UpdateController(iknowToken)
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

function UpdateController(iknowToken = IknowToken()) {
    const form = document.querySelectorAll('.form-control')

    const enableUpdate = () => {
        for (const f of form)
            f.disabled = false
    }

    const cancelUpdate = () => {
        for (const f of form)
            f.disabled = true
    }

    const confirmUpdate = async () => {
        const alertText = document.querySelector("#alertText")
        const token = (iknowToken.json() === null) ?
            "" :
            "Bearer " + iknowToken.json()["Token"]
        fetch(`/api/v1/product/update`, {
            method: "PUT",
            body: JSON.stringify([ProductFormExtractor().extractProduct()]),
            headers: new Headers({
                "Content-Type": "application/json",
                "Authorization": token,
            }),
        }).then(res => {
            let d = res.json()
            if (res.status === 200) {
                const banner = document.querySelector(".alert")
                banner.hidden = false
                alertText.innerHTML = "更新成功"
                return d
            } else {
                alert("驗證失敗, 請登入或重新取得驗證碼")
                return d.then(Promise.reject.bind(Promise));
            }
        }).catch(err => console.log(err))
    }

    return {
        enableUpdate: () => enableUpdate(),
        cancelUpdate: () => cancelUpdate(),
        confirmUpdate: () => confirmUpdate(),
    }
}

