package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tp2/comandos"
)

const (
	_AGREGAR_ARCHIVO   = "agregar_archivo"
	_VER_VISITANTES    = "ver_visitantes"
	_VER_MAS_VISITADOS = "ver_mas_visitados"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	logs := comandos.CrearDetectorDeLogs()

	for scanner.Scan() {
		expresion := scanner.Text()
		elementos := strings.Fields(expresion)
		comando := elementos[0]

		switch comando {
		case _AGREGAR_ARCHIVO:
			if len(elementos) != 2 {
				fmt.Fprintf(os.Stderr, "Error en comando %s\n", _AGREGAR_ARCHIVO)
				return
			}
			dos, err := logs.AgregarArchivo(elementos[1])

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error en comando %s\n", _AGREGAR_ARCHIVO)
				break
			}

			for _, ip := range dos {
				fmt.Fprintf(os.Stdout, "DoS: %s\n", ip)
			}
			fmt.Fprintln(os.Stdout, "OK")
		case _VER_VISITANTES:
			if len(elementos) != 3 {
				fmt.Fprintf(os.Stderr, "Error en comando %s\n", _VER_VISITANTES)
				break
			}
			visitantes := logs.VerVisitantes(elementos[1], elementos[2])
			fmt.Fprintf(os.Stdout, "%s\n", "Visitantes:")
			for _, visitante := range visitantes {
				fmt.Fprintf(os.Stdout, "\t%s\n", visitante)
			}
			fmt.Fprintln(os.Stdout, "OK")
		case _VER_MAS_VISITADOS:
			if len(elementos) != 2 {
				fmt.Fprintf(os.Stderr, "Error en comando %s\n", _VER_MAS_VISITADOS)
				break
			}
			n, _ := strconv.Atoi(elementos[1])
			mas_visitados := logs.VerMasVisitados(n)
			fmt.Fprintf(os.Stdout, "%s\n", "Sitios más visitados:")
			for _, visitado := range mas_visitados {
				sitio, visitas := visitado.VerPar()
				fmt.Fprintf(os.Stdout, "\t%s - %d\n", sitio, visitas)
			}
			fmt.Fprintln(os.Stdout, "OK")
		default:
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stdout, "ERROR", err)
		os.Exit(1)
	}
}
