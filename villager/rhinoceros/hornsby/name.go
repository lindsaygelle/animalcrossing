package hornsby

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Hornsby"
	nameCanadianFrench       string = "Cornio"
	nameDutch                string = "Hornsby"
	nameFrench               string = "Cornio"
	nameGerman               string = "Rüdiger"
	nameItalian              string = "Rino"
	nameJapanese             string = "みつお"
	nameLatinAmericanSpanish string = "Rino"
	nameKorean               string = "뿌람"
	nameRussian              string = "Хорнсби"
	nameSpanish              string = "Rino"
	nameSimplifiedChinese    string = "光雄"
	nameTraditionalChinese   string = "光雄"
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
