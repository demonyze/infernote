package fetcher

import (
	"errors"
	"fmt"

	"github.com/demonyze/infernote/internal/model"
)

func GetLanguage(lang string, url string) (model.LanguageImport, error) {
	var supportedLanguages = [1]string{"en"}

	var isSupportedLanguage bool
	for _, supported := range supportedLanguages {
		if lang == supported {
			isSupportedLanguage = true
			break
		}
	}

	if !isSupportedLanguage {
		message := fmt.Sprintf("%s is not a supported language.", lang)
		return model.LanguageImport{}, errors.New(message)
	}

	requestUrl := fmt.Sprintf("%s/%s.json", url, lang)
	languageImport, err := GetJson[model.LanguageImport](requestUrl)
	if err != nil {
		return model.LanguageImport{}, err
	}

	return languageImport, nil
}
