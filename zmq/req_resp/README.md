Tested the following cases:

* Initialization
It does not matter if server starts before client or client before server.
They both block until REQ - RESP is completed

* Blocking
When the client makes a request, blocks until the server replies
Server block until receives request

* Restarting
Client can restart and this does not affect server.
If server restarts client need to restart as well.
