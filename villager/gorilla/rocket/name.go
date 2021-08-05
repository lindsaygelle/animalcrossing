package rocket

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rocket"
	nameCanadianFrench       string = "Gertrude"
	nameDutch                string = "Rocket"
	nameFrench               string = "Gertrude"
	nameGerman               string = "Katrin"
	nameItalian              string = "Kinga"
	nameJapanese             string = "４ごう"
	nameLatinAmericanSpanish string = "Gloria"
	nameKorean               string = "4호"
	nameRussian              string = "Рокет"
	nameSpanish              string = "Gloria"
	nameSimplifiedChinese    string = "阿四"
	nameTraditionalChinese   string = "阿四"
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
