var ws = new WebSocket("ws://localhost:8080/conn");
ws.onopen = function (evt) {
    ws.send("ping")
};
ws.onmessage = function (event) {
    if(event.data == "reload"){
        window.location.reload()
    }
};
