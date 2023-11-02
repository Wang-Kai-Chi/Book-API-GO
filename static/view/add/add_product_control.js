function AddProductControl() {
    const confirmAddProduct = () => {
        const add = async (body) => {
            return fetch("/api/v1/product/insert", {
                method: "POST",
                body: JSON.stringify(body),
                headers: new Headers({
                    "Content-Type": "application/json",
                }),
            }).then(res => res.json())
        }

        add([ProductFormExtractor().extractProduct()])
        .catch(err => console.log(err))
        .then(res => console.log("Success", res))
    }
    document.querySelector("#confirmAdd").onclick = () => {
        confirmAddProduct()
    }
}