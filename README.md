### NodeJS vs Go Concurrent Network Performance Benchmark

Test scenario is as follows:

A TCP server (written in Go, and compiled) is provided, which opens port 1234 and welcomes connections. Each client that connects gets sent the byte value of the character "X" after a 1 second wait time, after which the client gets disconnected.

To start the server run the `server` binary that is located inside the server folder. A NodeJS version of the server is also provided.

After the server is started, run the tests using `node testRunner.js`. Number of connections in test can be changed by altering the `CONNECTIONS` constant inside each of the tests. There is also a limit on the OS as to how many TCP connections area allowed at any given time, to change this, see: http://stackoverflow.com/questions/7578594/how-to-increase-limits-on-sockets-on-osx-for-load-testing

Some results I've obtained for 5000 connections are below:

| Go        | NodeJS           |
| ------------- |:-------------:|
| 1.076 sec      | 1.356 sec |
| 1.072 sec      | 1.352 sec      |
| 1.068 sec      | 1.682 sec      |
| 1.047 sec      | 5.487 sec      |
