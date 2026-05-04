package main

import (
	"capass/src/adaptadores/controladores/env"
	controlador_error "capass/src/adaptadores/controladores/error"
	"fmt"
	"log"
	"os"
)

func main() {
	if !inicio_variablesEntorno() {
		err_ := controlador_error.NuevoControladorError()
		err_.Archivo = "main.go"
		err_.Modulo = "inicio_variablesEntorno"
		err_.Mensaje = "Error al iniciar variables de entorno"
		fmt.Println(err_)
		os.Exit(1)
	}

	fmt.Println("Capass se inicio correctamente")

}

func inicio_variablesEntorno() bool {
	//Probando variables de entorno
	env_ := "./.env"
	env_copia := "./src/adaptadores/controladores/env/env.copia"
	if env.VerificarEnv(env_, env_copia) {
		estado_env, err := env.GetEnv("estado_env", env_)
		if err == nil && estado_env == "1" {
			fmt.Println("Archivo env: " + env_ + " - cargado correctamente.")
		} else {
			if err == nil {
				log.Println("Verificar configuración de archivo env")
			} else {
				log.Println("Verificar configuración de archivo env: " + err.Error())
			}
			return false

		}
	}

	u_mdb, _ := env.GetEnv("user_mongodb", env_)
	pw_mdb, _ := env.GetEnv("pw_mongodb", env_)
	h_mdb, _ := env.GetEnv("host_mongodb", env_)
	p_mdb, _ := env.GetEnv("port_mongodb", env_)

	mongo_url := "mongodb://" + u_mdb + ":" + pw_mdb + "@" + h_mdb + ":" + p_mdb + "/"
	fmt.Println("Cadena de conexión mongodb: " + mongo_url)
	return true
}
