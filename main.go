package main

import (
	"adoption/colleghi"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var salva = flag.Bool("s", false, "salva")

func main() {

	rand.NewSource(time.Now().Unix())

	flag.Parse()

	if *salva {
		carica()

	}
	rr, err := colleghi.CaricaDati()
	if err != nil {
		log.Println(err)
	}

	for i := 0; i <= 10; i++ {
		fmt.Println("tranch:", i+1)
		for _, e := range rr {
			fmt.Printf("%s\t%s\t%s\n", e.Matricola, e.Settore, e.LastContact.Format("2006-01-02"))
		}
	}

	// for _, v := range rr {
	// 	fmt.Println(v)
	// }

	rrr := colleghi.Contattabili(rr, 10)

	for i := 0; i <= 10; i++ {
		fmt.Println("tranch:", i+1)
		for _, e := range rrr {
			fmt.Printf("%s\t%s\t%s\n", e.Matricola, e.Settore, e.LastContact.Format("2006-01-02"))
		}
	}

	colleghi.SalvaSuFile(rr)

}

func carica() {

	type dati struct {
		max     int
		settore string
	}

	o := []dati{
		{max: 42, settore: "CT"},
		{max: 10, settore: "CT.IO"},
		{max: 22, settore: "CT.5GDT"},
	}

	var risorse []*colleghi.Risorsa

	for _, a := range o {
		for i := 0; i < a.max; i++ {
			var r = &colleghi.Risorsa{
				Matricola:   strconv.Itoa(1000 + i),
				Settore:     a.settore,
				LastContact: time.Now().Local().Add(time.Duration(-1*rand.Intn(200)) * 24 * time.Hour), //time.Now(),
			}
			risorse = append(risorse, r)
		}
	}
	colleghi.SalvaSuFile(risorse)

}
