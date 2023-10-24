/**
 * Getting Product json data from api
 *
 * @return {object} return ProductService 
 */
export function ProductService() {
    const getByBarcode = (barcode) => {
        return fetch(`/api/v1/product/query/barcode/${barcode}`)
            .then(data => data.json())
            
    }
    const getByConditions = (conditions) => {
        return fetch(`/api/v1/product/query/?${conditions}`)
            .then(data => data.json())
    }
    return {
        getByBarcode: (barcode) => getByBarcode(barcode),
        getByConditions: (conditions) => getByConditions(conditions),
    }
}