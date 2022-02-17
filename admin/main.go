package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Stats struct {
	MeanTimes map[string]int `json:"meanTimes"`
}

type dbResponse struct {
	ID       int       `json:"-"`
	BackType string    `json:"backType"`
	ExecTime int       `json:"execTime"`
	CallDate time.Time `json:"callDate"`
}

type resp struct {
	Stats     Stats        `json:"stats"`
	Responses []dbResponse `json:"responses"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)

	html, err := os.ReadFile("index.html")
	if err != nil {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"code": 500, "err":true, "msg":"error opening the file: %s"}`, err.Error())))
		return
	}

	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(html)
}

func GetAllResponsesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	log.Println(r.Method, r.RequestURI)

	stats, resps, err := GetAllResponses()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"code": 500, "err":true, "msg":"error getting the data: %s"}`, err.Error())))
		return
	}
	// fmt.Println(stats)
	resp := resp{Stats: stats, Responses: resps}
	body, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"code": 500, "err":true, "msg":"error getting the data: %s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/getAll", GetAllResponsesHandler)

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)

}
