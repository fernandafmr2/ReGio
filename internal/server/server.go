package server

import(
	"ReGio/internal/database"
	"ReGio/internal/store"
)

func Start() {
	store.SetDBConnection(database.NewDBOptions())

	router := setRouter()

	// start listening
	router.Run(":8080")
}