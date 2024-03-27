package main

import (
	"database/sql"
	"fmt"

	// "forum/GestionDatabase"

	// "forum/GestionDatabase"

	// "forum/GestionDatabase"
	// "forum/GestionDatabase"

	// "forum/GestionDatabase"

	"forum/Handler"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const port = ":8080"

func main() {

	// Ouvre une connexion à la base de données forum.sqlite avec sqlite3
	db, err := sql.Open("sqlite3", "./forum.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a Version 4 UUID.
	// fmt.Println(GestionDatabase.CheckIfUsernameExists(db, "Patrick"))7

	// GestionDatabase.Hash(db, 7)

	// GestionDatabase.UpdateIdConnect(db, 7)

	// mytest := GestionDatabase.QueryDbMessage(db, GestionDatabase.Message{}) //envoie les donnes dans ma struct

	// linkMultimedia := "https://www.youtube.com/watch?v=mCigvageQlA&t=3395s&ab_channel=AmbientCinematics"

	//ICI YT VIDEO OU LINK IMG
	// GestionDatabase.CreateSujet(db, 7, "UN TOPIC CREE DANS LE MAIN", "12/06/03", 0, "OKCEC test test EEEEEEEE", Outils.MultimediaConverteur(linkMultimedia), Outils.TchekImage(linkMultimedia), Outils.GetIdImageFromYt(linkMultimedia))
	// ICI CREE UN USER
	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	log.Printf("Génération de la Version 4 UUID %v", u2)

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", Handler.PagePrincipal)
	http.HandleFunc("/pageConnect", Handler.PageConnect)
	http.HandleFunc("/pagePost", Handler.PagePost)

	fmt.Println("Mon site est en ligne! Port", port+".")
	fmt.Println("Mon adresse site à copier : ")
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(port, nil)

	// fmt.Println(GestionDatabase.GetMessageFromUserId(db, -6649866525117810327))
	// fmt.Println(GestionDatabase.GetUsernameFromMessageId(db, 0))

}
