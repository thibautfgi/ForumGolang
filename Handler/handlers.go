package Handler

import (
	"database/sql"
	"fmt"
	"forum/GestionDatabase"
	"forum/Outils"
	"log"
	"net/http"

	"strconv"
	"text/template"
)

var mytestFinal GestionDatabase.Forum

func PagePrincipal(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" { //demarre par default sur cette url
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		db, err := sql.Open("sqlite3", "./forum.sqlite")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		idTopicOpenInt := 999

		cookie, err := r.Cookie("session_id")

		if err != nil {
			// Le cookie n'est pas présent
			GestionDatabase.UpdateNotConnect(db, 1)
			fmt.Println(cookie)
		} else {
			GestionDatabase.UpdateNotConnect(db, -1)
			cookievalue := cookie.Value
			fmt.Println(cookievalue)
			toy, _ := GestionDatabase.GetIdFromUUID(db, cookievalue)
			GestionDatabase.UpdateIdConnect(db, toy)
		}

		fmt.Println("charge")
		GestionDatabase.UpdateTopicId(db, idTopicOpenInt)

		mytest := GestionDatabase.QueryDbUtilisateur(db, GestionDatabase.Utilisateur{}) //envoie les donnes dans ma struct
		mytest2 := GestionDatabase.QueryDbTopic(db, GestionDatabase.Topic{})
		mytest3 := GestionDatabase.QueryDbMessage(db, GestionDatabase.Message{})
		mytest4 := GestionDatabase.QueryDbLikeNumber(db, GestionDatabase.LikeNumber{})
		mytest5 := GestionDatabase.QueryDbTest(db, GestionDatabase.Test{})
		mytest6 := GestionDatabase.QueryDbSessions(db, GestionDatabase.Sessions{})

		mytestFinal = GestionDatabase.Forum{mytest, mytest2, mytest3, mytest4, mytest5, mytest6}

		fmt.Println("here1", mytestFinal.Tests)

		t, err := template.ParseFiles("./static/html/pageMenu.html") // le POST prend les infos de la templates pageMenu
		if err != nil {
			fmt.Println(err)
		}

		t.Execute(w, mytestFinal) // ici le deuxieme arg permet de cree une gotemplate avec les infos de mytestfinal

	case "POST":
		db, err := sql.Open("sqlite3", "./forum.sqlite")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		cookie, err := r.Cookie("session_id")

		if err != nil {
			// Le cookie n'est pas présent
			GestionDatabase.UpdateNotConnect(db, 1)
			fmt.Println(cookie)
		} else {
			GestionDatabase.UpdateNotConnect(db, -1)
			cookievalue := cookie.Value
			fmt.Println(cookievalue)
			toy, _ := GestionDatabase.GetIdFromUUID(db, cookievalue)
			GestionDatabase.UpdateIdConnect(db, toy)
		}

		testytest := r.FormValue("newcomment")
		commentdata := r.FormValue("commentData")

		commentdataInt, _ := strconv.Atoi(commentdata)

		iduser_connecter := 7 //LIER ICI A LUSER CONNECTE
		dates_test := Outils.Date()

		alredylike := r.FormValue("testlike")

		idTopicOpen := r.FormValue("topicOpen")
		idTopicOpenInt, _ := strconv.Atoi(idTopicOpen)

		fmt.Println(idTopicOpenInt)

		if idTopicOpen != "" {
			GestionDatabase.UpdateTopicId(db, idTopicOpenInt)
		}

		if alredylike != "" { //cree un like
			mytestuwu := GestionDatabase.QueryDbLikeNumber(db, GestionDatabase.LikeNumber{})

			alredy1, alredy2 := Outils.ParseUserIdAndLikeNumber(alredylike)

			creatstruct := GestionDatabase.LikeNumber{alredy1, alredy2}

			// fmt.Println("MessageIdLiked : ", alredy1)
			// fmt.Println("UserID : ", alredy2)
			if Outils.TestMatchMsg(mytestuwu, creatstruct) == false { //permet de like si un like a pas deja ete fais et l'enleve si reclick
				GestionDatabase.NewUserLikeMessage(db, alredy1, alredy2)
				GestionDatabase.UpdateLikesByOne(db, alredy1)
			} else {
				GestionDatabase.RemoveLikes(db, alredy1, alredy2)
				GestionDatabase.UpdateLikesMinusOne(db, alredy1)
			}

		}

		if testytest != "" { //cree un comment
			GestionDatabase.CreateMessage(db, GestionDatabase.GenerateUniqueID(db), commentdataInt, iduser_connecter, testytest, dates_test, 0) // test de creation d'un user
		}

		mytest := GestionDatabase.QueryDbUtilisateur(db, GestionDatabase.Utilisateur{}) //envoie les donnes dans ma struct
		mytest2 := GestionDatabase.QueryDbTopic(db, GestionDatabase.Topic{})
		mytest3 := GestionDatabase.QueryDbMessage(db, GestionDatabase.Message{})
		mytest4 := GestionDatabase.QueryDbLikeNumber(db, GestionDatabase.LikeNumber{})
		mytest5 := GestionDatabase.QueryDbTest(db, GestionDatabase.Test{})
		mytest6 := GestionDatabase.QueryDbSessions(db, GestionDatabase.Sessions{})

		mytestFinal = GestionDatabase.Forum{mytest, mytest2, mytest3, mytest4, mytest5, mytest6}

		fmt.Println("here2", mytestFinal.Tests)

		t, err := template.ParseFiles("./static/html/pageMenu.html") // le POST prend les infos de la templates pagelettre
		if err != nil {
			fmt.Println(err)
		}

		t.Execute(w, mytestFinal)
	}
}

func PageConnect(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/pageConnect" { //demarre par default sur cette url
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":

		db, err := sql.Open("sqlite3", "./forum.sqlite")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		mytest := GestionDatabase.QueryDbUtilisateur(db, GestionDatabase.Utilisateur{}) //envoie les donnes dans ma struct
		mytest2 := GestionDatabase.QueryDbTopic(db, GestionDatabase.Topic{})
		mytest3 := GestionDatabase.QueryDbMessage(db, GestionDatabase.Message{})
		mytest4 := GestionDatabase.QueryDbLikeNumber(db, GestionDatabase.LikeNumber{})
		mytest5 := GestionDatabase.QueryDbTest(db, GestionDatabase.Test{})
		mytest6 := GestionDatabase.QueryDbSessions(db, GestionDatabase.Sessions{})

		mytestFinal = GestionDatabase.Forum{mytest, mytest2, mytest3, mytest4, mytest5, mytest6}

		t, err := template.ParseFiles("./static/html/pageConnect.html") // le POST prend les infos de la templates pagelettre
		if err != nil {
			fmt.Println(err)
		}

		t.Execute(w, mytestFinal) // ici le deuxieme arg permet de cree une template

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err) // LA Templates choix POST CES INFOS
			return
		}

		db, err := sql.Open("sqlite3", "./forum.sqlite")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		NomUser := r.FormValue("NomUtilisateurTestInput")
		MdpUser := r.FormValue("MdpTestInput")
		EmailUser := r.FormValue("EmailTestInput")

		GestionDatabase.AddUser(db, NomUser, MdpUser, EmailUser, "/static/phtotos/moon")

		TryUser := r.FormValue("NomConnect")
		TryMdp := r.FormValue("MdpConnect")

		GestionDatabase.CreateSession(w, db, TryUser, TryMdp) // cookie

		mytest := GestionDatabase.QueryDbUtilisateur(db, GestionDatabase.Utilisateur{}) //envoie les donnes dans ma struct
		mytest2 := GestionDatabase.QueryDbTopic(db, GestionDatabase.Topic{})
		mytest3 := GestionDatabase.QueryDbMessage(db, GestionDatabase.Message{})
		mytest4 := GestionDatabase.QueryDbLikeNumber(db, GestionDatabase.LikeNumber{})
		mytest5 := GestionDatabase.QueryDbTest(db, GestionDatabase.Test{})
		mytest6 := GestionDatabase.QueryDbSessions(db, GestionDatabase.Sessions{})

		mytestFinal = GestionDatabase.Forum{mytest, mytest2, mytest3, mytest4, mytest5, mytest6}

		t, err := template.ParseFiles("./static/html/pageMenu.html") // le POST prend les infos de la templates pagelettre
		if err != nil {
			fmt.Println(err)
		}

		t.Execute(w, mytestFinal)
	}
}

func PagePost(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/pagePost" { //demarre par default sur cette url
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":

		db, err := sql.Open("sqlite3", "./forum.sqlite")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		mytest := GestionDatabase.QueryDbUtilisateur(db, GestionDatabase.Utilisateur{}) //envoie les donnes dans ma struct
		mytest2 := GestionDatabase.QueryDbTopic(db, GestionDatabase.Topic{})
		mytest3 := GestionDatabase.QueryDbMessage(db, GestionDatabase.Message{})
		mytest4 := GestionDatabase.QueryDbLikeNumber(db, GestionDatabase.LikeNumber{})
		mytest6 := GestionDatabase.QueryDbSessions(db, GestionDatabase.Sessions{})

		mytestFinal = GestionDatabase.Forum{mytest, mytest2, mytest3, mytest4, mytestFinal.Tests, mytest6}

		t, err := template.ParseFiles("./static/html/pagePost.html") // le POST prend les infos de la templates pagelettre
		if err != nil {
			fmt.Println(err)
		}

		t.Execute(w, mytestFinal) // ici le deuxieme arg permet de cree une template

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err) // LA Templates choix POST CES INFOS
			return
		}

		db, err := sql.Open("sqlite3", "./forum.sqlite")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		Titre := r.FormValue("titreTestInput")
		Content := r.FormValue("contentTestInput")
		Lien := r.FormValue("lienTestInput")

		if Titre != "" && Content != "" {
			GestionDatabase.CreateSujet(db, 7, Titre, Outils.Date(), 0, Content, Outils.MultimediaConverteur(Lien), Outils.TchekImage(Lien), Outils.GetIdImageFromYt(Lien))
		}

		mytest := GestionDatabase.QueryDbUtilisateur(db, GestionDatabase.Utilisateur{}) //envoie les donnes dans ma struct
		mytest2 := GestionDatabase.QueryDbTopic(db, GestionDatabase.Topic{})
		mytest3 := GestionDatabase.QueryDbMessage(db, GestionDatabase.Message{})
		mytest4 := GestionDatabase.QueryDbLikeNumber(db, GestionDatabase.LikeNumber{})
		mytest6 := GestionDatabase.QueryDbSessions(db, GestionDatabase.Sessions{})

		mytestFinal = GestionDatabase.Forum{mytest, mytest2, mytest3, mytest4, mytestFinal.Tests, mytest6}

		t, err := template.ParseFiles("./static/html/pageMenu.html") // le POST prend les infos de la templates pagelettre
		if err != nil {
			fmt.Println(err)
		}

		t.Execute(w, mytestFinal)
	}
}
