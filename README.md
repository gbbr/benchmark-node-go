benchmark-node-go
=================

Benchmarking Concurrent Network Performance

Server opens TCP port 1234 and replies to any connection with "X" after 1 second wait time.

Start server by running the `server/server` binary (or `node server/server.js` to use the NodeJS server)

After the server is running, run the tests using `node testRunner.js`

Some results I've obtained for 5000 connections are below:

| NodeJS        | Go           |
| ------------- |:-------------:|
| 1.076 sec      | 1.356 sec |
| 1.072 sec      | 1.352 sec      |
| 1.068 sec      | 1.682 sec      |
| 1.047 sec      | 5.487 sec      |
