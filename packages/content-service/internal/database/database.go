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
	log.Printf("Connection string: %v", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	// log.Printf("Migrating Image")
	// err = DB.AutoMigrate(&Image{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Migrating tracks")
	// err = DB.AutoMigrate(&Track{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Migrating Genre")
	// err = DB.AutoMigrate(&Genre{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Migrating Artist")
	// err = DB.AutoMigrate(&Artist{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Migrating Album")
	// err = DB.AutoMigrate(&Album{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Auto Migrate
	err = DB.AutoMigrate(&Track{}, &Artist{}, &Album{}, &Image{}, &Genre{})
	if err != nil {
		log.Fatal(err)
	}

	// Add foreign key constraint for Album - Artist many-to-many relationship
	migrator := DB.Migrator()
	err = migrator.CreateConstraint(&Album{}, "Artists", &gorm.Constraint{
		Name:             "fk_album_artist", // Constraint name
		ForeignDBName:    "artists",         // Name of the related table
		ForeignFieldName: "id",              // Name of the related table's primary key
		UpdateRule:       gorm.Restrict,     // Update rule (e.g. cascade, restrict, set null)
		DeleteRule:       gorm.Cascade,      // Delete rule (e.g. cascade, restrict, set null)
	})

	if err != nil {
		log.Fatal(err)
	}

	// Add foreign key constraint for Album - Genre many-to-many relationship
	err = migrator.CreateConstraint(&Album{}, "Genres", &gorm.Constraint{
		Name:             "fk_album_genre",
		ForeignDBName:    "genres",
		ForeignFieldName: "id",
		UpdateRule:       gorm.Restrict,
		DeleteRule:       gorm.Cascade,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Add foreign key constraint for Artist - Genre many-to-many relationship
	err = migrator.CreateConstraint(&Artist{}, "Genres", &gorm.Constraint{
		Name:             "fk_artist_genre",
		ForeignDBName:    "genres",
		ForeignFieldName: "id",
		UpdateRule:       gorm.Restrict,
		DeleteRule:       gorm.Cascade,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Add foreign key constraint for image many to many
	// Add foreign key constraint for Album - Image many-to-many relationship
	err = migrator.CreateConstraint(&Album{}, "Images", &gorm.Constraint{
		Name:             "fk_album_image",
		ForeignDBName:    "images",
		ForeignFieldName: "id",
		UpdateRule:       gorm.Restrict,
		DeleteRule:       gorm.Cascade,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Add foreign key constraint for Artist - Image many-to-many relationship
	err = migrator.CreateConstraint(&Artist{}, "Images", &gorm.Constraint{
		Name:             "fk_artist_image",
		ForeignDBName:    "images",
		ForeignFieldName: "id",
		UpdateRule:       gorm.Restrict,
		DeleteRule:       gorm.Cascade,
	})

	if err != nil {
		log.Fatal(err)
	}
}
