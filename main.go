package main

import (
	"fmt"
	"io/ioutil"
	"log"

	toml "github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	// File naming: translations.ja.toml, active.ja.toml.
	// langFiles := []string{"en_US.yaml", "ja_JP.yaml"}
	langFiles := []string{"en_US.toml", "ja_JP.toml"}
	bundle, err := CreateLocalizerBundle(langFiles)
	if err != nil {
		log.Fatal(err)
	}
	engLocalizer := i18n.NewLocalizer(bundle, "ja_JP")
	msg, err := engLocalizer.Localize(
		&i18n.LocalizeConfig{
			MessageID: "hello_world",
		},
	)
	fmt.Println(msg, err)
	msg, err = engLocalizer.Localize(
		&i18n.LocalizeConfig{
			MessageID: "numbers",
			TemplateData: map[string]interface{}{
				"Count": 1,
			},
			PluralCount: 1, // 1 - Singular, 2 - Plural.
		},
	)
	fmt.Println(msg, err)
}

func CreateLocalizerBundle(langFiles []string) (*i18n.Bundle, error) {
	bundle := i18n.NewBundle(language.English)
	// bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	for _, file := range langFiles {
		translations, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		bundle.MustParseMessageFileBytes(translations, file)
	}
	return bundle, nil
}
