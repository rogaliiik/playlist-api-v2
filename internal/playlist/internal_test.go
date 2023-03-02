package internal_test

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rogaliiik/playlist/internal/models"
	internal "github.com/rogaliiik/playlist/internal/playlist"
	"github.com/rogaliiik/playlist/internal/routes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var router *mux.Router

func TestMain(m *testing.M) {
	internal.APIServer.Store.DB.AutoMigrate(&models.Song{})
	internal.APIServer.Store.DB.AutoMigrate(&models.PlaylistDB{})
	router = mux.NewRouter()
	routes.RegisterRoutes(router)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	return writer
}

func TestGetAllSongs(t *testing.T) {
	resp := makeRequest("GET", "/song", nil)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetSongById(t *testing.T) {
	resp := makeRequest("GET", "/song/1", nil)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdateSong(t *testing.T) {
	resp := makeRequest("PUT", "/song/1", nil)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteSong(t *testing.T) {
	resp := makeRequest("DELETE", "/song/11", nil)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestPlaySong(t *testing.T) {
	resp := makeRequest("GET", "/play", nil)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestPauseSong(t *testing.T) {
	resp := makeRequest("GET", "/pause", nil)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestNextSong(t *testing.T) {
	resp := makeRequest("GET", "/next", nil)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestPrevSong(t *testing.T) {
	resp := makeRequest("GET", "/prev", nil)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestStatus(t *testing.T) {
	resp := makeRequest("GET", "/status", nil)
	assert.Equal(t, http.StatusOK, resp.Code)
}
