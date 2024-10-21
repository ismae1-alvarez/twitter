package bd

import (
	"context"
	"time"
	"twitter/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro es la parad final con la bd para insertar los dato del usuario*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitter")

	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	// Obtener id en string
	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
