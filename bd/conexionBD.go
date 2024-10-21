package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN el objeto de a la base de datos*/
var MongoCN = ConectarBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://ismaelalvarez514:mRAFL76HzuLJmwGt@cluster0.bk3go.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

/* ConectarBD es la funcion qeu me permite conectar la Bd*/
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

	log.Println("Conexion Exitosa con la BD")
	return client
}

/**ChequeoConnection es la funcion para ver la conexion a al BD**/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
