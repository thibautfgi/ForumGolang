package GestionDatabase

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/binary"
	"errors"
	"net/http"
	"time"

	// "fmt"
	"log"

	"github.com/gofrs/uuid"
)

// Fonctionalités tierces

// Fonction générant un ID unique

func GenerateUniqueID(db *sql.DB) int {
	// génère un entier aléatoire de 64 bits
	var id int
	b := make([]byte, 8)
	rand.Read(b)
	// convertit le tableau de bytes en entier
	id = int(binary.LittleEndian.Uint64(b))

	// vérifie si l'ID existe déjà dans la base de données

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM utilisateurs WHERE Id_utilisateur = ?", id).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	// si l'ID existe déjà, on rappelle la fonction
	if count > 0 {
		return GenerateUniqueID(db)
	}
	return id
}

// Fonction hashant un mot de passe

func HashPassword(password string) (string, error) {
	h := sha256.New()
	h.Write([]byte(password))
	hashed := h.Sum(nil)
	return string(hashed), nil
}

// Fonctionalités de vérification de la base de données

// Fonction qui vérifie si un nom d'utilisateur existe déjà dans la base de données
func CheckIfUsernameExists(db *sql.DB, nomUtilisateur string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM utilisateurs WHERE Nom_utilisateur = ?", nomUtilisateur).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

// Fonction qui vérifie si un email existe déjà dans la base de données

func CheckIfEmailExists(db *sql.DB, email string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM utilisateurs WHERE Email_utilisateur = ?", email).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

func CheckIfMsgLiked(db *sql.DB, IdUser int) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM likeNumber WHERE Id_utilisateur_likeNumber = ?", IdUser).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

// Fonction qui vérifie si un nom d'utilisateur et un mot de passe correspondent dans la base de données

func CheckIfUsernameAndPasswordMatch(db *sql.DB, nomUtilisateur string, motDePasse string) bool {
	var count int
	hash, _ := HashPassword(motDePasse)
	err := db.QueryRow("SELECT COUNT(*) FROM utilisateurs WHERE Nom_utilisateur = ? AND Mdp_utilisateur = ?", nomUtilisateur, hash).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

// Fonction pour vérifier si le nom d'utilisateur et le mot de passe correspondent

func CheckUser(db *sql.DB, nomUtilisateur string, motDePasse string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM utilisateurs WHERE Nom_utilisateur = ? AND Mdp_utilisateur = ?", nomUtilisateur, motDePasse).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

// Fonctionalités de récupération de données de la base de données

// Récupère le message en fonction de l'id du message

func GetMessageFromId(db *sql.DB, id int) (string, error) {
	var message string
	err := db.QueryRow("SELECT contenu_message FROM messages WHERE id_message = ?", id).Scan(&message)
	if err != nil {
		log.Fatal(err)
	}
	return message, nil
}

func GetUUIDFromID(db *sql.DB, id int) (string, error) {
	var message string
	err := db.QueryRow("SELECT UUID_session FROM sessions WHERE Id_utilisateur_session = ?", id).Scan(&message)
	if err != nil {
		log.Fatal(err)
	}
	return message, nil
}

func GetIdFromUUID(db *sql.DB, uuid string) (int, error) {
	var message int
	err := db.QueryRow("SELECT Id_utilisateur_session FROM sessions WHERE UUID_session = ?", uuid).Scan(&message)
	if err != nil {
		log.Fatal(err)
	}
	return message, nil
}

// Récupère le contenu du message en fonction de l'id de l'utilisateur

func GetMessageFromUserId(db *sql.DB, id int) (string, error) {
	var message string
	err := db.QueryRow("SELECT contenu FROM messages WHERE Id_utilisateur = ?", id).Scan(&message)
	if err != nil {
		log.Fatal(err)
	}
	return message, nil
}

// Récupère le nom de l'utilisateur en fonction du l'id du message

func GetUsernameFromMessageId(db *sql.DB, id int) (string, error) {
	var username string
	err := db.QueryRow("SELECT Nom_utilisateur FROM utilisateurs WHERE Id_message = (SELECT Id_utilisateur FROM messages WHERE Id_utilisateur = ?)", id).Scan(&username)
	if err != nil {
		log.Fatal(err)
	}
	return username, nil
}

// Fonctionalités de l'utilisateur

// Ajoute l'utilisateur à la base de données

func AddUser(db *sql.DB, nomUtilisateur string, motDePasse string, email string, avatar string) {
	_, err := db.Exec("INSERT INTO utilisateurs ( Nom_utilisateur, Mdp_utilisateur, Email_utilisateur, Avatar_utilisateur) VALUES ( ?, ?, ?, ?)", nomUtilisateur, motDePasse, email, avatar)
	if err != nil {
		log.Fatal(err)
	}
}

// Créer un message et l'enregistre dans la base de données

func CreateMessage(db *sql.DB, id int, id_sujet int, id_utilisateur int, contenu string, date_creation string, likes int) {
	_, err := db.Exec("INSERT INTO messages (Id_message,Id_topic_msg,Id_utilisateur_msg, Contenu_message,Date_message, Likes_message) VALUES (?, ?, ?, ?,?,  ?)", id, id_sujet, id_utilisateur, contenu, date_creation, likes)
	if err != nil {
		log.Fatal(err)
	}
}

// Créer un sujet et l'enregistre dans la base de données

func CreateSujet(db *sql.DB, Id_utilisateur_topic int, Titre_topic string, Date_topic string, Likes_topic int, Contenue_topic string, Multimedia_topic string, Imgtchek_topic int, MiniatureYt_topic string) {
	_, err := db.Exec("INSERT INTO topic (Id_utilisateur_topic,Titre_topic, Date_topic,Likes_topic, Contenu_topic, Multimedia_topic, Imgtchek_topic,MiniatureYt_topic) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)", Id_utilisateur_topic, Titre_topic, Date_topic, Likes_topic, Contenue_topic, Multimedia_topic, Imgtchek_topic, MiniatureYt_topic)
	if err != nil {
		log.Fatal(err)
	}
}

// Modifier un message dans la base de données

func UpdateMessage(db *sql.DB, id int, contenu string) {
	_, err := db.Exec("UPDATE messages SET Contenu_message = ? WHERE Id_message = ?", contenu, id)
	if err != nil {
		log.Fatal(err)
	}
}

// Modifier le nom d'utilisateur dans la base de données

func UpdateUsername(db *sql.DB, id int, nomUtilisateur string) {
	_, err := db.Exec("UPDATE utilisateurs SET Nom_utilisateur = ? WHERE Id_utilisateur = ?", nomUtilisateur, id)
	if err != nil {
		log.Fatal(err)
	}
}

// Modifier le mot de passe dans la base de données

func UpdatePassword(db *sql.DB, id int, motDePasse string) {
	_, err := db.Exec("UPDATE utilisateurs SET Mdp_utlisateur = ? WHERE Id_utilisateur = ?", motDePasse, id)
	if err != nil {
		log.Fatal(err)
	}
}

// Modifier l'email dans la base de données

func UpdateEmail(db *sql.DB, id int, email string) {
	_, err := db.Exec("UPDATE utilisateurs SET Email_utilisateur = ? WHERE Id_utilisateur = ?", email, id)
	if err != nil {
		log.Fatal(err)
	}
}

// Modifier l'avatar dans la base de données

func UpdateAvatar(db *sql.DB, id int, avatar string) {
	_, err := db.Exec("UPDATE Utilisateurs SET Avatar_utilisateur = ? WHERE Id_utilisateur = ?", avatar, id)
	if err != nil {
		log.Fatal(err)
	}
}

//update

func UpdateLikesByOne(db *sql.DB, id int) {
	_, err := db.Exec("UPDATE messages SET Likes_message = Likes_message +1 WHERE Id_message  = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateTopicId(db *sql.DB, id int) {
	_, err := db.Exec("UPDATE test SET TopicToOpen = ? WHERE id_TopicToOpen = 1", id)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateNotConnect(db *sql.DB, id int) {
	_, err := db.Exec("UPDATE test SET NotConnect = ? WHERE id_TopicToOpen = 1", id)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateIdConnect(db *sql.DB, id int) {
	_, err := db.Exec("UPDATE test SET IdConnect  = ? WHERE id_TopicToOpen = 1", id)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateLikesMinusOne(db *sql.DB, id int) {
	_, err := db.Exec("UPDATE messages SET Likes_message = Likes_message -1 WHERE Id_message  = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

func RemoveLikes(db *sql.DB, idMsg, id int) {
	_, err := db.Exec("DELETE FROM likeNumber WHERE Id_message_likeNumber = ? AND Id_utilisateur_likeNumber = ?", idMsg, id)
	if err != nil {
		log.Fatal(err)
	}
}

func NewUserLikeMessage(db *sql.DB, tochange int, id int) {
	_, err := db.Exec("INSERT INTO likeNumber (Id_message_likeNumber,Id_utilisateur_likeNumber) VALUES (?, ?)", tochange, id)
	if err != nil {
		log.Fatal(err)
	}
}

func InspectUser(db *sql.DB, id int) string {
	row, err := db.Query("SELECT Nom_utilisateur FROM utilisateurs WHERE Id_utilisateur = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	var nom_utilisateur string
	if row.Next() {
		err := row.Scan(&nom_utilisateur)
		if err != nil {
			return ""
		}
	}

	row.Scan(&nom_utilisateur)
	return nom_utilisateur

}

//test inspect db = utilisateurs

func QueryDbUtilisateur(db *sql.DB, p Utilisateur) []Utilisateur {

	tabs := []Utilisateur{}
	row, err := db.Query("SELECT * FROM utilisateurs")
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		var Id_utilisateur int
		var Nom_utilisateur string
		var Mdp_utilisateur string
		var Email_utilisateur string
		var Avatar_utilisateur string
		row.Scan(&Id_utilisateur, &Nom_utilisateur, &Mdp_utilisateur, &Email_utilisateur, &Avatar_utilisateur)
		// fmt.Println("id = " + strconv.Itoa(id) + " nomutilisateur = " + nom_utilisateur + " motdepasse = " + mot_de_passe + "email" + email + "avatar" + avatar)
		test1 := Utilisateur{
			Id_utilisateur:     Id_utilisateur,
			Nom_utilisateur:    Nom_utilisateur,
			Mdp_utilisateur:    Mdp_utilisateur,
			Email_utilisateur:  Email_utilisateur,
			Avatar_utilisateur: Avatar_utilisateur,
		}

		// fmt.Println(test1)
		tabs = append(tabs, test1)

	}
	return tabs
}

func QueryDbTopic(db *sql.DB, p Topic) []Topic {

	tabs := []Topic{}
	row, err := db.Query("SELECT * FROM topic")
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		var Id_utilisateur_topic int
		var Id_topic int
		var Titre_topic string
		var Date_topic string
		var Contenu_topic string
		var Likes_topic int
		var Multimedia_topic string
		var Imgtchek_topic int
		var MiniatureYt_topic string

		row.Scan(&Id_utilisateur_topic, &Id_topic, &Titre_topic, &Date_topic, &Contenu_topic, &Likes_topic, &Multimedia_topic, &Imgtchek_topic, &MiniatureYt_topic)
		// fmt.Println("id = " + strconv.Itoa(id) + " nomutilisateur = " + nom_utilisateur + " motdepasse = " + mot_de_passe + "email" + email + "avatar" + avatar)
		test2 := Topic{
			Id_utilisateur_topic: Id_utilisateur_topic,
			Id_topic:             Id_topic,
			Titre_topic:          Titre_topic,
			Date_topic:           Date_topic,
			Contenu_topic:        Contenu_topic,
			Likes_topic:          Likes_topic,
			Multimedia_topic:     Multimedia_topic,
			Imgtchek_topic:       Imgtchek_topic,
			MiniatureYt_topic:    MiniatureYt_topic,
		}

		// fmt.Println(test1)
		tabs = append(tabs, test2)

	}
	return tabs
}

func QueryDbMessage(db *sql.DB, p Message) []Message {
	tabs := []Message{}
	// ("SELECT utilisateurs.Nom_utilisateur, utilisateurs.Mdp_utilisateur, utilisateurs.Email_utilisateur, utilisateurs.Avatar_utilisateur, utilisateurs.Id_utilisateur, topic.Id_topic, topic.Titre_topic, topic.Date_topic, topic.Contenu_topic, topic.Likes_topic FROM utilisateurs INNER JOIN topic ON utilisateurs.Id_utilisateur=topic.Id_utilisateur")
	row, err := db.Query("SELECT * FROM messages")
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		var Id_utilisateur_msg int
		var Id_topic_msg int
		var Id_message int
		var Contenu_message string
		var Date_message string
		var Likes_message int

		row.Scan(&Id_utilisateur_msg, &Id_topic_msg, &Id_message, &Contenu_message, &Date_message, &Likes_message)

		test3 := Message{
			Id_utilisateur_msg: Id_utilisateur_msg,
			Id_topic_msg:       Id_topic_msg,
			Id_message:         Id_message,
			Contenu_message:    Contenu_message,
			Date_message:       Date_message,
			Likes_message:      Likes_message,
		}

		tabs = append(tabs, test3)
	}

	return tabs
}

func QueryDbLikeNumber(db *sql.DB, p LikeNumber) []LikeNumber {
	tabs := []LikeNumber{}
	// ("SELECT utilisateurs.Nom_utilisateur, utilisateurs.Mdp_utilisateur, utilisateurs.Email_utilisateur, utilisateurs.Avatar_utilisateur, utilisateurs.Id_utilisateur, topic.Id_topic, topic.Titre_topic, topic.Date_topic, topic.Contenu_topic, topic.Likes_topic FROM utilisateurs INNER JOIN topic ON utilisateurs.Id_utilisateur=topic.Id_utilisateur")
	row, err := db.Query("SELECT * FROM likeNumber")
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		var Id_message_likeNumber int
		var Id_utilisateur_likeNumber int

		row.Scan(&Id_message_likeNumber, &Id_utilisateur_likeNumber)

		test3 := LikeNumber{
			Id_message_likeNumber:     Id_message_likeNumber,
			Id_utilisateur_likeNumber: Id_utilisateur_likeNumber,
		}

		tabs = append(tabs, test3)
	}

	return tabs
}

func QueryDbTest(db *sql.DB, p Test) []Test {
	tabs := []Test{}
	// ("SELECT utilisateurs.Nom_utilisateur, utilisateurs.Mdp_utilisateur, utilisateurs.Email_utilisateur, utilisateurs.Avatar_utilisateur, utilisateurs.Id_utilisateur, topic.Id_topic, topic.Titre_topic, topic.Date_topic, topic.Contenu_topic, topic.Likes_topic FROM utilisateurs INNER JOIN topic ON utilisateurs.Id_utilisateur=topic.Id_utilisateur")
	row, err := db.Query("SELECT * FROM test")
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		var TopicToOpen int
		var Id_TopicToOpen int
		var NotConnect int
		var IdConnect int

		row.Scan(&TopicToOpen, &Id_TopicToOpen, &NotConnect, &IdConnect)

		test3 := Test{
			TopicToOpen:    TopicToOpen,
			Id_TopicToOpen: Id_TopicToOpen,
			NotConnect:     NotConnect,
			IdConnect:      IdConnect,
		}

		tabs = append(tabs, test3)
	}

	return tabs
}

func QueryDbSessions(db *sql.DB, p Sessions) []Sessions {
	tabs := []Sessions{}
	// ("SELECT utilisateurs.Nom_utilisateur, utilisateurs.Mdp_utilisateur, utilisateurs.Email_utilisateur, utilisateurs.Avatar_utilisateur, utilisateurs.Id_utilisateur, topic.Id_topic, topic.Titre_topic, topic.Date_topic, topic.Contenu_topic, topic.Likes_topic FROM utilisateurs INNER JOIN topic ON utilisateurs.Id_utilisateur=topic.Id_utilisateur")
	row, err := db.Query("SELECT * FROM sessions")
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		var UUID_session string
		var Id_utilisateur_session int

		row.Scan(&UUID_session, &Id_utilisateur_session)

		test3 := Sessions{
			UUID_session:           UUID_session,
			Id_utilisateur_session: Id_utilisateur_session,
		}

		tabs = append(tabs, test3)
	}

	return tabs
}

// Récupère l'id de l'utilisateur en fonction du nom d'utilisateur

func GetIdFromUsername(db *sql.DB, nomUtilisateur string) int {
	var id int
	err := db.QueryRow("SELECT Id_utilisateur FROM utilisateurs WHERE Nom_utilisateur = ?", nomUtilisateur).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

// Gère la session de l'utilisateur avec le cookie et l'UUID sur la database

func CreateSession(w http.ResponseWriter, db *sql.DB, nomUtilisateur string, motDePasse string) error {
	// Vérifier l'authentification de l'utilisateur
	if !CheckUser(db, nomUtilisateur, motDePasse) {
		return errors.New("Authentification invalide")
	}

	// Générer un nouvel UUID de session
	sessionUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}

	// Vérifie si l'UUID existe déjà dans la base de données
	if CheckIfUUIDExists(db, sessionUUID.String()) {
		return errors.New("UUID invalide")
	}

	//verifie si id exist

	// Récupérer l'ID utilisateur
	idUtilisateur := GetIdFromUsername(db, nomUtilisateur)
	if idUtilisateur == 0 {
		return errors.New("Utilisateur introuvable")
	}

	if CheckIfIDExists(db, idUtilisateur) {
		_, err := db.Exec("DELETE FROM sessions WHERE Id_utilisateur_session = ?", idUtilisateur)
		if err != nil {
			return err
		}
	}

	// Insérer le hachage de l'UUID de session dans la base de données
	_, err = db.Exec("INSERT INTO sessions (Id_utilisateur_session, UUID_session) VALUES (?, ?)", idUtilisateur, sessionUUID.String())
	if err != nil {
		return err
	}

	// Créer un cookie pour l'identifiant de session

	expiration := time.Now().Add(30 * time.Second) //100s

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sessionUUID.String(),
		Path:    "/",
		Expires: expiration,
	}

	// Ajouter le cookie à la réponse

	http.SetCookie(w, cookie)
	UpdateNotConnect(db, -1)

	return nil
}

// Vérifie si l'UUID existe dans la base de données

func CheckIfUUIDExists(db *sql.DB, UUID string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM sessions WHERE UUID_session = ?", UUID).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

func CheckIfIDExists(db *sql.DB, id int) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM sessions WHERE Id_utilisateur_session = ?", id).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

// Gère la déconnexion de la session de l'utilisateur

func DeleteSession(w http.ResponseWriter, db *sql.DB, UUID string) error {
	// Vérifier si l'UUID existe dans la base de données
	if !CheckIfUUIDExists(db, UUID) {
		return errors.New("UUID invalide")
	}

	// Supprimer l'UUID de la base de données
	_, err := db.Exec("DELETE FROM sessions WHERE UUID_session = ?", UUID)
	if err != nil {
		return err
	}

	// Créer un cookie pour l'identifiant de session
	// Créer un cookie vide pour supprimer le cookie existant
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	}

	// Ajouter le cookie à la réponse
	http.SetCookie(w, cookie)

	return nil
}

func Hash(db *sql.DB, idUtilisateur int) {

	sessionUUID, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO sessions (UUID_session, Id_utilisateur_session ) VALUES (?, ?)", sessionUUID.String(), idUtilisateur)
	if err != nil {
		log.Fatal(err)
	}
}
