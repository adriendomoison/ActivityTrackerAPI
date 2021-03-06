package main

import (
	"github.com/adriendomoison/ActivityTrackerAPI/repository"
	"github.com/adriendomoison/ActivityTrackerAPI/api"
	"github.com/nicksnyder/go-i18n/i18n"
	"github.com/adriendomoison/ActivityTrackerAPI/translate"
)

func main() {
	i18n.MustLoadTranslationFile("translation/en-us.all.json")
	i18n.MustLoadTranslationFile("translation/fr-fr.all.json")
	translate.SetLang()
	repository.InitDBConnection()
	api.InitApi()
}