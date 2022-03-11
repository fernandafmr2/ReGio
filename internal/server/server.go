package server

import(
	"ReGio/internal/database"
	"ReGio/internal/store"
	"ReGio/internal/conf"
)

func Start(cfg conf.Config) {
	store.SetDBConnection(database.NewDBOptions(cfg))

	router := setRouter()

	// start listening
	router.Run(":8080")
}