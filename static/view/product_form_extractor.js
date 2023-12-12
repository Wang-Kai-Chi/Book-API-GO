// eslint-disable-next-line no-unused-vars
function ProductFormExtractor () {
  /**
     *Product entity
    *@constructor
    * @return {object}
    */
  const Product = () => {
    const product = {
      Product_id: 0,
      Product_title: '名稱',
      Price: '價格',
      Barcode: '條碼',
      Publisher: '出版商',
      Publication_date: '發行日',
      Quantity: 0,
      Description: '說明'
    }

    return {
      this: () => { return product },
      keys: () => { return Object.keys(product) }
    }
  }

  /**
     *Extracting Product() from detail list
     *
     * @return {Product().this()}
     */
  const extractProduct = () => {
    const product = Product().this()
    const keys = Product().keys()

    const setValueMatchDataType = (data, value) => {
      if (Number.isInteger(data)) data = parseInt(value)
      else data = value
      return data
    }
    for (const i in keys) {
      const current = keys[i]
      const value = document.querySelector(`#${current}`).value

      const data = setValueMatchDataType(product[current], value)
      product[current] = data
    }
    return product
  }

  return {
    extractProduct: () => extractProduct()
  }
}
