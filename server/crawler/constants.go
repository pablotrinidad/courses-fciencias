package crawler

const (
	baseURL   = "http://www.fciencias.unam.mx"
	majorsURL = "licenciatura/resumen"
)

var majorsList = map[int]string{
	101: "Actuariía",
	201: "Biología",
	104: "Ciencias de la Computacioón",
	127: "Ciencias de la Tierra",
	106: "Física",
	134: "Física Biomeédica",
	122: "Matemáticas",
	136: "Matemáticas Aplicadas",
}