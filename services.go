package main

import (
    "fmt"
    "net/http"
    "html/template"
)

func addClientPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("web/addClient.gtpl")
        t.Execute(w, "")
    } else if r.Method == "POST" {
        r.ParseForm()
        firstname := r.PostFormValue("firstname")
        lastname := r.PostFormValue("lastname")
        middlename := r.PostFormValue("middlename")
        addClientSQL(firstname, lastname, middlename)
        t, _ := template.ParseFiles("web/addClient.gtpl")
        name := fmt.Sprintf("%s %s %s", firstname, middlename, lastname)
        t.Execute(w, name)
    }
}

func listClientsPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        ClientList := listClients()
        t, _ := template.ParseFiles("web/listClients.gtpl")
        t.Execute(w, ClientList)
    }
}

func addServicesPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("web/addService.gtpl")
        t.Execute(w, "")
    } else if r.Method == "POST" {
        r.ParseForm()
        name := r.PostFormValue("name")
        group := r.PostFormValue("group")
        addServiceSQL(name, group)
        t, _ := template.ParseFiles("web/addService.gtpl")
        msg := fmt.Sprintf("%s has been added to the %s group", name, group)
        t.Execute(w, msg)
    }
}

func listServicesPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        ServiceList := listServices()
        t, _ := template.ParseFiles("web/listServices.gtpl")
        t.Execute(w, ServiceList)
    }
}

