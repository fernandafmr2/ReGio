package store

import(
	"log"
	"github.com/go-pg/pg/v10"
)

// database connector
var db *pg.DB

func SetDBConnection(dbOpts *pg.Options){
	if dbOpts == nil {
		log.Panicln("DB options can't be nil")
	} else {
		db = pg.Connect(dbOpts)
	}
}

func GetDBConnetion() *pg.DB { return db }