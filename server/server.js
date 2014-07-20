var net = require("net")

net.createServer(function (conn) {
	setTimeout(function () {
		conn.write("X");
	}, 1000);
	return;
}).listen(1234);