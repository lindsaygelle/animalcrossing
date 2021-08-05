package belle

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Belle"
	nameCanadianFrench       string = "meuh-euh"
	nameDutch                string = ""
	nameFrench               string = "meuh-euh"
	nameGerman               string = "knutsch"
	nameItalian              string = "amuuur"
	nameJapanese             string = "にゅう"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "cencerrito"
	nameSimplifiedChinese    string = "Unknown"
	nameTraditionalChinese   string = ""
)

var (
	name = map[language.Tag]string{
		language.AmericanEnglish:      nameAmericanEnglish,
		language.CanadianFrench:       nameCanadianFrench,
		language.Dutch:                nameDutch,
		language.French:               nameFrench,
		language.German:               nameGerman,
		language.Italian:              nameItalian,
		language.Japanese:             nameJapanese,
		language.Korean:               nameKorean,
		language.LatinAmericanSpanish: nameLatinAmericanSpanish,
		language.Russian:              nameRussian,
		language.Spanish:              nameSpanish,
		language.SimplifiedChinese:    nameSimplifiedChinese,
		language.TraditionalChinese:   nameTraditionalChinese}
)
