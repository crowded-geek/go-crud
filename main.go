package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "io/ioutil"
    "strconv"
)

type Food struct { 
    Name string
    Price int
}

var m []Food

func create(w http.ResponseWriter, r *http.Request){
    var food Food
    body, er := ioutil.ReadAll(r.Body)
    er = json.Unmarshal(body, &food)
    if er != nil {
        fmt.Fprintf(w, er.Error())
        panic(er)
    }

    m = append(m, Food{Name: food.Name, Price: food.Price})
    fmt.Fprintf(w, "Created: ", m)
}

func read(w http.ResponseWriter, r *http.Request){
    for i := range m {
        var food Food = m[i]
        var k string  = strconv.Itoa(i)

        fmt.Fprintf(w, "ID: "+k+", Name: %s", string(food.Name)+", ")
        fmt.Fprintf(w, "Price: %d", food.Price)
        fmt.Fprintf(w, "\n")
    }
}

func update(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    i, err := strconv.Atoi(vars["id"])
    if err != nil {
        panic(err)
    }
    if(i>len(m)-1){
        fmt.Fprintf(w, "ID Doesn't exist")
    }else{
        var food Food
        body, er := ioutil.ReadAll(r.Body)
        er = json.Unmarshal(body, &food)
        if er != nil {
            fmt.Fprintf(w, er.Error())
            panic(er)
        }
        m[i] = Food{Name: food.Name, Price: food.Price}
        fmt.Fprintf(w, "Updated: ", m)
    }
}

func get(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    i, err := strconv.Atoi(vars["id"])
    if err != nil {
        panic(err)
    }
    if(i>len(m)-1){
        fmt.Fprintf(w, "ID Doesn't exist")
    }else{
        var food Food = m[i]
        fmt.Fprintf(w, "Name: %s", string(food.Name)+"\n")
        fmt.Fprintf(w, "Price: %d", food.Price)
    }
}

func delete(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    i, err := strconv.Atoi(vars["id"])
    if err != nil {
        panic(err)
    }
    if(i>len(m)-1){
        fmt.Fprintf(w, "ID Doesn't exist")
    }else{
        m = remove(m, i)
        fmt.Fprintf(w, "Deleted!")
    }
}

func remove(s []Food, i int) []Food {
    s[len(s)-1], s[i] = s[i], s[len(s)-1]
    return s[:len(s)-1]
}

func home(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Working")
}


func handleRequests() {
    m = make([]Food, 0)
    r := mux.NewRouter()
    r.HandleFunc("/", home).Methods("GET")
    r.HandleFunc("/api", read).Methods("GET")
    r.HandleFunc("/api/{id}", get).Methods("GET")
    r.HandleFunc("/api/", create).Methods("POST")
    r.HandleFunc("/api/{id}", update).Methods("PUT")
    r.HandleFunc("/api/{id}", delete).Methods("DELETE")
    http.ListenAndServe(":4000", r)
}

func main() {
    handleRequests()
}