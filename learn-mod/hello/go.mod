module example-go/niksolaz/hello

go 1.22.4

replace example-go/niksolaz/greetings => ../greetings

require (
	example-go/niksolaz/greetings v0.0.0-00010101000000-000000000000
	example-go/niksolaz/messages v0.0.0-00010101000000-000000000000
)

replace example-go/niksolaz/messages => ../messages
