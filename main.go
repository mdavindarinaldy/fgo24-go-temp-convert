package main

import (
	"fgo24-go-temp-convert/models"
	"fgo24-go-temp-convert/utils"
	"fmt"
	"os"
)

func main() {
	for {
		utils.Menu()
		option := utils.Input("Pilih Menu 1-5 : ")
		if option < 1 || option > 5 {
			fmt.Println("Input tidak valid!")
		} else if option == 5 {
			os.Exit(0)
		} else if option == 1 {
			utils.MenuCToF(&models.HistoryData)
		} else if option == 2 {
			utils.MenuCToK(&models.HistoryData)
		} else if option == 3 {
			utils.MenuCToR(&models.HistoryData)
		} else if option == 4 {
			utils.MenuHistory(&models.HistoryData)
		}
	}
}
