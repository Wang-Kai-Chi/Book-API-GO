export default function Register () {
  const User = () => {
    const user = {
      Id: '',
      Name: '',
      Email: '',
      Phone: '',
      Password: ''
    }
    return {
      this: () => { return user },
      keys: () => { return Object.keys(user) }
    }
  }
  const component = (selector = '') => { return document.querySelector(selector) }

  const submit = component('#submit')

  submit.onclick = () => {
    const user = User().this()

    if (component('#password').value !== component('#confirmpw').value) {
      const alertText = component('#alertText')
      alertText.innerHTML = '確認密碼錯誤'

      const alert = component('.alert')
      alert.hidden = false
    } else {
      for (const k of User().keys()) {
        const compV = (component(`#${k.toLowerCase()}`) === null)
          ? ''
          : component(`#${k.toLowerCase()}`).value
        user[k] = compV
      }

      fetch('/api/v1/user/insert', {
        method: 'POST',
        body: JSON.stringify(user),
        headers: new Headers({
          'Content-Type': 'application/json'
        })
      }).then(res => {
        const d = res.json()
        if (res.status === 200) {
          return d
        } else {
          console.log('Register failed')
          return d.then(Promise.reject.bind(Promise))
        }
      }).then(() => location.reload())
        .catch(err => console.log(err))
    }
  }
}
