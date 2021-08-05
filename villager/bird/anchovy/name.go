package anchovy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Anchovy"
	nameCanadianFrench       string = "Miguel"
	nameDutch                string = "Anchovy"
	nameFrench               string = "Miguel"
	nameGerman               string = "Armin"
	nameItalian              string = "Acciuga"
	nameJapanese             string = "アンチョビ"
	nameLatinAmericanSpanish string = "Boquerón"
	nameKorean               string = "안쵸비"
	nameRussian              string = "Анчоуи"
	nameSpanish              string = "Boquerón"
	nameSimplifiedChinese    string = "凤尾"
	nameTraditionalChinese   string = "鳳尾"
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
