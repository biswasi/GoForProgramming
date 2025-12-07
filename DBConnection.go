package main
import (
"database/sql"
"fmt"
 "log"
 "time"

 _ "github.com/go-sql-driver/mysql"
 )


func main(){
db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/root?parseTime=true")
if err != nil {
        log.Fatal(err)
    }
if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

 { // Create a new table
        query := `
            CREATE TABLE students (
                id INT AUTO_INCREMENT,
                studentname TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

        if _, err := db.Exec(query); err != nil {
            log.Fatal(err)
        }

   { 
        studentname := "showandtell"
        password := "donttalkmuch"
        createdAt := time.Now()

        result, err := db.Exec(`INSERT INTO students (studentname, password, created_at) VALUES (?, ?, ?)`, studentname, password, createdAt)
        if err != nil {
            log.Fatal(err)
        }
        id, err := result.LastInsertId()
        fmt.Println(id)
    }
    }  
{ // Query a single student
        var (
            id        int
            studentname  string
            password  string
            createdAt time.Time
        )

        query := "SELECT id, studentname, password, created_at FROM students WHERE id = ?"
        if err := db.QueryRow(query, 1).Scan(&id, &studentname, &password, &createdAt); err != nil {
            log.Fatal(err)
        }

        fmt.Println(id, studentname, password, createdAt)
    }

  
    { 
        type student struct {
            id        int
            studentname  string
            password  string
            createdAt time.Time
        }

        rows, err := db.Query(`SELECT id, studentname, password, created_at FROM students`)
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        var students []student
        for rows.Next() {
            var s student

            err := rows.Scan(&s.id, &s.studentname, &u.password, &u.createdAt)
            if err != nil {
                log.Fatal(err)
            }
            students = append(students, u)
        }
        if err := rows.Err(); err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%#v", students)
    }

  }
