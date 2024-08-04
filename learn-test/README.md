# GO test

Running the go test command in a module or package directory will automatically run all tests defined in that package's _test.go files.
```
go test
```

-v (verbose): Print detailed information when running tests.
```
go test -v
```

-run: Run only tests that match a specific pattern.
```
go test -run TestXxx
```

-bench: Run benchmarks that match a specific pattern.
```
go test -bench=.
```

-cover: Generate a code coverage report.
```
go test -cover
```

-race: Detect race conditions in the code.
```
go test -race
```


