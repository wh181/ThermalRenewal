package ishttpa

var web_str = "<script>var ws = new WebSocket(\"ws://localhost:8080/conn\");\nws.onopen = function (evt) {\n    ws.send(\"ping\")\n};\nws.onmessage = function (event) {\n    if(event.data == \"reload\"){\n        window.location.reload()\n    }\n};</script>"