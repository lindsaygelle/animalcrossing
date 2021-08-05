package jeremiah

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Jeremiah"
	nameCanadianFrench       string = "Jérémie"
	nameDutch                string = "Jeremiah"
	nameFrench               string = "Jérémie"
	nameGerman               string = "Jörg"
	nameItalian              string = "Geremia"
	nameJapanese             string = "クワトロ"
	nameLatinAmericanSpanish string = "Jeremías"
	nameKorean               string = "드리미"
	nameRussian              string = "Джеремия"
	nameSpanish              string = "Jeremías"
	nameSimplifiedChinese    string = "阿刻"
	nameTraditionalChinese   string = "阿刻"
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
