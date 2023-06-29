package main

// _ antes do pacote (por exemplo, _ github.com/lib) significa que será utilizada em tempo de execução
import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Login struct {
	User     string
	Password string
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func connectBD() *sql.DB {
	connection := "user=postgres dbname=kien_users password=sql123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	db := connectBD()
	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

	// site := "https://google.com"
	// response, error := http.Get(site)

	// // fmt.Println(response)
	// if response.StatusCode == 200 {
	// 	fmt.Println("!")
	// } else {
	// 	fmt.Println("Satus code: ", response.StatusCode)
	// 	fmt.Println(error)
	// }
}

func index(w http.ResponseWriter, r *http.Request) {
	logins := []Login{
		{User: "carlos1", Password: "123"},
		{User: "carlos2", Password: "456"},
		{User: "rodrigo", Password: "123"},
		{User: "marina", Password: "123"},
		{User: "jummy", Password: "123"},
		{User: "takeshi", Password: "123"},
	}

	templates.ExecuteTemplate(w, "Index", logins)
}
