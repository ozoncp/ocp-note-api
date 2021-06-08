package core

//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-note-api/core/flusher Flusher
//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-note-api/core/repo Repo
//go:generate mockgen -destination=./mocks/alarmer_mock.go -package=mocks github.com/ozoncp/ocp-note-api/core/alarmer Alarmer
//go:generate mockgen -destination=./mocks/saver_mock.go -package=mocks github.com/ozoncp/ocp-note-api/core/saver Saver
