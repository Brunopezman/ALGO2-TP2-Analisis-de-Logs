package comandos

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADic "tdas/diccionario"
)

type detectorLogs struct {
	visitantes       TDADic.DiccionarioOrdenado[string, []string]
	sitios_visitados TDADic.Diccionario[string, int]
}

type parSitioVisitas struct {
	sitio   string
	visitas int
}

func crearParSitioVisitas(sitio string, visitas int) *parSitioVisitas {
	return &parSitioVisitas{sitio: sitio, visitas: visitas}
}

func CrearDetectorDeLogs() ServidorLogs {

	return &detectorLogs{
		visitantes:       TDADic.CrearABB[string, []string](compararIp),
		sitios_visitados: TDADic.CrearHash[string, int](),
	}
}

func compararIp(ip1, ip2 string) int {
	valoresIp1 := strings.Split(ip1, ".")
	valoresIp2 := strings.Split(ip2, ".")

	for i := range valoresIp1 {
		valorIp1, _ := strconv.Atoi(valoresIp1[i])
		valorIp2, _ := strconv.Atoi(valoresIp2[i])

		if valorIp1 > valorIp2 {
			return 1
		}
		if valorIp1 < valorIp2 {
			return -1
		}
	}
	return 0
}

func (detector *detectorLogs) Agregar_archivo(ruta string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		expresion := scanner.Text()
		elementos := strings.Fields(expresion)
		ip := elementos[0]
		tiempo := elementos[1]

		if !detector.visitantes.Pertenece(ip) {
			detector.visitantes.Guardar(ip, []string{})
		} else {
			detector.visitantes.Guardar(ip, append(detector.visitantes.Obtener(ip), tiempo))
		}

		sitio := elementos[2]

		if !detector.sitios_visitados.Pertenece(sitio) {
			detector.sitios_visitados.Guardar(sitio, 1)
		} else {
			detector.sitios_visitados.Guardar(sitio, detector.sitios_visitados.Obtener(sitio)+1)
		}

	}

}

func (detector *detectorLogs) Ver_visitantes(desde string, hasta string) []string {

	visitantesEnRango := []string{}

	visitar := func(clave string, dato []string) bool {
		visitantesEnRango = append(visitantesEnRango, clave)
		return true
	}

	detector.visitantes.IterarRango(&desde, &hasta, visitar)

	return visitantesEnRango
}

func (detector *detectorLogs) Ver_mas_visitados(n int) []parSitioVisitas {

	var resultado []parSitioVisitas

	mas_visitados := TDAHeap.CrearHeap(func(a, b parSitioVisitas) int {
		if a.visitas-b.visitas > 0 {
			return a.visitas
		}
		return b.visitas
	})

	for iter := detector.sitios_visitados.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		sitio, cantidad := iter.VerActual()
		mas_visitados.Encolar(*crearParSitioVisitas(sitio, cantidad))
	}

	for i := 0; i < n && !mas_visitados.EstaVacia(); i++ {
		resultado = append(resultado, mas_visitados.Desencolar())
	}

	return resultado
}
