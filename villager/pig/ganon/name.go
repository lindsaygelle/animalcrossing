package ganon

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ganon"
	nameCanadianFrench       string = "Ganon"
	nameDutch                string = ""
	nameFrench               string = "Ganon"
	nameGerman               string = "Ganon"
	nameItalian              string = "Ganon"
	nameJapanese             string = "ガノン"
	nameLatinAmericanSpanish string = "Ganon"
	nameKorean               string = "가논"
	nameRussian              string = ""
	nameSpanish              string = "Ganon"
	nameSimplifiedChinese    string = ""
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
