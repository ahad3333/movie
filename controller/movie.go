package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"add/models"
	"add/storage"
)

func (c *Controller) Movie(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		c.CreateMovie(w, r)
	}

	// if r.Method == "GET" {
	// 	c.GetByIdMovie(w, r)
	// }
	if r.Method == "GET" {
		c.GetByIdMovie1(w, r)
	}
}

func (c *Controller) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error whiling movie post method 11: ", err)
		return
	}

	err = json.Unmarshal(body, &movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error whiling movie json unmarshal: ", err)
		return
	}

	id, err := storage.InsertMovie(c.db, movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error whiling storage create movie: ", err)
		return
	}

	movie, err = storage.GetByIdMovie(c.db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error whiling storage get by id movie: ", err)
		return
	}

	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error whiling encode movie: ", err)
		return
	}
}

func (c *Controller) GetByIdMovie(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	movie, err := storage.GetByIdMovie(c.db,id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error whiling storage get by id movie: ", err)
		return
	}

	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error whiling encode movie: ", err)
		return
	}
}
func (c *Controller) GetByIdMovie1(w http.ResponseWriter, r *http.Request) {

    // all:=r.URL.Query().Get("movie")

	movie, err := storage.GetByIdMovie1(c.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error whiling storage get by id movie: ", err)
		return
	}

	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error whiling encode movie: ", err)
		return
	}
}