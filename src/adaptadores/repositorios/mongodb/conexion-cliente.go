package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//db_url=mongodb+srv://<user>:<password>@<host:port>/

func ConexionCliente(db_url string) (cliente *mongo.Client, err error) {
	//Se configura el tiempo de espera antes de cancelar la conexión.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//Conexión al servidor de MongoDB

	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(db_url))
	if err != nil {
		return nil, err
	}

	//Verificar conexión
	err = cliente.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return cl, nil

}
