package comandos

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADic "tdas/diccionario"
	"time"
)

type detectorLogs struct {
	visitantes       TDADic.DiccionarioOrdenado[string, []time.Time]
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
		visitantes:       TDADic.CrearABB[string, []time.Time](compararIp),
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
		tiempo, _ := time.Parse("2006-01-02T15:04:05-07:00", elementos[1])

		if !detector.visitantes.Pertenece(ip) {
			detector.visitantes.Guardar(ip, []time.Time{})
		} else {
			detector.visitantes.Guardar(ip, append(detector.visitantes.Obtener(ip), tiempo))
		}

		sitio := elementos[3]

		if !detector.sitios_visitados.Pertenece(sitio) {
			detector.sitios_visitados.Guardar(sitio, 1)
		} else {
			detector.sitios_visitados.Guardar(sitio, detector.sitios_visitados.Obtener(sitio)+1)
		}

	}

}

func (detector *detectorLogs) DOS() []string {
	var dos []string
	for iter := detector.visitantes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		ip, listaTiempo := iter.VerActual()
		if esMenorADos(listaTiempo) {
			dos = append(dos, ip)
		}
	}
	return dos
}

func esMenorADos(listaTiempo []time.Time) bool {
	if len(listaTiempo) >= 5 {
		inicio := 0
		fin := 4
		for fin < len(listaTiempo) {
			diferencia := listaTiempo[fin].Sub(listaTiempo[inicio])
			segundos := diferencia.Seconds()
			if segundos < 2 {
				return true
			}
			inicio++
			fin++
		}
	}
	return false

}

func (detector *detectorLogs) Ver_visitantes(desde string, hasta string) []string {

	visitantesEnRango := []string{}

	visitar := func(clave string, dato []time.Time) bool {
		visitantesEnRango = append(visitantesEnRango, clave)
		return true
	}

	detector.visitantes.IterarRango(&desde, &hasta, visitar)

	return visitantesEnRango
}

func (detector *detectorLogs) Ver_mas_visitados(n int) []parSitioVisitas {

	var resultado []parSitioVisitas

	mas_visitados := TDAHeap.CrearHeap(func(a, b parSitioVisitas) int {
		if a.visitas > b.visitas {
			return a.visitas
		}
		return b.visitas
	})

	for iter := detector.sitios_visitados.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		sitio, visitas := iter.VerActual()
		mas_visitados.Encolar(*crearParSitioVisitas(sitio, visitas))
	}

	for i := 0; i < n && !mas_visitados.EstaVacia(); i++ {
		resultado = append(resultado, mas_visitados.Desencolar())
	}

	return resultado
}

func (par *parSitioVisitas) Ver_par() (string, int) {
	return par.sitio, par.visitas
}
