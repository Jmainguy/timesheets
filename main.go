package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/addClient", addClientPage)
    http.HandleFunc("/listClients", listClientsPage)
    http.HandleFunc("/addService", addServicesPage)
    http.HandleFunc("/listServices", listServicesPage)
    http.HandleFunc("/", timeSheetPage)
    http.ListenAndServe(":8000", nil)
}
