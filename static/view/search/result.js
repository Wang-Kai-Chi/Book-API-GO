Result()

function Result() {
    let filters = document.querySelector("#searchInput").value

    const getByConditions = async (conditions) => {
        return fetch(`/api/v1/product/query/?${conditions}`)
            .then(data => data.json())
    }

    const getByBarcode = async (barcode) => {
        return fetch(`/api/v1/product/query/barcode/${barcode}`)
            .then(data => data.json())
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
