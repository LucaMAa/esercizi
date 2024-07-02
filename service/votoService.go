package service

import (
	"bufio"
	"esercizi/model"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func VotoService() {
	registro := CreaRegistro()
	epsilon := 0.5

	file, err := os.Open("voti.txt")
	if err != nil {
		fmt.Println("Errore nell'apertura del file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 3 {
			fmt.Println("Errore nel formato della riga:", line)
			continue
		}

		matricola, _ := strconv.Atoi(fields[0])
		voto, _ := strconv.ParseFloat(fields[1], 64)

		cfu, _ := strconv.Atoi(fields[2])

		registro = AggiungiVoto(registro, matricola, int(voto), cfu)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Errore durante la lettura del file:", err)
	}

	coppie := TrovaCoppieSimili(registro, epsilon)

	for _, coppia := range coppie {
		fmt.Println(coppia[0], coppia[1])
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

	sum := 0.0
	sumCFU := 0
	for _, voto := range voti {
		sum += float64(voto.Voto) * float64(voto.CFU)
		sumCFU += voto.CFU
	}

	return sum / float64(sumCFU)
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

			if math.Abs(media1-media2) <= epsilon {
				coppia := []int{matricola1, matricola2}
				coppie = append(coppie, coppia)
			}
		}
	}

	return coppie
}
