package main

import (
	"encoding/json"
	"fmt"
	"hello-world-kubernetes/common"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	assetDir    string
	templateDir string
)

func main() {
	assetDir = os.Getenv("ASSET_DIR")
	templateDir = os.Getenv("TEMPLATE_DIR")

	http.HandleFunc("/", renderIndex)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(assetDir))))
	log.Println("Start serving on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func renderIndex(w http.ResponseWriter, _ *http.Request) {
	data := make(map[string]interface{})

	data["self"] = common.GetData()

	backendUrl := os.Getenv("BACKEND_URL")
	if backendUrl != "" {
		data["backend"] = loadBackendInfo(backendUrl)
	}

	tmpl, err := template.ParseGlob(templateDir + "/default/*")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	noAdditionalTemplates, err := IsDirEmpty(templateDir + "/additional")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if !noAdditionalTemplates {
		tmpl, err = tmpl.ParseGlob(templateDir + "/additional/*")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	err = tmpl.ExecuteTemplate(w, "index.html.tmpl", data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func loadBackendInfo(backendUrl string) map[string]string {

	resp, err := http.Get("http://" + backendUrl + "/system")
	if err != nil {
		log.Println(err)
		return nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	var backendInfo map[string]string
	err = json.Unmarshal(body, &backendInfo)
	if err != nil {
		log.Println(err)
		return nil
	}
	return backendInfo
}

func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
