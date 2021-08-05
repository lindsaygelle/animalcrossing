package leopold

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Leopold"
	nameCanadianFrench       string = "Leandro"
	nameDutch                string = "Leopold"
	nameFrench               string = "Leandro"
	nameGerman               string = "Leandro"
	nameItalian              string = "Leandro"
	nameJapanese             string = "ティーチャー"
	nameLatinAmericanSpanish string = "Leopoldo"
	nameKorean               string = "티처"
	nameRussian              string = "Леопольд"
	nameSpanish              string = "Leopoldo"
	nameSimplifiedChinese    string = "老狮"
	nameTraditionalChinese   string = "老獅"
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
