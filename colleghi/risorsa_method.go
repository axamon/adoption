package colleghi

import "time"

// AggiornaLastContact aggiorna il contatto.
func (r *Risorsa) AggiornaLastContact() {
	r.LastContact = time.Now()
	return
}
