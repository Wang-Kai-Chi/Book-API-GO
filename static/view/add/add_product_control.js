function AddProductControl() {
    const confirmAddProduct = () => {
        const add = async (body) => {
            return fetch("/api/v1/product/insert", {
                method: "POST",
                body: JSON.stringify(body),
                headers: new Headers({
                    "Content-Type": "application/json",
                }),
            }).then(res => {
                let d = res.json()
                if (res.status === 200) {
                    banner.hidden = false
                    alertText.innerHTML = "更新成功"
                    return d
                } else {
                    alert("驗證失敗, 請登入或重新取得驗證碼")
                    return d.then(Promise.reject.bind(Promise));
                }
            })
        }

        add([ProductFormExtractor().extractProduct()])
            .catch(err => console.log(err))
    }
    document.querySelector("#confirmAdd").onclick = () => {
        confirmAddProduct()
    }
}