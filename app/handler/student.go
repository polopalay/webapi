package handler

import (
	"encoding/json"
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
	if id == 0 {
		sc.getMany(w, r)
	} else {
		sc.getone(w, r, id)
	}
}

func (sc *StudentController) getMany(w http.ResponseWriter, r *http.Request) {
	var d = dao.StudentDAO{Name: sc.DbName}
	students := d.GetAll()
	json, _ := json.Marshal(students)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))
}

func (sc *StudentController) getone(w http.ResponseWriter, r *http.Request, id int) {
	var d = dao.StudentDAO{Name: sc.DbName}
	student := d.Get(id)
	json, _ := json.Marshal(student)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))

}

func (sc *StudentController) post(w http.ResponseWriter, r *http.Request) {
	student := entity.Student{}
	json.NewDecoder(r.Body).Decode(&student)
	var d = dao.StudentDAO{Name: sc.DbName}
	err := d.Add(student)
	mss := "Success"
	if err != nil {
		mss = err.Error()
	}
	w.Write([]byte((mss)))
}

func (sc *StudentController) del(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.Join(r.URL.Query()["id"], ""))
	var d = dao.StudentDAO{Name: sc.DbName}
	err := d.Del(id)
	mss := "Success"
	if err != nil {
		mss = err.Error()
	}
	w.Write([]byte((mss)))
}

func (sc *StudentController) notfound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
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
