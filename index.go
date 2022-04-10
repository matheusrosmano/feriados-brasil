package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	fmt.Println("Feriados no brasil")

	fmt.Println("Digite o ano desejado: ")

	var ano int
	fmt.Scan(&ano)

	data := getDataNumeroDourado(ano)
	pascoa := getDayOfPascoa(data)
	carnaval := getDayOfCarnaval(pascoa)
	fmt.Println("O carnaval é:", carnaval)
}

func getDataNumeroDourado(ano int) time.Time{
	result := math.Mod(float64(ano), 19)

	var timeDate time.Time

	switch int(result) + 1 {
		case 1:
			timeDate = time.Date(ano, 4, 14, 5, 0, 0, 0, time.UTC)
		case 2:
			timeDate = time.Date(ano, 4, 3, 5, 0, 0, 0, time.UTC)
		case 3:
			timeDate = time.Date(ano, 3, 23, 5, 0, 0, 0, time.UTC)
		case 4:
			timeDate = time.Date(ano, 4, 11, 5, 0, 0, 0, time.UTC)
		case 5:
			timeDate = time.Date(ano, 3, 31, 5, 0, 0, 0, time.UTC)
		case 6:
			timeDate = time.Date(ano, 4, 18, 5, 0, 0, 0, time.UTC)
		case 7:
			timeDate = time.Date(ano, 4, 8, 5, 0, 0, 0, time.UTC)
		case 8:
			timeDate = time.Date(ano, 3, 28, 5, 0, 0, 0, time.UTC)
		case 9:
			timeDate = time.Date(ano, 4, 16, 5, 0, 0, 0, time.UTC)
		case 10:
			timeDate = time.Date(ano, 4, 5, 5, 0, 0, 0, time.UTC)
		case 11:
			timeDate = time.Date(ano, 3, 25, 5, 0, 0, 0, time.UTC)
		case 12:
			timeDate = time.Date(ano, 4, 13, 5, 0, 0, 0, time.UTC)
		case 13:
			timeDate = time.Date(ano, 4, 2, 5, 0, 0, 0, time.UTC)
		case 14:
			timeDate = time.Date(ano, 3, 22, 5, 0, 0, 0, time.UTC)
		case 15:
			timeDate = time.Date(ano, 4, 10, 5, 0, 0, 0, time.UTC)
		case 16:
			timeDate = time.Date(ano, 3, 30, 5, 0, 0, 0, time.UTC)
		case 17:
			timeDate = time.Date(ano, 4, 17, 5, 0, 0, 0, time.UTC)
		case 18:
			timeDate = time.Date(ano, 4, 7, 5, 0, 0, 0, time.UTC)
		case 19:
			timeDate = time.Date(ano, 3, 27, 5, 0, 0, 0, time.UTC)
		default:
			fmt.Println("Invalido")
			os.Exit(-1)
	}
	return timeDate
}

func getDayOfPascoa(data time.Time) time.Time {
	dia := data.Weekday()
	var novaData time.Time

	switch dia {
		case 0:
			novaData = data.AddDate(0, 0, 7)
		case 1:
			novaData = data.AddDate(0, 0, 6)
		case 2:
			novaData = data.AddDate(0, 0, 5)
		case 3:
			novaData = data.AddDate(0, 0, 4)
		case 4:
			novaData = data.AddDate(0, 0, 3)
		case 5:
			novaData = data.AddDate(0, 0, 2)
		case 6:
			novaData = data.AddDate(0, 0, 1)
	}
	fmt.Println("A pascoa vai ser:", novaData)
	return novaData
}

func getDayOfCarnaval(data time.Time) time.Time {
	novaData := data.AddDate(0, 0, -47)
	return novaData
}