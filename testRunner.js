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

	exec(tests[number].cmd, function (err, stdout, stderr) {
		
			console.log(stdout);
		
		if (number < tests.length - 1) {
			nextTest(++number);
		}
	});
})(0)
