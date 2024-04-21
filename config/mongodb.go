package config

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongo() *MongoConfig {
	log.Info().Msgf("connecting to mongo... 🕰️")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	dsn := "mongodb://" + host + ":" + port

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatal().Msgf("error connecting to mongo: %v", err)
	}

	log.Info().Msgf("pinging mongo... 👋")

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal().Msgf("error pinging mongo: %v", err)
	}

	log.Info().Msg("connected to mongo ok! 🚀")

	dbName := os.Getenv("MONGO_DATABASE")

	return &MongoConfig{
		Client: client,
		DB:     client.Database(dbName),
	}
}
