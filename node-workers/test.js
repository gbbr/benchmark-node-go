var CONNECTIONS = 100,
	net = require("net"),
	completed = 0;

for (var i = 0; i < CONNECTIONS; i++) {
	var client = net.createConnection({ port: 1234 })
	client.on("data", function (data) {
		if (data.toString() == "ERROR") {
			process.stdout.write("Received error.")
			process.exit(1);
		} else {
			process.stdout.write(data.toString());
			if (++completed === CONNECTIONS) {
				process.exit(0);
			}
		}
	});
}
