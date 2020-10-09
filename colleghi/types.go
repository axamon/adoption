package colleghi

import "time"

// Risorsa contiene informazioni dell'utente.
type Risorsa struct {
	Matricola   string
	Settore     string
	LastContact time.Time
}
