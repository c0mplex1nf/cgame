package deck
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DeckRepo struct {
	db *sql.Db
}

func list(){
	
}