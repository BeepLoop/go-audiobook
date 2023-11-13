const form = document.getElementById('story-form')
const submitBtn = document.getElementById('submit')

form.addEventListener('submit', () => {
    submitBtn.disabled = true
    submitBtn.innerText = "submitting..."
})
