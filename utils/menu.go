package utils

import (
	"fgo24-go-temp-convert/calculate"
	"fgo24-go-temp-convert/models"
	"fmt"
	"time"
)

func Menu() {
	fmt.Printf("\nMenu : \n")
	fmt.Println("1. Konversi Celcius ke Fahrenheit")
	fmt.Println("2. Konversi Celcius ke Kelvin")
	fmt.Println("3. Konversi Celcius ke Reamur")
	fmt.Println("4. Lihat Histori")
	fmt.Printf("5. Exit Program\n\n")
}

func MenuCToF(history *models.History) {
	fmt.Printf("\n\n---Konversi Celcius ke Fahrenheit---\n")
	celcius := InputTemp("Masukkan suhu dalam Celcius : ")
	fahrenheit := calculate.CalculateCtoF(celcius)
	fmt.Printf("Hasil konversi suhu dalam Fahrenheit : %f\n", fahrenheit)
	models.AddHistory(history, models.Data{
		TempC:       celcius,
		TypeConvert: "Fahrenheit",
		Result:      fahrenheit,
	})
	time.Sleep(2 * time.Second)
}

func MenuCToK(history *models.History) {
	fmt.Printf("\n\n---Konversi Celcius ke Kelvin---\n")
	celcius := InputTemp("Masukkan suhu dalam Celcius : ")
	kelvin := calculate.CalculateCtoK(celcius)
	fmt.Printf("Hasil konversi suhu dalam Kelvin : %f\n", kelvin)
	models.AddHistory(history, models.Data{
		TempC:       celcius,
		TypeConvert: "Kelvin",
		Result:      kelvin,
	})
	time.Sleep(2 * time.Second)
}

func MenuCToR(history *models.History) {
	fmt.Printf("\n\n---Konversi Celcius ke Reamur---\n")
	celcius := InputTemp("Masukkan suhu dalam Celcius : ")
	reamur := calculate.CalculateCtoR(celcius)
	fmt.Printf("Hasil konversi suhu dalam Reamur : %f\n", reamur)
	models.AddHistory(history, models.Data{
		TempC:       celcius,
		TypeConvert: "Reamur",
		Result:      reamur,
	})
	time.Sleep(2 * time.Second)
}

func MenuHistory(history *models.History) {
	fmt.Printf("\n\n---Histori Konversi Suhu yang telah dilakukan---\n")
	if len(*history) == 0 {
		fmt.Println("Belum ada histori!")
		time.Sleep(time.Second)
		return
	}
	for i, data := range *history {
		fmt.Printf("%d. Konversi dari Celcius (Suhu: %f) ke %s dengan hasil %f\n", i+1, data.TempC, data.TypeConvert, data.Result)
	}
	time.Sleep(2 * time.Second)
}
