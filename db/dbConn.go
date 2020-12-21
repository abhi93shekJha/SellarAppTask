package db

import(
	"database/sql"
    "time"
    _ "github.com/go-sql-driver/mysql"
    "SELLARAPP/model"
) 


func Insert(data model.Scrapped_data) {
    
    db := create_database()
    currentTime := time.Now()
    time := currentTime.Format("2006.01.02 15:04:05")
    var query = "INSERT INTO scrappeddata VALUES ('"+string(time)+"','"+data.Title+"','"+data.Image_url+"','"+data.Total_reviews+"','"+data.Price+"','"+data.Description+"')"
    _,err := db.Exec(query)
    if err != nil {
        panic(err)
    }
    
    //insForm.Exec(time, data.Title, data.Image_url, data.Total_reviews, data.Price, data.Description)
    defer db.Close()

}

func create_database() (db *sql.DB){
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "root"
    dbName := "mysql"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err)
    }
 
    _,err = db.Exec("CREATE DATABASE IF NOT EXISTS "+ "scrapped")
    if err != nil {
        panic(err)
    }
 
    _,err = db.Exec("USE "+ "scrapped")
    if err != nil {
        panic(err)
    }

    _,err = db.Exec("CREATE TABLE IF NOT EXISTS scrappeddata (timestamp varchar(100), title varchar(1000), images LONGTEXT, total_reviews varchar(1000), price varchar(1000), description LONGTEXT)")
    if err != nil {
        panic(err)
    }
    return db
}
