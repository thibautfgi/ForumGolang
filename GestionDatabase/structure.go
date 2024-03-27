package GestionDatabase

type Utilisateur struct {
	Id_utilisateur     int
	Nom_utilisateur    string
	Mdp_utilisateur    string
	Email_utilisateur  string
	Avatar_utilisateur string
}

type Message struct {
	Id_utilisateur_msg int
	Id_topic_msg       int
	Id_message         int
	Contenu_message    string
	Date_message       string
	Likes_message      int
}

type Topic struct {
	Id_utilisateur_topic int
	Id_topic             int
	Titre_topic          string
	Date_topic           string
	Contenu_topic        string
	Likes_topic          int
	Multimedia_topic     string
	Imgtchek_topic       int
	MiniatureYt_topic    string
}

type LikeNumber struct {
	Id_message_likeNumber     int
	Id_utilisateur_likeNumber int
}

type Forum struct {
	Utilisateurs []Utilisateur //YA UN S ATTENTION
	Topics       []Topic
	Messages     []Message
	LikeNumbers  []LikeNumber
	Tests        []Test
	Sessions     []Sessions
}

type Test struct {
	TopicToOpen    int
	Id_TopicToOpen int
	NotConnect int
	IdConnect  int
}

type Sessions struct {
	UUID_session           string
	Id_utilisateur_session int
}
