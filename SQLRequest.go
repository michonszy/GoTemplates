package main
import (
	//"log"
	"fmt"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	//"reflect"
	"os"
)
func main(){

	db, err := sql.Open("mssql", "server="+"FQDN\\SQLEXPRESS"+";user id="+"USER"+";password="+"PASSEORD")
	if err != nil {
		fmt.Println("From Open() attempt: " + err.Error())
	   }
	defer db.Close()
	
	
	// Run the query
	rows, err := db.Query("use main;SELECT * FROM  table;")
	if err != nil {
	 fmt.Println(err)
	}
	
	var result [][]string

	cols,err := rows.Columns()
	if err != nil {
		fmt.Println(err)
   	}
	pointers := make([]interface{}, len(cols))
	container := make([]string, len(cols))

	plik2, err := os.OpenFile("CHECKSQL.txt", os.O_APPEND|os.O_WRONLY, 0)


	for i, _ := range pointers {	
		pointers[i] = &container[i]
	}
	for rows.Next() {
		rows.Scan(pointers...)
		fmt.Println(container)
		fmt.Println("==========================================")
		result = append(result, container)
		for _, word := range container {
			plik2.WriteString(word)
			plik2.WriteString(" | ")
		}
		plik2.WriteString("\n")
	}
		
}