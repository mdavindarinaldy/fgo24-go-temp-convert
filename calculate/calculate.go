package calculate

func CalculateCtoF(celcius float64) float64 {
	return (celcius * 9 / 5) + 32
}

func CalculateCtoK(celcius float64) float64 {
	return celcius + 273.15
}

func CalculateCtoR(celcius float64) float64 {
	return celcius * 4 / 5
}
