import NodeScriptReplace from './node_script_replace.js'

export default function ResponseHandler () {
  const handleAuthurizationExpired = () => {
    const reverify = confirm('驗證已過期，將重新驗證')

    if (reverify) {
      fetch('/static/view/auth.html').then(res => res.text())
        .then(data => {
          document.body.innerHTML = data
          NodeScriptReplace(document.body)
        })
        .catch(err => console.log(err))
    }
  }

  /**
   *
   *
   * @param {*} [res=Promise()]
   * @param {*} [success=() => { }]
   * @return {Promise().json()}
   */
  const handleResponse = (res = Promise(), success = () => {}) => {
    const d = res.json()
    if (res.status === 200) {
      success()
      return d
    } else if (res.status === 401) {
      handleAuthurizationExpired()
      return d.then(Promise.reject.bind(Promise))
    } else if (res.status === 400) {
      alert('資料錯誤，請再次確認')
      return d.then(Promise.reject.bind(Promise))
    } else {
      alert('系統錯誤，請重新操作')
      return d.then(Promise.reject.bind(Promise))
    }
  }

  return {
    run: (res = Promise(), success = () => {}) => handleResponse(res, success)
  }
}
