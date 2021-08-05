package kevin

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Kevin"
	nameCanadianFrench       string = "Jean-Bon"
	nameDutch                string = "Kevin"
	nameFrench               string = "Jean-Bon"
	nameGerman               string = "Kevin"
	nameItalian              string = "Kotekiño"
	nameJapanese             string = "イノッチ"
	nameLatinAmericanSpanish string = "Porcinio"
	nameKorean               string = "멧지"
	nameRussian              string = "Кевин"
	nameSpanish              string = "Porcinio"
	nameSimplifiedChinese    string = "伊比利"
	nameTraditionalChinese   string = "伊比利"
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
