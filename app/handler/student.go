package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"webapi/app/dao"
	"webapi/app/entity"
)

//StudentController is ...
type StudentController struct {
	DbName string
}

func (sc *StudentController) get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.Join(r.URL.Query()["id"], ""))
	var d = dao.DAO{Name: sc.DbName}
	student := d.Get(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{name: %s,age:%d}", student.Name, student.Age)))
}

func (sc *StudentController) post(w http.ResponseWriter, r *http.Request) {
	student := entity.Student{}
	json.NewDecoder(r.Body).Decode(&student)
	var d = dao.DAO{Name: sc.DbName}
	err := d.Add(student)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{mss: %s}", err.Error())))
	} else {
		w.Write([]byte(fmt.Sprintf("{mss: 'success'}")))
	}
}

func (sc *StudentController) del(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.Join(r.URL.Query()["id"], ""))
	var d = dao.DAO{Name: sc.DbName}
	err := d.Del(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{mss: %s}", err.Error())))
	} else {
		w.Write([]byte(fmt.Sprintf("{mss: 'success'}")))
	}
}

func (sc *StudentController) notfound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

//Route is ..
func (sc *StudentController) Route(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		sc.get(w, r)
	case "POST":
		sc.post(w, r)
	case "DELETE":
		sc.del(w, r)
	default:
		sc.notfound(w)
	}
}
