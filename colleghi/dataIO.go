package colleghi

import (
	"encoding/gob"
	"log"
	"os"
)

// SalvaSuFile salva i dati in StoreFile.
func SalvaSuFile(dati []*Risorsa) error {

	dataFile, err := os.OpenFile(StoreFile, 666, 666)
	if err != nil {
		log.Println(err)
	}
	defer dataFile.Close()

	// serialize the data
	dataEncoder := gob.NewEncoder(dataFile)
	err = dataEncoder.Encode(dati)
	if err != nil {
		log.Println(err)
	}

	return err
}

// CaricaDati recupera i dati da StoreFile.
func CaricaDati() ([]*Risorsa, error) {

	dataFile, err := os.Open(StoreFile)
	if err != nil {
		log.Println(err)
	}
	defer dataFile.Close()

	var data []*Risorsa
	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&data)
	if err != nil {
		log.Println(err)
	}

	return data, err
}
