package models

type Data struct {
	TempC       float64
	TypeConvert string
	Result      float64
}

type History []Data

var HistoryData History

func AddHistory(history *History, data Data) {
	*history = append(*history, data)
}
