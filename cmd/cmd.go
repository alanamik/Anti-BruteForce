package main

import (
	"OTUS_hws/Anti-BruteForce/internal/config"
	"OTUS_hws/Anti-BruteForce/internal/gen/restapi"
	"OTUS_hws/Anti-BruteForce/internal/gen/restapi/operations"
	"OTUS_hws/Anti-BruteForce/internal/handlers"
	"OTUS_hws/Anti-BruteForce/internal/redisdb"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-openapi/loads"
	"github.com/pkg/errors"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	conf, err := config.New()
	if err != nil {
		err = errors.Wrap(err, "[config.New()]")
		err = errors.Wrap(err, "[main()]")

		panic(err)
	}

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		err = errors.Wrap(err, "[loads.Analyzed()]")
		err = errors.Wrap(err, "[main()]")
		panic(err)
	}

	ctx := context.Background()
	redisClient := redisdb.NewClient(*conf)
	h := handlers.NewHandler(redisClient)

	api := operations.NewAntiBrutForceAPI(swaggerSpec)
	h.Register(api)
	server := restapi.NewServer(api)
	server.Port = conf.Service.Port
	server.Host = conf.Service.Host

	//	err = redisClient.AddToList(ctx, "128.10.10.10/10", redisDB.Blacklist)
	//	fmt.Println(err)
	err = redisClient.DeleteFromList(ctx, "128.10.10.10/10", redisdb.Blacklist)
	fmt.Println(err)

	// Сёрвим сервис на наличие panic ошибок
	if err = server.Serve(); err != nil {
		err = errors.Wrap(err, "[server.Serve()]")
		err = errors.Wrap(err, "[main()]")
		panic(err)
	}

	// До тех пор, пока не будет прокинут системый shutdown, канал будет залочен
	sig := <-quit

	err = redisClient.Client.Close()
	if err != nil {
		err = errors.Wrapf(err, "[db.Close(%v)]", sig)
		err = errors.Wrap(err, "[main()]")
		panic(err)
	}

	// Завершаем сеанс
	err = server.Shutdown()
	if err != nil {
		err = errors.Wrap(err, "[server.Shutdown()]")
		err = errors.Wrap(err, "[main()]")
		panic(err)
	}
}
