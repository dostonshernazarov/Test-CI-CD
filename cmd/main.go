package main

import (
	"github/test-ci-cd-github/api"
	"github/test-ci-cd-github/config"
	"github/test-ci-cd-github/pkg/logger"

	"github/test-ci-cd-github/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	// casbin with postgres -----------------------------------
	// psqlString := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`, "localhost", 5432, "postgres", "1234", "backend")
	// db, err := gormadapter.NewAdapter("postgres", psqlString, true)
	// if err != nil {
	// 	log.Error("gormadapter error", logger.Error(err))
	// 	return
	// }
	// casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, db)
	// if err != nil {
	// 	log.Error("NewEnforcer error", logger.Error(err))
	// 	return
	// }

	// Kafka ----
	// writer, err := producer.NewKafkaProducerInit([]string{"localhost:9092"})
	// if err != nil {
	// 	log.Error("NewKafkaProducerInit: %v", logger.Error(err))
	// }

	// err = writer.ProduceMessage("test-topic", []byte("message"))
	// if err != nil {
	// 	log.Fatal("failed to run http server", logger.Error(err))
	// }

	// defer writer.Close()

	// -----------------

	apiServer := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := apiServer.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
