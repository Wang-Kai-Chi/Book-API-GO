import UserController from './controller/user_controller.js'
import UserInfo from './localstorage/user_info.js'

export default function Auth () {
  const submit = () => {
    const password = document.querySelector('#password')

    const user = UserInfo().json()
    user.Password = password.value

    UserController().login(JSON.stringify(user))
  }

  document.querySelector('#submit').onclick = () => submit()
}
