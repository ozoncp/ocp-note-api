module github.com/ozoncp/ocp-note-api

go 1.16

require (
	github.com/Masterminds/squirrel v1.5.0
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
	github.com/rs/zerolog v1.22.0
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	google.golang.org/genproto v0.0.0-20210426193834-eac7f76ac494
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/ozoncp/ocp-note-api/pkg/ocp-note-api => ./pkg/ocp-note-api
