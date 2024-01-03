import UserController from './controller/user_controller.js'
import User from './data/user.js'
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

      const loadingMessage = document.querySelector('#loadingMessage')
      const returnHome = document.querySelector('#returnHome')
      returnHome.onclick = () => {
        location.reload()
      }

      const userStr = JSON.stringify(user)

      const sendMail = () => {
        EmailController().sendVerificationMail(userStr)
          .then(() => {
            returnHome.hidden = false
            loadingMessage.innerHTML = '驗證電子郵件已寄出，請至電子信箱查看'
          }).catch(err => {
            console.log(err)
            returnHome.hidden = false
            loadingMessage.innerHTML = '驗證失敗，請重新註冊'
          })
      }

      UserController().addUser(userStr)
        .then(() => sendMail())
    }
  }
}
