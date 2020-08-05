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

func dbConn() (db *sql.DB){ //establishing db connection
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
	if err != nil {
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

func Show(w http.ResponseWriter, r *http.Request){
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("Select * From Contact_List wwhere id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	list := ContactList{}
	for selDB.Next(){
		var id int
		var name, address string
		err = selDB.Scan(&id, &name, &address)

		if err != nil {
			panic(err.Error())
		}
		list.Id = id
		list.Name = name
		list.Address = address
	}
	tmpl.ExecuteTemplate(w,"Show", list)
	defer db.Close()
}
func New(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit (w http.ResponseWriter, r *http.Request){
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("Select * from Contact_List where id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	list := ContactList{}
	for selDB.Next(){
		var id int
		var name, address string
		err = selDB.Scan(&id, &name, &address)
		if err != nil{
			panic(err.Error())
		}
		list.Id = id
		list.Name = name
		list.Address = address
	}
	tmpl.ExecuteTemplate(w, "Edit", list)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request)  {
	db := dbConn()
	if r.Method == "POST"{
		name := r.FormValue("name")
		address:= r.FormValue("address")
		insForm, err := db.Prepare("INSERT INTO Contact_List(name, address) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, address)
		log.Println("INSERT: Name: " + name + "| Address: "+ address)
	}

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request){
	db := dbConn()
	if r.Method == "POST"{
		name := r.FormValue("name")
		address := r.FormValue("address")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("Update Contact_List set name=?, city=? where id=?")
		if err != nil{
			panic(err.Error())
		}
		insForm.Exec(name,address,id)
		log.Println("Update: Name: "+ name + " | address: " + address)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)

}

func Delete (w http.ResponseWriter, r *http.Request){
	db := dbConn()
	list := r.URL.Query().Get("id")
	delForm, err := db.Prepare("Delete From Contact_List Where id=?")
	if err != nil{
		panic(err.Error())
	}
	delForm.Exec(list)
	log.Println("Delete")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}