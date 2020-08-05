package main
import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)
type ContactList struct { //creating struct
	Id int
	Name string
	Address string
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

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request){
	db := dbConn()
	selDB, err := db.Query("Select * From Contact_List ORDER BY id DESC" )
	if err !nil {
		panic(err.Error())
	}
	list := ContactList{}
	res := []ContactList{}

	for selDB.Next(){
		var id int
		var name, address string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		list.Id = id
		list.Name = name
		list.Address = address
		res = append (res, emp)
	}
	tmpl.ExecuteTemplate(w,"Index", res)
	defer db.Close()
}
func main() {

}
