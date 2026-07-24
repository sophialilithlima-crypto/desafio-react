package utils

import (
	"strings"
)


func ValidarTexto(valor string) bool {

	return strings.TrimSpace(valor) != ""

}