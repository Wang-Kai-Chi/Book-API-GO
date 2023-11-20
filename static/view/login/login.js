
function Login() {
    const User = () => {
        let user = {
            Id: "",
            Name: "",
            Email: "",
            Phone: "",
            Password: "",
        }
        return {
            this: () => { return user },
            keys: () => { return Object.keys(user) },
        }
    }

    const submit = () => {
        const email = document.querySelector("#email")
        const password = document.querySelector("#password")
        const user = User()

        user.this()["Email"] = email.value
        user.this()["Password"] = password.value

        fetch("/api/v1/user/login", {
            method: "POST",
            body: JSON.stringify(user.this()),
            headers: new Headers({
                "Content-Type": "application/json",
            }),
        }).then(res => {
            let d = res.json()
            if (res.status === 200) {
                return d
            } else {
                return d.then(Promise.reject.bind(Promise));
            }
        }).then(data => {
            let user = data[0]
            user.Password = ""
            localStorage.setItem("userinfo", JSON.stringify(user))
            window.location.href = '/'
        }).catch(err => console.log(err))

    }

    document.querySelector("#submit").onclick = () => submit()

}