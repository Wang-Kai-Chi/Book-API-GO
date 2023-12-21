import NodeScriptReplace from '../node_script_replace.js'

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
  const handleResponse = (res = Promise(), success = () => { }) => {
    const d = res.json()
    if (res.status === 200) {
      success()
      return d
    } else if (res.status === 401) {
      handleAuthurizationExpired()
      return d.then(Promise.reject.bind(Promise))
    } else {
      alert('驗證失敗, 請登入')
      return d.then(Promise.reject.bind(Promise))
    }
  }

  return {
    run: (res = Promise(), success = () => {}) => handleResponse(res, success)
  }
}
