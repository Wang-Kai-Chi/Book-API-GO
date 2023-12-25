import UserController from './controller/user_controller.js'
import User from './data/user.js'

export default function Register () {
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

      UserController().addUser(JSON.stringify(user))
    }
  }
}
