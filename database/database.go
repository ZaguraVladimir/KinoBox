// DataBase project DataBase.go
package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(pathToDB string, logger *log.Logger) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", pathToDB)
	if err != nil {
		logger.Println(err)
		return nil, err
	}

	if _, err := os.Stat(pathToDB); os.IsNotExist(err) {
		logger.Printf("Файл %q не найден", pathToDB)
		migration(db, logger)
	}

	if err := db.Ping(); err != nil {
		logger.Fatalln(err)
	}
	return db, err
}

func InitModels(db *sql.DB, logger *log.Logger) {
	createCountries(db, logger)
}

func createCountries(db *sql.DB, logger *log.Logger) {
	// запрос
	rows, err := db.Query("SELECT * FROM countries_view")
	if err != nil {
		logger.Fatalln(err)
	}

	for rows.Next() {
		var name, wiki, flag string
		err = rows.Scan(&name, &wiki, &flag)
		Countries[name] = country{wiki, flag}
	}
}

//	// вставка
//	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
//	checkErr(err)

//	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
//	checkErr(err)

//	id, err := res.LastInsertId()
//	checkErr(err)

//	fmt.Println(id)
//	// обновление
//	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
//	checkErr(err)

//	res, err = stmt.Exec("astaxieupdate", id)
//	checkErr(err)

//	affect, err := res.RowsAffected()
//	checkErr(err)

//	fmt.Println(affect)

//	// запрос
//	rows, err := db.Query("SELECT * FROM countries_view WHERE country = ?", "Чили")
//	if err != nil {
//		logger.Fatalln(err)
//	}

//	for rows.Next() {
//		var country, wiki, flag string
//		err = rows.Scan(&country, &wiki, &flag)
//		log.Println(country, wiki, flag)
//	}

//	// удаление
//	stmt, err = db.Prepare("delete from userinfo where uid=?")
//	checkErr(err)

//	res, err = stmt.Exec(id)
//	checkErr(err)

//	affect, err = res.RowsAffected()
//	checkErr(err)

//	fmt.Println(affect)

//	db.Close()
