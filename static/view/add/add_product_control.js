setTimeout(() => AddProductControl(IknowToken()), 150)
function AddProductControl (iknowToken = IknowToken()) {
  const token = (iknowToken.json() === null)
    ? ''
    : 'Bearer ' + iknowToken.json().Token

  const confirmAddProduct = () => {
    const add = async (body) => {
      return fetch('/api/v1/product/insert', {
        method: 'POST',
        body: JSON.stringify(body),
        headers: new Headers({
          'Content-Type': 'application/json',
          Authorization: token
        })
      }).then(res => {
        const d = res.json()
        if (res.status === 200) {
          return d
        } else {
          alert('驗證失敗, 請登入或重新取得驗證碼')
          return d.then(Promise.reject.bind(Promise))
        }
      })
    }

    add([ProductFormExtractor().extractProduct()])
      .then(() => alert('新增成功'))
      .catch(err => console.log(err))
  }
  document.querySelector('#confirmAdd').onclick = () => {
    confirmAddProduct()
  }
}
