package config

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	dsn := "mongodb://" + host + ":" + port

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatal().Msgf("error connecting to mongo: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal().Msgf("error pinging mongo: %v", err)
	}

	log.Info().Msg("connected to mongo ok!")

	dbName := os.Getenv("MONGO_DATABASE")
	return client.Database(dbName)
}
