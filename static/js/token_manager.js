import NodeScriptReplace from './request/node_script_replace.js'

export default function TokenManager () {
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

  return {
    handleAuthurizationExpired: () => handleAuthurizationExpired()
  }
}
