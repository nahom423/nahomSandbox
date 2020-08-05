package main
import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)
type ContactList struct { //creating struct
	id int
	name string
	address string
}

func db() (db *sql.DB){ //establishing db connection
	dbDriver := "mysql"
	dbUser:= "nnegash"
	dbPass:= "1234"
	dbName:= "IFM_Contact_List"
	db, err := sql.Open(dbDriver, dbUser + ":" +dbPass+"@/"+dbName )
	if err!= nil{
		panic(err.Error())
	}
	return db
}
func main() {
	
}
