const statbar = document.getElementById('status')

const urlParams = new URLSearchParams(window.location.search)
const state = urlParams.get('status')

let statMessage = createNotification(state)

statbar.append(statMessage)

function createNotification(success) {
    const div = document.createElement('div')
    const p = document.createElement('p')

    p.innerText = ""

    if (success === "success") {
        div.classList.add("bg-green-400")
        p.innerText = "login success"
    }

    if (success === "failed") {
        div.classList.add("bg-red-400")
        p.innerText = "invalid credentials"
    }

    div.append(p)
    return div
}
