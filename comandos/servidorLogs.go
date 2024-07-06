package comandos

type DetectorLogs interface {
	// Procesa de forma completa un archivo de log e Identifica posibles situaciones de sobrecarga de solicitudes de un mismo cliente
	// en un tiempo breve sobre un mismo servidor.
	AgregarArchivo(string) ([]string, error)

	//muestra todas las IPs que solicitaron algún recurso en el servidor, dentro del rango de IPs determinado.
	VerVisitantes(desde string, hasta string) []string

	// Muestra los n recursos más solicitados.
	VerMasVisitados(n int) []parSitioVisitas
}
