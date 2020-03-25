// DataBase project DataBase.go
package database

type country struct{ Wiki, Flag string }

var Countries = make(map[string]country)
