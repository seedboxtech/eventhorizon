module github.com/seedboxtech/eventhorizon

require (
	cloud.google.com/go v0.26.0
	github.com/aws/aws-sdk-go v1.15.26
	github.com/cenkalti/backoff v2.0.0+incompatible // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/globalsign/mgo v0.0.0-20180828104044-6f9f54af1356
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75
	github.com/gorilla/websocket v1.4.0
	github.com/guregu/dynamo v1.0.0
	github.com/jpillora/backoff v0.0.0-20170918002102-8eab2debe79d
	github.com/kr/pretty v0.1.0
	github.com/looplab/eventhorizon v0.4.0
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.2.2
	golang.org/x/net v0.0.0-20180925072008-f04abc6bdfa7 // indirect
)

replace github.com/looplab/eventhorizon => ./
