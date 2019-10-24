package todo

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type DatabaseInformation struct {
	dbdriver string
	password string
	username string
	database string
}

var database = DatabaseInformation{
	dbdriver: "mysql",
	password: "passowrd",
	username: "root",
	database: "mysql123",
}

const (
	insert      string = "insert into todos (todo ,completed) values (?,?)"
	selectquery string = "select todo,id from todos "
	remotequery string = "delete from todos where todo=?"
)

var datasourcename = database.username + ":" + database.password + "@/" + database.database

type todoname struct {
	Todo string
}
type Bınd struct {
	Todo string
}

func AddTodo(c *gin.Context) {
	var Param Bınd

	db, err := connectToDb()
	defer db.Close()
	if err == nil {
		prepare, err := db.Prepare(insert)
		if err == nil {
			c.BindJSON(&Param)

			_, err := prepare.Exec(Param.Todo, false)
			if err == nil {
				log.Print("Başarıyla kayıt yapıldı ")
				prepare.Close()

				c.JSON(201, nil)
			}
		}

	} else {
		log.Fatal("applicaon is down ")
	}
}

type Bind2 struct {
	Todo []string
}

func GetTodos(c *gin.Context) {
	var todo string
	var id int
	var m Bind2

	db, err := connectToDb()
	if err == nil {
		results, _ := db.Query(selectquery)

		for results.Next() {

			results.Scan(&todo, &id)
			m.Todo = append(m.Todo, todo)

		}

	} else {
		log.Fatal("Bir sorunla karşılıldı" + err.Error())
	}

	c.JSON(200, m)
}

func DeleteTodo(c *gin.Context) {

	var Param Bınd
	db, err := connectToDb()
	defer db.Close()
	if err == nil {
		prepare, err := db.Prepare(remotequery)
		if err == nil {
			c.BindJSON(&Param)
			fmt.Print("param")
			_, err := prepare.Exec(Param.Todo)
			if err == nil {
				log.Print("Silinidi")

				c.JSON(200, nil)
				fmt.Print("tirrej")
			} else {
				fmt.Print("nasıl gırıyon buraya")
			}
		}

	} else {
		log.Fatal("applicaon is down ")
	}

}
func UpdateTodo(c *gin.Context) {

}

func connectToDb() (db *sql.DB, err error) {

	db, err = sql.Open(database.dbdriver, datasourcename)
	return
}
