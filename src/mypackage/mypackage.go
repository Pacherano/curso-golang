package mypackage

//CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

//PrintMessage imprime un mensaje
func PrintMessage(texto string) {
	println(texto)
}
