export default function User () {
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
