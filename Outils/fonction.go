package Outils

import (
	"fmt"
	"forum/GestionDatabase"
	"strconv"
	"strings"
	"time"
)

func ParseUserIdAndLikeNumber(userNlike string) (int, int) {
	split := strings.Split(userNlike, "+")
	tabs1 := split[0]

	tabs2 := split[1]

	tabs1final, _ := strconv.Atoi(tabs1)
	tabs2final, _ := strconv.Atoi(tabs2)

	return tabs1final, tabs2final
}

func TestMatchMsg(mytestuwu []GestionDatabase.LikeNumber, creatstruct GestionDatabase.LikeNumber) bool {
	for i := range mytestuwu {
		count := 0
		if creatstruct == mytestuwu[i] {
			count++
		} else {

		}
		if count > 0 {
			return true
		}

	}
	return false
}

func MultimediaConverteur(content string) string { // permet d'envoye des liens yt
	var final string

	if TestYtLink(content) == 1 {

		if Index(content, "youtube") == 1 {
			split := strings.Split(content, "?v=")
			split2 := strings.Split(split[1], "&t=")
			final = "https://www.youtube.com/embed/" + split2[0] + "?&autoplay=1&mute=1"

		} else {
			final = content
		}
	} else {
		if Index(content, "youtube") == 1 {
			split := strings.Split(content, "?v=")
			split2 := strings.Split(split[1], "&ab")
			final = "https://www.youtube.com/embed/" + split2[0] + "?&autoplay=1&mute=1"

		} else {
			final = content
		}
	}

	fmt.Println(final)
	return final
}

func Index(s string, toFind string) int { //j'ai repris index de la piscine
	for i := 0; i < len(s); i++ {
		if i+len(toFind) <= len(s) {
			slicetest := s[i : i+len(toFind)]
			if slicetest == toFind {
				return 1
			}
		}
	}
	return -1
}

func TchekImage(content string) int { // permet de differencier video et img
	if Index(content, "youtube") == 1 {
		return 1
	} else {
		return -1
	}

}

func TestYtLink(content string) int {
	if Index(content, "&t=") == 1 {
		return 1
	} else {
		return -1
	}
}

func GetIdImageFromYt(content string) string {

	if TestYtLink(content) == 1 {
		var final string
		split := strings.Split(content, "?v=")
		split2 := strings.Split(split[1], "&t=")
		final = "https://img.youtube.com/vi/" + split2[0] + "/maxresdefault.jpg"
		return final
	} else {
		var final string
		split := strings.Split(content, "?v=")
		split2 := strings.Split(split[1], "&ab")
		final = "https://img.youtube.com/vi/" + split2[0] + "/maxresdefault.jpg"
		return final
	}

}

func Date() string {
	date := time.Now()
	annee, mois, jour := date.Date()
	heure, minutes := date.Hour(), date.Minute()
	dateStr := fmt.Sprintf("%d-%02d-%02d", jour, mois, annee)
	time := strconv.Itoa(heure) + ":" + strconv.Itoa(minutes)
	AllDate := dateStr + " " + time
	fmt.Print(AllDate)
	return AllDate
}
