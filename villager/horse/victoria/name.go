package victoria

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Victoria"
	nameCanadianFrench       string = "Victoria"
	nameDutch                string = "Victoria"
	nameFrench               string = "Victoria"
	nameGerman               string = "Emma"
	nameItalian              string = "Vittoria"
	nameJapanese             string = "セントアロー"
	nameLatinAmericanSpanish string = "Victoria"
	nameKorean               string = "센트엘로"
	nameRussian              string = "Виктория"
	nameSpanish              string = "Victoria"
	nameSimplifiedChinese    string = "圣萝"
	nameTraditionalChinese   string = "聖蘿"
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
