module example-go/niksolaz/newmod

go 1.22.4

replace example-go/niksolaz/greetings => ../greetings

// Remove the unnecessary require statement

require example-go/niksolaz/greetings v0.0.0-00010101000000-000000000000 // indirect
