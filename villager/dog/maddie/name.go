package maddie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Maddie"
	nameCanadianFrench       string = "Olympe"
	nameDutch                string = "Maddie"
	nameFrench               string = "Olympe"
	nameGerman               string = "Agnes"
	nameItalian              string = "Cristina"
	nameJapanese             string = "マロン"
	nameLatinAmericanSpanish string = "Martina"
	nameKorean               string = "마롱"
	nameRussian              string = "Мэдди"
	nameSpanish              string = "Martina"
	nameSimplifiedChinese    string = "麻蓉"
	nameTraditionalChinese   string = "麻蓉"
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
