.SILENT:



test-cmd:
	# runs cmd's tests
	cd cmd && go test -cover

test-error:
	# runs error's tests
	cd error-h && go test -cover

test-utils:
	# runs utils' tests
	cd utils && go test -cover

test-ternary:
	# runs ternary's tests
	cd ternary && go test -cover
