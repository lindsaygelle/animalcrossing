package barold

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Barold"
	nameCanadianFrench       string = "Manu"
	nameDutch                string = "Barold"
	nameFrench               string = "Manu"
	nameGerman               string = "Hubert"
	nameItalian              string = "Placido"
	nameJapanese             string = "ニッシー"
	nameLatinAmericanSpanish string = "Plácido"
	nameKorean               string = "곰시"
	nameRussian              string = "Барольд"
	nameSpanish              string = "Plácido"
	nameSimplifiedChinese    string = "阿西"
	nameTraditionalChinese   string = "阿西"
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
