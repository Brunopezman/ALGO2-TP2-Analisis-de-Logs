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
	for scanner.Scan() {

		expresion := scanner.Text()
		elementos := strings.Fields(expresion)
		comando := elementos[0]
		logs := comandos.CrearDetectorDeLogs()
		switch comando {
		case _AGREGAR_ARCHIVO:
			logs.Agregar_archivo(elementos[1])
			listaIPs := logs.DOS()
			if len(listaIPs) != 0 {
				for _, ip := range listaIPs {
					fmt.Fprintf(os.Stdout, "DoS: %s", ip)
				}
			}
			fmt.Fprintln(os.Stdout, "OK")
		case _VER_VISITANTES:
			visitantes := logs.Ver_visitantes(elementos[1], elementos[2])
			fmt.Fprintf(os.Stdout, "%s\n", "Visitantes:")
			for _, visitante := range visitantes {
				fmt.Fprintf(os.Stdout, "\t%s\n", visitante)

			}
			fmt.Fprintln(os.Stdout, "OK")
		case _VER_MAS_VISITADOS:
			n, _ := strconv.Atoi(elementos[1])
			mas_visitados := logs.Ver_mas_visitados(n)
			fmt.Fprintf(os.Stdout, "%s\n", "Sitios más visitados:")
			for _, visitado := range mas_visitados {
				sitio, visitas := visitado.Ver_par()
				fmt.Fprintf(os.Stdout, "\t%s - %d\n", sitio, visitas)
			}
			fmt.Fprintln(os.Stdout, "OK")
		default:
			fmt.Fprintf(os.Stderr, "Error en comando %s", comando)
			return
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stdout, "ERROR", err)
		os.Exit(1)
	}
}
