package main

type Client struct {
    ID int
    FirstName string
    MiddleName string
    LastName string
}

type Service struct {
    Name string
    GroupName string
}

type ServiceGroup struct {
    GroupName string
}

type timeSheet struct {
    Services []Service
    ServiceGroups []ServiceGroup
}
