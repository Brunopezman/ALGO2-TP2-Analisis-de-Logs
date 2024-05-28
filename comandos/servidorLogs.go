package comandos

type ServidorLogs interface {
	// Procesa de forma completa un archivo de log
	Agregar_archivo(string)

	//muestra todas las IPs que solicitaron algún recurso en el servidor, dentro del rango de IPs determinado.
	Ver_visitantes(desde string, hasta string) []string

	// Muestra los n recursos más solicitados.
	Ver_mas_visitados(n int) []parSitioVisitas
}
