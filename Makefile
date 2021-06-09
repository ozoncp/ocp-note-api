PHONY: .generate
.generate:
		mkdir swagger
		mkdir pkg\ocp-note-api
		protoc -I ./vendor.protogen -I ./api/ocp-note-api \
				--go_out=pkg/ocp-note-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-note-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-note-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-note-api \
				--swagger_out=allow_merge=true,merge_file_name=api:. \
				./api/ocp-note-api/ocp-note-api.proto
		move pkg\ocp-note-api\github.com\ozoncp\ocp-note-api\pkg\ocp-note-api\* pkg\ocp-note-api
		rmdir /s /q pkg\ocp-note-api\github.com