var exec = require("child_process").exec,
	tests = [{
		name: "Go-Lang using goroutines: ",
		cmd: "go-routines/test"
	}, {
		name: "Node Async: ",
		cmd: "node node-workers/test.js"
	}];

(function nextTest(number) {
	process.stdout.write(tests[number].name);
	var startTime = new Date().getTime();

	exec(tests[number].cmd, function (err, stdout, stderr) {
		console.log((new Date().getTime() - startTime) / 1000 + "s");
		if (number < tests.length - 1) {
			nextTest(++number);
		}
	});
})(0)
