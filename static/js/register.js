import UserController from './controller/user_controller.js'
import User from './data/user.js'
import UserInfo from './localstorage/user_info.js'
import EmailController from './controller/email_controller.js'

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

      const userStr = JSON.stringify(user)

      UserController().addUser(userStr).then(data => {
        UserInfo().set(JSON.stringify(data))
        EmailController().sendVerificationMail(userStr)
      })
    }
  }
}
