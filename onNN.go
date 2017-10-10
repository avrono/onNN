package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
	"github.com/onNN/routes"
	"github.com/onNN/models"
)


// routes
func handleRequests(env *routes.Env) {
	// this is just for when browser requests favicon
	http.HandleFunc("/favicon.ico", routes.HandlerICon)
	http.HandleFunc("/neurons", env.AddNeuron)
	http.HandleFunc("/build", env.Build)
	http.HandleFunc("/test", routes.TestTemplate)
	http.HandleFunc("/css/", routes.ServeResource)
	log.Fatal(http.ListenAndServe(":3334", nil))
}

// log
func mutter(s string) {
	fmt.Print(time.Now().Format(time.RFC850))
	fmt.Print(" - ")
	fmt.Println(s)
}

// main
func main() {
	//Create a new DB Handle TODO: Support more that one DB
	db, err := models.NewDB("MGO_DIAL")
	if err != nil {
		panic(err)
	}

	/* Can create Env with a DB obj, since DB class implements The Interface methods */
	env := &routes.Env{db}

	cwd, _ := os.Getwd()
	fmt.Println(filepath.Join(cwd, "./templates/todos.html"))
	handleRequests(env)
	defer db.Close()
}
