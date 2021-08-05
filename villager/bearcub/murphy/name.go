package murphy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Murphy"
	nameCanadianFrench       string = "Eddie"
	nameDutch                string = "Murphy"
	nameFrench               string = "Eddie"
	nameGerman               string = "Michael"
	nameItalian              string = "Marcello"
	nameJapanese             string = "のりお"
	nameLatinAmericanSpanish string = "Orsón"
	nameKorean               string = "머피"
	nameRussian              string = "Мерфи"
	nameSpanish              string = "Orsón"
	nameSimplifiedChinese    string = "宪雄"
	nameTraditionalChinese   string = "憲雄"
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
