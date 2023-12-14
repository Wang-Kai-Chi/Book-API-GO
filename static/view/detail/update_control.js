import IknowToken from '../iknow_token.js'
import NodeScriptReplace from '../node_script_replace.js'
import ProductFormExtractor from '../product_form_extractor.js'

export default function UpdateControl (iknowToken = IknowToken()) {
  const updateBtn = document.querySelector('#updateBtn')
  const confirmBtn = document.querySelector('#confirmUpdateBtn')
  const cancelBtn = document.querySelector('#cancelUpdateBtn')

  const updateController = UpdateController(iknowToken)
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

function UpdateController (iknowToken = IknowToken()) {
  const form = document.querySelectorAll('.form-control')

  const enableUpdate = () => {
    for (const f of form) { f.disabled = false }
  }

  const cancelUpdate = () => {
    for (const f of form) { f.disabled = true }
  }

  const confirmUpdate = async () => {
    const token = (iknowToken.json() === null)
      ? ''
      : 'Bearer ' + iknowToken.json().Token

    const handleResponse = (res, success = () => { }) => {
      const d = res.json()
      if (res.status === 200) {
        success()
        return d
      } else if (res.status === 401) {
        const reverify = confirm('驗證已過期，將重新驗證')

        if (reverify) {
          fetch('/static/view/auth/auth.html').then(res => res.text())
            .then(data => {
              document.body.innerHTML = data
              NodeScriptReplace(document.body)
            })
            .catch(err => console.log(err))
        }
        return d.then(Promise.reject.bind(Promise))
      } else {
        alert('驗證失敗, 請登入')
        return d.then(Promise.reject.bind(Promise))
      }
    }

    fetch('/api/v1/product/update', {
      method: 'PUT',
      body: JSON.stringify([ProductFormExtractor().extractProduct()]),
      headers: new Headers({
        'Content-Type': 'application/json',
        Authorization: token
      })
    }).then(res => handleResponse(res, () => {
      const banner = document.querySelector('.alert')
      banner.hidden = false
      const alertText = document.querySelector('#alertText')
      alertText.innerHTML = '更新成功'
    })).catch(err => console.log(err))
  }

  return {
    enableUpdate: () => enableUpdate(),
    cancelUpdate: () => cancelUpdate(),
    confirmUpdate: () => confirmUpdate()
  }
}
