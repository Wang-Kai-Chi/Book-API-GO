import UserController from './controller/user_controller.js'

export default function Login () {
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
