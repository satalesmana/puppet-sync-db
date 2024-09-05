package puppet

import "go.mongodb.org/mongo-driver/bson/primitive"

type SmtpAccount struct {
	ID           primitive.ObjectID `bson:"_id"`
	HostName     string             `bson:"host_name"`
	Port         int                `bson:"port"`
	AuthEmail    string             `bson:"auth_email"`
	AuthPassword string             `bson:"auth_password"`
	Status       string             `bson:"status"`
}
