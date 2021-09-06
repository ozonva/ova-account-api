package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonva/ova-account-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozonva/ova-account-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/producer_mock.go -package=mocks github.com/ozonva/ova-account-api/internal/kafka Producer
//go:generate mockgen -destination=./mocks/metrics_mock.go -package=mocks github.com/ozonva/ova-account-api/internal/metrics AccountMetrics
