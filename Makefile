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

PHONY: .vendor-proto
.vendor-proto:
		mkdir vendor.protogen
		mkdir vendor.protogen\api\ocp-note-api
		copy /y api\ocp-note-api\ocp-note-api.proto vendor.protogen\api\ocp-note-api

		git clone https://github.com/googleapis/googleapis vendor.protogen\googleapis
		mkdir vendor.protogen\google
		move vendor.protogen\googleapis\google\api vendor.protogen\google
		rmdir /s /q vendor.protogen\googleapis

		mkdir vendor.protogen\validate
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen\protoc-gen-validate
		move vendor.protogen\protoc-gen-validate\validate\*.proto vendor.protogen\validate
		rmdir /s /q vendor.protogen\protoc-gen-validate

		mkdir vendor.protogen\google\protobuf
		git clone https://github.com/protocolbuffers/protobuf vendor.protogen\protobuf
		move vendor.protogen\protobuf\src\google\protobuf\*.proto vendor.protogen\google\protobuf
		rmdir /s /q vendor.protogen\protobuf