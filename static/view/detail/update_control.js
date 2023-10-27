UpdateControl()

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
        const extractProduct = () => {
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

        const update = async (body) => {
            return fetch(`/api/v1/product/update`, {
                method: "PUT",
                body: JSON.stringify(body),
                headers: new Headers({
                    "Content-Type": "application/json",
                }),
            }).then(res => res.json())
        }

        update([extractProduct()])
            .catch(err => console.log(err))
            .then(response => console.log("Success", response))
    }

    return {
        enableUpdate: () => enableUpdate(),
        cancelUpdate: () => cancelUpdate(),
        confirmUpdate: () => confirmUpdate(),
    }
}

