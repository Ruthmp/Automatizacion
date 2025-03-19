package main
import (
	"encoding/csv"
	"fmt"
	"os"
)
func main(){
	file, err := os.Open("./fichero.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo", err)
		return
	}
	defer file.Close()
	reader:=csv.NewReader(file)
	records, err:= reader.ReadAll()
	if err!= nil{
		fmt.Println("Error al leer el archivo", err)
		return
	}
	for i, record:= range records{
		fmt.Println("línea", i, "es", record)
	}
		
	}
