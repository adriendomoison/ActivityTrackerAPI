package translate

import (
	"github.com/nicksnyder/go-i18n/i18n"
	"log"
	"fmt"
	"os"
	"github.com/adriendomoison/ActivityTrackerAPI/utils"
)

var T i18n.TranslateFunc

func SetLang() (utils.ReturnMsg) {
	var err error
	
	T, err = i18n.Tfunc(os.Getenv("LANG"))
	if err != nil {
		log.Println(err)
		return utils.ReturnMsg{-1, T("errorLanguage")}
	}
	log.Println(T("greetingNewLanguage"))
	return utils.ReturnMsg{0, T("greetingNewLanguage")}
}

func ChangeLang(cookieLang string, acceptLang string) (utils.ReturnMsg) {
	var err error
	
	defaultLang := "en-us"
	T, err = i18n.Tfunc(cookieLang, acceptLang, defaultLang)
	if err != nil {
		log.Println(err)
		return utils.ReturnMsg{-1, T("errorLanguage")}
	}
	fmt.Println(T("greetingNewLanguage"))
	return utils.ReturnMsg{0, T("greetingNewLanguage")}
}