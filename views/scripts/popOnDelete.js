const confirmModal = document.getElementById('confirm-modal')
const cancelBtn = document.getElementById('cancel')
const confirmBtn = document.getElementById('confirm')

let storyId

document.querySelectorAll("[aria-label='delete']").forEach(btn => {
    btn.addEventListener('click', (e) => {
        storyId = e.target.id
        confirmModal.classList.replace("hidden", "block")
    })
})

cancelBtn.addEventListener('click', () => {
    confirmModal.classList.replace("block", "hidden")
})

confirmBtn.addEventListener('click', async () => {
    console.log(storyId)

    const res = await fetch(`/story/delete/${storyId}`, {
        method: "POST"
    })
    const data = await res.json()
    console.log(data)

    window.location.reload()
})

