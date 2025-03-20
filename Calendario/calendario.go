package main

import (
	"fmt"
	"time"
)

func main() {
	//Fechas
	fechaInicio := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	fechaFin := time.Date(2025, 8, 30, 0, 0, 0, 0, time.UTC)

	//Nombre día de la semana
	for fechaInicio.Before(fechaFin) || fechaInicio.Equal(fechaFin) {
		var dayName string
		weekDay := fechaInicio.Weekday()
		switch weekDay {
		case time.Sunday:
			dayName = "Domingo"
		case time.Monday:
			dayName = "Lunes"
		case time.Tuesday:
			dayName = "Martes"
		case time.Wednesday:
			dayName = "Miércoles"
		case time.Thursday:
			dayName = "Jueves"
		case time.Friday:
			dayName = "Viernes"
		case time.Saturday:
			dayName = "Sábado"
		}
		//Imprimir los días
		fmt.Println(fechaInicio.Format("2006-01-02"), dayName)

		//Incrementar un día
		fechaInicio = fechaInicio.AddDate(0, 0, 1)
	}

}
