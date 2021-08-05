package cousteau

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cousteau"
	nameCanadianFrench       string = "Figaro"
	nameDutch                string = "Cousteau"
	nameFrench               string = "Figaro"
	nameGerman               string = "Jacques"
	nameItalian              string = "Jacques"
	nameJapanese             string = "ハルマキ"
	nameLatinAmericanSpanish string = "Cousteau"
	nameKorean               string = "왕서방"
	nameRussian              string = "Кусто"
	nameSpanish              string = "Cousteau"
	nameSimplifiedChinese    string = "春卷"
	nameTraditionalChinese   string = "春卷"
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
