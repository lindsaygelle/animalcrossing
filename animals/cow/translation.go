package cow

import "github.com/lindsaygelle/animalcrossing/translations"

var (
	_ translations.Translation = (translation{})
)

type translation struct {
	chineseSimplified
	chineseTraditional
	dutch
	french
	german
	italian
	japanese
	korean
	russian
	spanish
}

func (t translation) ChineseSimplified() translations.Language {
	return t.chineseSimplified
}

func (t translation) ChineseTraditional() translations.Language {
	return t.chineseTraditional
}

func (t translation) Dutch() translations.Language {
	return t.dutch
}

func (t translation) French() translations.Language {
	return t.french
}

func (t translation) German() translations.Language {
	return t.german
}

func (t translation) Italian() translations.Language {
	return t.italian
}

func (t translation) Japanese() translations.Language {
	return t.japanese
}

func (t translation) Korean() translations.Language {
	return t.korean
}

func (t translation) Russian() translations.Language {
	return t.russian
}

func (t translation) Spanish() translations.Language {
	return t.spanish
}
