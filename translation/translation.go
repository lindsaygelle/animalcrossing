package translation

import "golang.org/x/text/language"

// Translation is a language translation.
type Translation interface {
	Language() language.Tag
	Value() string
}

// translation implements Translation.
type translation struct {
	language language.Tag
	value    string
}

func (v translation) Language() language.Tag {
	return v.language
}

func (v translation) Value() string {
	return v.value
}

// NewTranslation returns a new Translation.
func NewTranslation(language language.Tag, value string) Translation {
	return translation{
		language: language,
		value:    value}
}
