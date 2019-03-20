package api

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PhotoBase interface {
	List() ([]Photo, error)
	Add(Photo) error
	Get(string) (Photo, error)
}

type Photo struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Data   []byte `json:"-"`
	Hero   string `json:"hero"`
}

func New(photobase PhotoBase) http.Handler {
	handler := httprouter.New()

	photoHandler := PhotoHandler{photobase}

	handler.GET("/photos", photoHandler.listPhotos)
	handler.GET("/photos/:id", photoHandler.getPhoto)
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

func (p *PhotoHandler) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	photo, err := p.photobase.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	buffer := bytes.NewBuffer(photo.Data)

	//if err := jpeg.Encode(buffer, photo.Data, nil); err != nil {
	//log.Println("unable to encode image.")
	//}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
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
