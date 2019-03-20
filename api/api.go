package api

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PhotoBase interface {
	List() ([]Photo, error)
	Add(Photo) error
}

type Photo struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Data   []byte `json:"data"`
	Hero   string `json:"hero"`
}

func New(photobase PhotoBase) http.Handler {
	handler := httprouter.New()

	photoHandler := PhotoHandler{photobase}

	handler.GET("/photos", photoHandler.listPhotos)
	handler.POST("/photos", photoHandler.addPhoto)

	return handler
}

type PhotoHandler struct {
	photobase PhotoBase
}

func (p *PhotoHandler) listPhotos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	photoList, err := p.photobase.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err = json.NewEncoder(w).Encode(&photoList)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *PhotoHandler) addPhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Adding photo")

	file, _, err := r.FormFile("photo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	hero := md5.Sum([]byte(fmt.Sprintf("%s-%s", title, author)))

	photo := Photo{
		Title:  title,
		Author: author,
		Data:   bytes,
		Hero:   fmt.Sprintf("%x", hero),
	}

	err = p.photobase.Add(photo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
}
