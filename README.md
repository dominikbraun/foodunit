# Prerequisites
- yarn (plain npm may also work, not tested)

# How to run
First start a webserver which serves the client.
You can use the fileserver provided by foodunit or serve the files via an external webserver like nginx or (for development) yarn serve.

The go server can also proxy a webserver serving the client files.
This is needed if you run the client in debug mode via yarn. The proxy prevents a cross-origin error your webbrowser may throw.