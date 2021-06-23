module github.com/achelovekov/telemetry-collector

go 1.15

replace github.com/achelovekov/collectorutils => /Users/achelove/Go/src/github.com/achelovekov/collectorutils
require (
	github.com/achelovekov/collectorutils v0.0.0-20210401112550-6f70067e1724
	github.com/elastic/go-elasticsearch v0.0.0 // indirect
	go.mongodb.org/mongo-driver v1.5.1
)
