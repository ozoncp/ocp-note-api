module github.com/ozoncp/ocp-note-api

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/golang/mock v1.5.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.16.3
	github.com/onsi/gomega v1.13.0
	github.com/rs/zerolog v1.22.0
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	google.golang.org/genproto v0.0.0-20210426193834-eac7f76ac494
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)

replace github.com/ozoncp/ocp-note-api/pkg/ocp-note-api => ./pkg/ocp-note-api
