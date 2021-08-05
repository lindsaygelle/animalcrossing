package marcel

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Marcel"
	nameCanadianFrench       string = "Ismaël"
	nameDutch                string = "Marcel"
	nameFrench               string = "Ismaël"
	nameGerman               string = "Ronaldo"
	nameItalian              string = "Giosuè"
	nameJapanese             string = "もんじゃ"
	nameLatinAmericanSpanish string = "Eto"
	nameKorean               string = "에드워드"
	nameRussian              string = "Марсель"
	nameSpanish              string = "Eto"
	nameSimplifiedChinese    string = "文字烧"
	nameTraditionalChinese   string = "文字燒"
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
