module github.com/ozoncp/ocp-note-api

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/Masterminds/squirrel v1.5.0
	github.com/Shopify/sarama v1.29.0
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jmoiron/sqlx v1.3.4
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.2.0
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.16.3
	github.com/onsi/gomega v1.13.0
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.22.0
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.8.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/genproto v0.0.0-20210426193834-eac7f76ac494
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/ozoncp/ocp-note-api/pkg/ocp-note-api => ./pkg/ocp-note-api
