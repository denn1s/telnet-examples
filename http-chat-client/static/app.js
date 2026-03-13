const messagesDiv = document.getElementById("messages")
const form = document.getElementById("form")

async function loadMessages() {
    const resp = await fetch("/api/messages")
    const messages = await resp.json()

    messagesDiv.innerHTML = ""
    for (const msg of messages) {
        const div = document.createElement("div")
        div.className = "message"
        div.innerHTML = `<span class="user">${msg.user}:</span> ${msg.text}`
        messagesDiv.appendChild(div)
    }

    messagesDiv.scrollTop = messagesDiv.scrollHeight
}

form.addEventListener("submit", async (e) => {
    e.preventDefault()

    const user = form.user.value
    const text = form.text.value

    await fetch("/api/messages", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ user, text }),
    })

    form.text.value = ""
    await loadMessages()
})

loadMessages()
setInterval(loadMessages, 2000)
