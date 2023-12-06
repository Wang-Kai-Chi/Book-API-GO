function Result(iknowToken = IknowToken()) {
    let filters = document.querySelector("#searchInput").value
    const token = "Bearer "+iknowToken.json()["Token"]
    const handleResponse = (res) => {
        let d = res.json()
        if (res.status === 200) {
            return d
        } else {
            alert("驗證失敗, 請登入或重新取得驗證碼")
            return d.then(Promise.reject.bind(Promise));
        }
    }
    const getByConditions = async (conditions) => {
        return fetch(`/api/v1/product/query/?${conditions}`, {
            method: "GET",
            headers: new Headers({
                "Content-Type": "application/json",
                "Authorization": token,
            }),
        }).then(res => handleResponse(res))
    }

    const getByBarcode = async (barcode) => {
        return fetch(`/api/v1/product/query/barcode/${barcode}`, {
            method: "GET",
            headers: new Headers({
                "Content-Type": "application/json",
                "Authorization": token,
            }),
        }).then(res => handleResponse(res))
    }

    if (filters.includes("=")) {
        if (!filters.includes("max"))
            getByConditions(filters + "max=500")
                .then(value => CardRenderer("#cardResult").render(value))
                .catch(err => console.log(err))
        else
            getByConditions(filters)
                .then(value => CardRenderer("#cardResult").render(value))
                .catch(err => console.log(err))
    }
    else
        getByBarcode(filters)
            .then(value => CardRenderer("#cardResult").render(value))
            .catch(err => console.log(err))
    htmx.process(document.querySelector("#cardResult"))
}
