package config
import (
	"context"
	"log"
	"os"
	
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// FirebaseAuth adalah instance Firebase Auth yang dipakai untuk verify token
	var FirebaseAuth *auth.Client
	func InitFirebase() {
	credPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")
// Inisialisasi Firebase App dengan service account credentials
	opt := option.WithCredentialsFile(credPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
	log.Fatalf("Gagal init Firebase: %v", err)
}