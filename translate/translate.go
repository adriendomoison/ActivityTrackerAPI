package translate

import (
	"github.com/nicksnyder/go-i18n/i18n"
	"github.com/adriendomoison/ActivityTrackerAPI/DTO"
	"log"
	"fmt"
)

var T i18n.TranslateFunc

func SetLang() (DTO.ReturnMsg) {
	var err error
	//os.Getenv("LANG")
	T, err = i18n.Tfunc("en-us")
	if err != nil {
		log.Println(err)
		return DTO.ReturnMsg{-1, T("errorLanguage")}
	}
	log.Println(T("greetingNewLanguage"))
	return DTO.ReturnMsg{0, T("greetingNewLanguage")}
}

func ChangeLang(cookieLang string, acceptLang string) (DTO.ReturnMsg) {
	var err error
	
	defaultLang := "en-us"
	T, err = i18n.Tfunc(cookieLang, acceptLang, defaultLang)
	if err != nil {
		log.Println(err)
		return DTO.ReturnMsg{-1, T("errorLanguage")}
	}
	fmt.Println(T("greetingNewLanguage"))
	return DTO.ReturnMsg{0, T("greetingNewLanguage")}
}