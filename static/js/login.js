import UserController from './controller/user_controller.js'
import User from './user.js'

export default function Login () {
  const submit = () => {
    const email = document.querySelector('#email')
    const password = document.querySelector('#password')
    const user = User()

    user.this().Email = email.value
    user.this().Password = password.value

    UserController().login(JSON.stringify(user.this()))
  }

  document.querySelector('#submit').onclick = () => submit()
}
