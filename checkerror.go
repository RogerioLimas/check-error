package checkerror

import "log"

// CheckError Verifica se o objeto error está nulo. Se não estiver, imprime a mensagem de erro passada, junto com a stack do error
func CheckError(err error, message string) {
	if err != nil {
		log.Fatalf("%s\n%v\n\n", message, err)
	}
}
