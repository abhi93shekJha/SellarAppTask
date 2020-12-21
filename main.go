package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
    "bytes"
    "SELLARAPP/db"
    "SELLARAPP/model"
    "SELLARAPP/controller"
)

var title, image_url, description, price, total_reviews string
var images[] string

type Url struct {
    URL string `json:"url"`
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
	
    myRouter.HandleFunc("/scrap_url", scrap_url).Methods("POST")
    myRouter.HandleFunc("/save_data", saveData).Methods("POST")
   
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func scrap_url(w http.ResponseWriter, r *http.Request) {
    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var new_url Url 
    json.Unmarshal(reqBody, &new_url)
 
    fmt.Println("Address1: ", new_url.URL) 
	fmt.Fprintf(w, "URL being scrapped: " + new_url.URL)
    
    // fill data after scrapping
    var data = controller.Get_data(new_url.URL)

    // runs a new post to save data
    postToSaveData(data)
    resJson, _ := json.Marshal(data)
    fmt.Fprintf(w, "\nJson scraped and saved to DB: \n" + string(resJson))
}


func postToSaveData(data model.Scrapped_data){
    var jsonData []byte
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Println(err)
    }
    fmt.Println(string(jsonData))

    url := "http://localhost:10000/save_data"
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    // body, _ := ioutil.ReadAll(resp.Body)
    // fmt.Println("response Body:", string(body))
}


func saveData(w http.ResponseWriter, r *http.Request) {
      
    reqBody, _ := ioutil.ReadAll(r.Body)

    var scrapped_data model.Scrapped_data 
    json.Unmarshal(reqBody, &scrapped_data)
    db.Insert(scrapped_data)

    // updating the response
    json.NewEncoder(w).Encode(scrapped_data)
}


func main() {
    handleRequests()
}