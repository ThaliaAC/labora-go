package main

import (
	"fmt"
	"strings"
	"time"
)

type estructuraMensaje struct {
	tipo      string
	contenido string
}

/*const (
    detenido = 0
    pausado = 1
    ejecutando = 2
)*/

func main() {
	canalEmail := make(chan estructuraMensaje)
	canalSms := make(chan estructuraMensaje)
	canalPush := make(chan estructuraMensaje)
	//estadoCanal := make(chan bool)

	for {
		var nuevoMensaje estructuraMensaje
		fmt.Println("[Para salir escriba OUT]")
		fmt.Println("Ingrese el tipo de Mensaje [email / sms / push]: ")
		fmt.Scan(&nuevoMensaje.tipo)
		if strings.ToUpper(nuevoMensaje.tipo) == "OUT" {
			break
		}
		if !validarTipo(nuevoMensaje.tipo) {
			fmt.Println("==Tipo inválido, intente nuevamente==")
			continue
		}
		fmt.Println("Ingrese el contenido del Mensaje: ")
		fmt.Scan(&nuevoMensaje.contenido)
		if strings.ToUpper(nuevoMensaje.tipo) == "EMAIL" {
			go enviarEmail(nuevoMensaje, canalEmail)
			go func() {
				for mensaje := range canalEmail {
					fmt.Println("Mensaje recibido", mensaje)
				}
			}()
		}
		if strings.ToUpper(nuevoMensaje.tipo) == "SMS" {
			go enviarSms(nuevoMensaje, canalSms)
			go func() {
				for mensaje := range canalSms {
					fmt.Println("Mensaje recibido", mensaje)
				}
			}()
		}
		if strings.ToUpper(nuevoMensaje.tipo) == "PUSH" {
			go enviarPush(nuevoMensaje, canalPush)
			go func() {
				for mensaje := range canalPush {
					fmt.Println("Mensaje recibido", mensaje)
				}
			}()
		}
	}
	defer close(canalEmail)
	defer close(canalSms)
	defer close(canalPush)

}

func enviarEmail(nuevoMensaje estructuraMensaje, canalEmail chan estructuraMensaje) {
	time.Sleep(10 * time.Second)
	canalEmail <- nuevoMensaje
}

func enviarSms(nuevoMensaje estructuraMensaje, canalSms chan estructuraMensaje) {
	time.Sleep(10 * time.Second)
	canalSms <- nuevoMensaje
}
func enviarPush(nuevoMensaje estructuraMensaje, canalPush chan estructuraMensaje) {
	time.Sleep(10 * time.Second)
	canalPush <- nuevoMensaje
}

func validarTipo(tipo string) bool {
	tipoValido := strings.ToUpper(tipo) == "EMAIL" || strings.ToUpper(tipo) == "SMS" || strings.ToUpper(tipo) == "PUSH"
	return tipoValido
}
