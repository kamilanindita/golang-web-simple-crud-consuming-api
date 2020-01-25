package main

import (
    "log"
    "net/http"
    "text/template"
    "io/ioutil"
    "github.com/parnurzeal/gorequest"
    "encoding/json"

)

type Buku struct {
    Id    int `json:"Id"`
    Penulis  string `json:"Penulis"`
    Judul string `json:"Judul"`
    Kota string `json:"Kota"`
    Penerbit string `json:"Penerbit"`
    Tahun int `json:"Tahun"`
}

type Response struct{
    Status bool `json:"Status"`
    Message string `json:"Message"`
    Data []Buku `json:"Data"`
}



func HandlerIndex(w http.ResponseWriter, r *http.Request) {
    var tmp = template.Must(template.ParseFiles(
        "views/Header.html",
        "views/Menu.html",
        "views/Index.html",
        "views/Footer.html",
    ))
    data:=""
    var error = tmp.ExecuteTemplate(w,"Index",data)
    if error != nil {
        http.Error(w, error.Error(), http.StatusInternalServerError)
    }
}


func HandlerBuku(w http.ResponseWriter, r *http.Request) {
    request := gorequest.New()
    resp, _, errs := request.Get("http://127.0.0.1:8080/buku").End()
    if errs != nil {
    log.Fatal(errs)
    } 

    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    
    var responseObject Response
    json.Unmarshal(responseData, &responseObject)
      
    if responseObject.Status==true {
        var tmp = template.Must(template.ParseFiles(
            "views/Header.html",
            "views/Menu.html",
            "views/Buku.html",
            "views/Footer.html",
        ))

        var error = tmp.ExecuteTemplate(w,"Buku",responseObject.Data)
        if error != nil {
            http.Error(w, error.Error(), http.StatusInternalServerError)
        }
    }
}


func HandlerTamabah(w http.ResponseWriter, r *http.Request) {
    var tmp = template.Must(template.ParseFiles(
        "views/Header.html",
        "views/Menu.html",
        "views/Tambah.html",
        "views/Footer.html",
    ))
    data:=""
    var error = tmp.ExecuteTemplate(w,"Tambah",data)
    if error != nil {
        http.Error(w, error.Error(), http.StatusInternalServerError)
    }
}


func HandlerSave(w http.ResponseWriter, r *http.Request) {
    penulis := r.FormValue("penulis")
    judul := r.FormValue("judul")
    kota := r.FormValue("kota")
    penerbit := r.FormValue("penerbit")
    tahun := r.FormValue("tahun")
    
    m := map[string]string{"penulis":penulis, "judul": judul, "kota": kota, "penerbit":penerbit, "tahun": tahun}
    
    request := gorequest.New()
    resp, _, errs := request.Post("http://127.0.0.1:8080/buku").Set("Content-Type", "application/x-www-form-urlencoded").Send(m).End()
    if errs != nil {
    log.Fatal(errs)
    } 
    
    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    var responseObject Response
    json.Unmarshal(responseData, &responseObject)
    
    if responseObject.Status==true {
        http.Redirect(w, r, "/buku", 301)
    }

}


func HandlerEdit(w http.ResponseWriter, r *http.Request) {
    Id := r.URL.Query().Get("id")

    request := gorequest.New()
    resp, _, errs := request.Get("http://127.0.0.1:8080/buku/"+Id).End()
    if errs != nil {
    log.Fatal(errs)
    } 

    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    
    var responseObject Response
    json.Unmarshal(responseData, &responseObject)
      
    if responseObject.Status==true {
        var tmp = template.Must(template.ParseFiles(
            "views/Header.html",
            "views/Menu.html",
            "views/Edit.html",
            "views/Footer.html",
        ))
    
        var error = tmp.ExecuteTemplate(w,"Edit",responseObject.Data[0])
        if error != nil {
            http.Error(w, error.Error(), http.StatusInternalServerError)
        }
    }

}


func HandlerUpdate(w http.ResponseWriter, r *http.Request) {
    Id := r.URL.Query().Get("id")
    
    penulis := r.FormValue("penulis")
    judul := r.FormValue("judul")
    kota := r.FormValue("kota")
    penerbit := r.FormValue("penerbit")
    tahun := r.FormValue("tahun")
    
    m := map[string]string{"penulis":penulis, "judul": judul, "kota": kota, "penerbit":penerbit, "tahun": tahun}
    
    request := gorequest.New()
    resp, _, errs := request.Put("http://127.0.0.1:8080/buku/"+Id).Set("Content-Type", "application/x-www-form-urlencoded").Send(m).End()
    if errs != nil {
    log.Fatal(errs)
    } 

    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject Response
    json.Unmarshal(responseData, &responseObject)
    
    if responseObject.Status==true {
        http.Redirect(w, r, "/buku", 301)
    }

}


func HandlerDelete(w http.ResponseWriter, r *http.Request) {
    Id := r.URL.Query().Get("id")
    
    request := gorequest.New()
    resp, _, errs := request.Delete("http://127.0.0.1:8080/buku/"+Id).End()
    if errs != nil {
    log.Fatal(errs)
    } 

    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject Response
    json.Unmarshal(responseData, &responseObject)
      
    if responseObject.Status==true {
        http.Redirect(w, r, "/buku", 301)
    }

}



func main() {
    log.Println("Server started on: http://localhost:8000")
    http.HandleFunc("/", HandlerIndex)
    http.HandleFunc("/buku", HandlerBuku)
    http.HandleFunc("/buku/tambah", HandlerTamabah)
    http.HandleFunc("/buku/edit", HandlerEdit)
    http.HandleFunc("/buku/save", HandlerSave)
    http.HandleFunc("/buku/update", HandlerUpdate)
    http.HandleFunc("/buku/delete", HandlerDelete)
    http.ListenAndServe(":8000", nil)
}