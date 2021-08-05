package chuck

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Chuck"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "dugong"
	nameGerman               string = "daddel"
	nameItalian              string = "mummamia"
	nameJapanese             string = "ってんだ"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "cobarde"
	nameSimplifiedChinese    string = "伙计"
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
