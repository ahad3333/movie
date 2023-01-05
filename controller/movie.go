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

	if r.Method == "GET" {
	id := r.URL.Query().Get("id")
	if id == "" {
		c.GetByIdMovie(w, r)
		
	}else{
		c.GetListMovie(w, r)
	}
	}

	if r.Method == "PUT" {
		c.UpdateMovie(w, r)
	}

	if r.Method == "DELETE" {
		c.DeleteMovie(w, r)
	}
}

func (c *Controller) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error whiling movie post method: ", err)
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

	movie, err := storage.GetByIdMovie(c.db, id)
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

func (c *Controller) GetListMovie(w http.ResponseWriter, r *http.Request) {

	movie, err := storage.GetListMovie(c.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error whiling storage get List movie: ", err)
		return
	}

	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error whiling encode movie: ", err)
		return
	}
}

func (c *Controller) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error whiling movie put method: ", err)
		return
	}

	err = json.Unmarshal(body, &movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error whiling movie json unmarshal: ", err)
		return
	}

	err = storage.UpdateMovie(c.db, movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error whiling storage update movie: ", err)
		return
	}

	err = json.NewEncoder(w).Encode("updated Movie") 
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error updated movie: ", err)
		return
	}

}

func (c *Controller) DeleteMovie(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	err := storage.DeleteMovie(c.db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error get delete  id movie: ", err)
		return
	}

	err = json.NewEncoder(w).Encode("delete Movie")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error delete movie: ", err)
		return
	}

}