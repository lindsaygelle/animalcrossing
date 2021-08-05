package pashmina

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Pashmina"
	nameCanadianFrench       string = "Chavrina"
	nameDutch                string = "Pashmina"
	nameFrench               string = "Chavrina"
	nameGerman               string = "Pamela"
	nameItalian              string = "Pashmina"
	nameJapanese             string = "バーバラ"
	nameLatinAmericanSpanish string = "Carla"
	nameKorean               string = "바바라"
	nameRussian              string = "Пашмина"
	nameSpanish              string = "Carla"
	nameSimplifiedChinese    string = "芭芭拉"
	nameTraditionalChinese   string = "芭芭拉"
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
