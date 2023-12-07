function ProductFormExtractor() {
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

    return {
        extractProduct: () => extractProduct()
    }
}