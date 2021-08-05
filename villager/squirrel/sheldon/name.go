package sheldon

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Sheldon"
	nameCanadianFrench       string = "Roy"
	nameDutch                string = "Sheldon"
	nameFrench               string = "Roy"
	nameGerman               string = "Steffen"
	nameItalian              string = "Frugolo"
	nameJapanese             string = "クリス"
	nameLatinAmericanSpanish string = "Matracas"
	nameKorean               string = "크리스"
	nameRussian              string = "Шелдон"
	nameSpanish              string = "Matracas"
	nameSimplifiedChinese    string = "克栗斯"
	nameTraditionalChinese   string = "克栗斯"
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
