benchmark-node-go
=================

Benchmarking Concurrent Network Performance

Server opens TCP port 1234 and replies to any connection with "X" after 1 second wait time.

Start server by running the `server/server` binary if you wish to use the Go server, or `node server/server.js` to use the NodeJS server.

After the server is running, run the tests using `node testRunner.js`
