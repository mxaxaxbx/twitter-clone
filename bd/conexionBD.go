package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
/* objeto de conexion a la base de datos */
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://user:BQSLTpaF0p9oGc1c@cluster0.gdt1r.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
/* conexion a la base de datos */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa a la BD")

	return client
	
}
/* ping a la base de datos */
func ChequeoConection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
