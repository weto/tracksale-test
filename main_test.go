package main

import "testing"

func TestCheckNetPromoterScoreSend(t *testing.T) {
	if CheckNetPromoterScore("2018-12-11T00:00:00.000Z") {
		t.Error("Esperado retorno verdadeiro", )
	}
}

func TestCheckNetPromoterScoreNoSend(t *testing.T) {
	if CheckNetPromoterScore("2019-03-11T00:00:00.000Z") {
		t.Error("Esperado retorno Falso")
	}
}
