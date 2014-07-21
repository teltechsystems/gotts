package tts

import (
	"github.com/clbanning/mxj/x2j"
	"sync"
)

type Translator struct {
	defaultLanguageCode string
	languagesDictionary map[string]map[string]string
	mu                  *sync.Mutex
	xmlData             string
}

func (t *Translator) Get(languageCode, key string) string {
	t.mu.Lock()
	defer t.mu.Unlock()

	languageMap, exists := t.languagesDictionary[languageCode]
	if !exists {
		languageMap = t.languagesDictionary[t.defaultLanguageCode]
	}

	return languageMap[key]
}

func NewTranslator(xmlData, defaultLanguageCode string) (*Translator, error) {
	languagesDictionary, err := parseXml(xmlData)
	if err != nil {
		return nil, err
	}

	translator := &Translator{
		defaultLanguageCode: defaultLanguageCode,
		languagesDictionary: languagesDictionary,
		mu:                  &sync.Mutex{},
		xmlData:             xmlData,
	}

	return translator, nil
}

func parseXml(xmlData string) (map[string]map[string]string, error) {
	languages, err := x2j.XmlToMap([]byte(xmlData))
	if err != nil {
		return nil, err
	}

	languagesMap := languages["translations"].(map[string]interface{})

	languagesDictionary := make(map[string]map[string]string)

	for key, value := range languagesMap {
		stringsMap := ((value.(map[string]interface{}))["strings"]).(map[string]interface{})

		languageDictionary := make(map[string]string)
		for k, v := range stringsMap {
			languageDictionary[k] = v.(string)
		}

		languagesDictionary[key] = languageDictionary
	}

	return languagesDictionary, nil
}
