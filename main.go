package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	comandos "tp2/comandos"
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
		case _VER_VISITANTES:
			if len(elementos) == 1 {
				logs.Ver_visitantes("0.0.0.0", "255.255.255.255")
				fmt.Fprintln(os.Stdout, "OK")
				continue
			}
			logs.Ver_visitantes(elementos[1], elementos[2])

		case _VER_MAS_VISITADOS:
			n, _ := strconv.Atoi(elementos[1])
			mas_visitados := logs.Ver_mas_visitados(n)

		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stdout, "ERROR", err)
		os.Exit(1)
	}
}
