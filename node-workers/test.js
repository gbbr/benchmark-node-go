var CONNECTIONS = 200,
	net = require("net"),
	completed = 0,
	startTime = new Date().getTime();

for (var i = 0; i < CONNECTIONS; i++) {
	var client = net.createConnection({ port: 1234 })
	client.on("data", function (data) {
		if (data.toString() == "ERROR") {
			process.stdout.write("Received error.")
			process.exit(1);
		} else {
			process.stdout.write(data.toString());
			if (++completed === CONNECTIONS) {
				console.log("Time: ", (new Date().getTime() - startTime) / 1000 + "s");
				process.exit(0);
			}
		}
	});
}
