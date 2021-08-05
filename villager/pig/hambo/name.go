package hambo

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Hambo"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "pti boudin"
	nameGerman               string = "hajaaa"
	nameItalian              string = "yo"
	nameJapanese             string = "だス"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "morcillita"
	nameSimplifiedChinese    string = "哈罗"
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
