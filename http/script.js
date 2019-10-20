let
    ws,
    loopCount = 0,
    messagesCount = 0,

    status = document.querySelector(".status"),
    statusUsers = document.querySelector(".status-users"),
    statusMessages = document.querySelector(".status-messages"),
    statusText = document.querySelector(".status-text"),

    ths = document.querySelectorAll(".log th"),
    logElem = document.querySelector(".log tbody")

function log(time, type, value) {
    let tr = document.createElement("tr")
    tr.innerHTML = `<td>${time}</td><td>${type}</td><td title="${value}">${value}</td>`
    logElem.appendChild(tr)

    for (let i = 0; i < ths.length; i++) {
        ths[i].width = tr.childNodes[i].offsetWidth + "px"
    }
    logElem.scrollTop = logElem.scrollHeight
}

function loop() {
    if (loopCount === 20) {
        statusText.innerHTML = "Please reload page"
        return
    }
    loopCount++
    statusText.innerHTML = "Connect..."
    ws = new WebSocket("ws://192.168.1.37:80/ws")
    ws.onopen = () => {
        status.classList.add("online")
        statusText.innerHTML = "Connection established"
    }

    ws.onmessage = (event) => {
        messagesCount++
        statusMessages.innerHTML = messagesCount
        let data = JSON.parse(event.data)
        console.log("Incoming message", data)
        if ("Users" in data) {
            statusUsers.innerHTML = data.Users
        }
    }

    ws.onclose = (event) => {
        if (event.wasClean) {
            statusText.innerHTML = "Connection close"
        } else {
            statusText.innerHTML = "Connection abort"
        }
        status.classList.remove("online")
        setTimeout(loop, 2500)
    }
}

loop()
