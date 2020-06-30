package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", mux("/static")))
}

const jsonParam = "json"

func mux(dir string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodDelete {
			deleteFileIfExists(w, r, dir)
			return
		}
		if _, ok := r.URL.Query()[jsonParam]; ok {
			listFilesAsJson(w, dir)
			return
		}
		http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
	})
	return mux
}

func listFilesAsJson(w http.ResponseWriter, dir string) {
	files, err := ioutil.ReadDir(dir)
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().After(files[j].ModTime())
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ret := []string{}
	for _, f := range files {
		ret = append(ret, f.Name())
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ret)
}

func deleteFileIfExists(w http.ResponseWriter, r *http.Request, dir string) {
	fileName := strings.TrimPrefix(r.URL.Path, "/")
	filePath := filepath.Join(dir, fileName)
	_, err := os.Stat(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unknown file %s", fileName), http.StatusNotFound)
		return
	}
	err = os.Remove(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete file %s: %v", fileName, err), http.StatusInternalServerError)
		return
	}
}
