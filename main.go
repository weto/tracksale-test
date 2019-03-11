package main

import (
	"fmt"
	"time"
	"net/http"
)

/*CheckNetPromoterScore é o responsável por verificar se o cliente está ou não para
receber uma nova consulta*/
func CheckNetPromoterScore(date string) (check bool){
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}
	newDate := time.Now()
	days := newDate.Sub(t).Hours() / 24
	if days > 90 {
		check = true
	}
	return
}

func sendNetPromoterScore(key string, w http.ResponseWriter) {
	SetNetPromoterScore(key)
	fmt.Fprintln(w, "Liberado para o envio de avaliação!")
}

func handleApp(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		fmt.Fprintln(w, "URL INVÁLIDA!")
		return
	}
	key := keys[0]
	netPromoterScore, err := GetNetPromoterScore(key)

	if err != nil {
		sendNetPromoterScore(key, w)
		return
	}

	check := CheckNetPromoterScore(netPromoterScore)
	if check {
		sendNetPromoterScore(key, w)
	} else {
		fmt.Fprintln(w, "Não está liberado para o envio de avaliação!")
	}
}

func main() {
	http.HandleFunc("/", handleApp)
	fmt.Println("Aplicação rodando na URL http://localhost:6060/")
	http.ListenAndServe(":8080", nil)
}
