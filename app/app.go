package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"webapi/app/handler"
)

//App is ...
type App struct {
	DbName  string
	Address string
}

//Init is ...
func Init() App {
	byteValue, _ := ioutil.ReadFile("app.json")
	var app App
	json.Unmarshal(byteValue, &app)
	return app
}

//Start is ...
func (app *App) Start() {
	fs := http.FileServer(http.Dir("./client"))
	sc := handler.StudentController{DbName: app.DbName}

	http.Handle("/", fs)
	http.HandleFunc("/api/data", sc.Route)
	http.ListenAndServe(app.Address, nil)
	log.Printf("Listening on %s", app.Address)
	log.Fatal(http.ListenAndServe(app.Address, nil))
	fmt.Printf("Start")
}
