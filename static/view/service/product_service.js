/**
 * Getting Product json data from api
 *@constructor
 * @return {object} return ProductService 
 */
export function ProductService() {
    const getByBarcode = async (barcode) => {
        return fetch(`/api/v1/product/query/barcode/${barcode}`)
            .then(data => data.json())

    }
    const getByConditions = async (conditions) => {
        return fetch(`/api/v1/product/query/?${conditions}`)
            .then(data => data.json())
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
    return {
        getByBarcode: (barcode) => getByBarcode(barcode),
        getByConditions: (conditions) => getByConditions(conditions),
        update: (body=[{}]) => update(body)
    }
}