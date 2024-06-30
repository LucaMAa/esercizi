package service

import (
	"esercizi/model"
	"fmt"
	"sort"
)

func VotoService() {
	registro := CreaRegistro()

	var matricola, voto, cfu int
	for {
		_, err := fmt.Scan(&matricola, &voto, &cfu)
		if err != nil {
			break
		}
		registro = AggiungiVoto(registro, matricola, voto, cfu)
	}

	coppie := TrovaCoppieSimili(registro, 0.5)

	for _, coppia := range coppie {
		fmt.Println(coppia[0], coppia[1])
	}
	fmt.Println("Coppie con medie simili trovate:", coppie)
	for matricola := range registro {
		ultimoVoto := Ultimo(registro, matricola)
		fmt.Printf("Ultimo voto per matricola %d: %d CFU %d\n", matricola, ultimoVoto.Voto, ultimoVoto.CFU)
	}
}

// CreaRegistro crea un registro vuoto
func CreaRegistro() model.Registro {
	return make(model.Registro)
}

// CreaVoto crea un oggetto Voto con voto e CFU specificati
func CreaVoto(voto, cfu int) model.Voto {
	if voto < 18 || voto > 30 || cfu < 2 || cfu > 12 {
		return model.Voto{}
	}
	return model.Voto{Voto: voto, CFU: cfu}
}

// AggiungiVoto aggiunge un voto al registro per la matricola specificata
func AggiungiVoto(r model.Registro, matricola, voto, cfu int) model.Registro {
	if v := CreaVoto(voto, cfu); v.Voto != 0 {
		if _, ok := r[matricola]; !ok {
			r[matricola] = []model.Voto{}
		}
		r[matricola] = append(r[matricola], v)
	}
	return r
}

// Ultimo restituisce l'ultimo voto registrato per la matricola specificata
func Ultimo(r model.Registro, matricola int) model.Voto {
	if voti, ok := r[matricola]; ok && len(voti) > 0 {
		return voti[len(voti)-1]
	}
	return model.Voto{}
}

// MediaPesata calcola la media pesata dei voti per la matricola specificata
func MediaPesata(r model.Registro, matricola int) float64 {
	voti, ok := r[matricola]
	if !ok || len(voti) == 0 {
		return 0.0
	}

	sum := 0
	sumCFU := 0
	for _, voto := range voti {
		sum += voto.Voto * voto.CFU
		sumCFU += voto.CFU
	}

	return float64(sum) / float64(sumCFU)
}

// TrovaCoppieSimili trova tutte le coppie di matricole con media pesata simile
func TrovaCoppieSimili(r model.Registro, epsilon float64) [][]int {
	var coppie [][]int
	matricole := make([]int, 0, len(r))

	for matricola := range r {
		matricole = append(matricole, matricola)
	}
	sort.Ints(matricole)

	for i := 0; i < len(matricole); i++ {
		for j := i + 1; j < len(matricole); j++ {
			matricola1 := matricole[i]
			matricola2 := matricole[j]
			media1 := MediaPesata(r, matricola1)
			media2 := MediaPesata(r, matricola2)

			if media1-epsilon <= media2 && media2 <= media1+epsilon {
				coppia := []int{matricola1, matricola2}
				coppie = append(coppie, coppia)
			}
		}
	}

	return coppie
}
