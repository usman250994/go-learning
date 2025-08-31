module hello

go 1.25.0

require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	greetings v0.0.0
	rsc.io/sampler v1.3.0 // indirect
)

replace greetings => ../greetings
