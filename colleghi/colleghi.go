// Package colleghi facilita la gestione delle comunicazioni aziendali.
package colleghi

import (
	"log"
	"os"
	"time"
)

// StoreFile è il file dove salvare gli aggiornamenti.
var StoreFile = "storedRisorse.gob"

// TempoMinimo è il tempo che deve passare prima di
// poter essere ricontattati.
var TempoMinimo = 90 * 24 * time.Hour //tre mesi

func init() {
	_, err := os.Stat(StoreFile)
	if os.IsNotExist(err) {
		_, err := os.Create(StoreFile)
		if err != nil {
			log.Println(err)
		}
	}
}

// Contattabili restituisce una percentuale delle risorse contattabili
// per ogni settore.
func Contattabili(r []*Risorsa, percentuale int) []*Risorsa {
	var settori = make(map[string]int)
	var contattabili []*Risorsa
	for _, v := range r {
		settori[v.Settore]++
	}
	for settore := range settori {

		var num int
		for _, v := range r {
			if v.Settore == settore && time.Now().After(v.LastContact.Add(TempoMinimo)) {
				v.AggiornaLastContact()
				contattabili = append(contattabili, v)
				num++
				if num >= settori[v.Settore]/percentuale {
					break
				}
			}
		}
	}
	return contattabili
}
