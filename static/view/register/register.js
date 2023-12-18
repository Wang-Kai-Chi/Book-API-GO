Register()

function Register () {
  const component = (selector = '') => { return document.querySelector(selector) }

  const userName = component('#user')
  const email = component('#email')
  const password = component('#password')
  const confirmpw = component('#confirmpw')
  const submit = component('#submit')

  const alertText = component('#alertText')
  const alert = component('.alert')

  submit.onclick = () => {
    console.log(userName.value)
    console.log(email.value)
    console.log(password.value)
    console.log(confirmpw.value)

    if (password.value !== confirmpw.value) {
      alertText.innerHTML = '確認密碼錯誤'
      alert.hidden = false
    }
  }
}
