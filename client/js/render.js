function login() {
    fetch("/components/login.html").then(rs => rs.text()).then(rs => {
        document.getElementById("root").innerHTML = rs
    })
}

submit = () => {
    event.preventDefault()
    console.log(event.target)
    return false;
}