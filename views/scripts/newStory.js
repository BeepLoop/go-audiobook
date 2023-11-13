const statbar = document.getElementById('status')

const urlParams = new URLSearchParams(window.location.search)
const state = urlParams.get('status')
const reason = urlParams.get('reason')

let statMessage
if (state == "failed") {
    statMessage = createNotification("failed")
} else if (state == "success") {
    statMessage = createNotification("success")
} else {
    statMessage = createNotification(undefined)
}

statbar.append(statMessage)

function createNotification(success) {
    const div = document.createElement('div')
    const p = document.createElement('p')
    const rsn = document.createElement('p')

    if (success == "success") {
        div.classList.add("bg-green-400")
        p.innerText = "story saved successfully"
    } else if (success == "failed") {
        div.classList.add("bg-red-400")
        p.innerText = "failed to save story"
        rsn.innerText = reason
    } else {
        p.innerText = ""
    }

    div.append(p)
    div.append(rsn)
    return div
}
