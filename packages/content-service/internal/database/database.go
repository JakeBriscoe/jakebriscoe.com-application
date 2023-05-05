package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var DB *gorm.DB

func ConnectDB() {

	// Get the Kubernetes configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	secretNamespace := os.Getenv("MY_NAMESPACE")
	secretName := "content-db-credentials"
	secret, err := clientset.CoreV1().Secrets(secretNamespace).Get(context.Background(), secretName, metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	username := string(secret.Data["username"])
	password := string(secret.Data["password"])

	log.Print("Connecting to db")
	dsn := fmt.Sprintf("host=postgres-service.dev.svc.cluster.local user=%v password=%v dbname=content_db port=5432 sslmode=disable TimeZone=Europe/London", username, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	// Auto Migrate
	err = DB.AutoMigrate(&Track{}, &Artist{}, &Album{}, &Image{}, &Genre{})
	if err != nil {
		log.Fatal(err)
	}
}
