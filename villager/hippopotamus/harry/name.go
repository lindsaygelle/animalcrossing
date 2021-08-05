package harry

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Harry"
	nameCanadianFrench       string = "Bob"
	nameDutch                string = "Harry"
	nameFrench               string = "Bob"
	nameGerman               string = "Jürgen"
	nameItalian              string = "Ercolino"
	nameJapanese             string = "オリバー"
	nameLatinAmericanSpanish string = "Harpo"
	nameKorean               string = "올리버"
	nameRussian              string = "Гарри"
	nameSpanish              string = "Harpo"
	nameSimplifiedChinese    string = "欧世豪"
	nameTraditionalChinese   string = "歐世豪"
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
