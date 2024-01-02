import UserController from './controller/user_controller.js'
import UserInfo from './localstorage/user_info.js'

export default function VerifySuccess () {
  const verifyUser = () => {
    const message = document.querySelector('#message')

    const getSuccessMsg = () => {
      return `
    <h2>Email Verify Successful</h2>
    <p>You can close this page, and return to website now.</p>
        `
    }

    UserController().authurize(UserInfo().string(), function () {
      message.innerHTML = getSuccessMsg()
      UserInfo().remove()
    })
  }

  verifyUser()
}
