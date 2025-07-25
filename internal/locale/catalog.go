// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package locale // import "miniflux.app/v2/internal/locale"

import (
	"embed"
	"encoding/json"
	"fmt"
)

type translationDict map[string]any
type catalog map[string]translationDict

var defaultCatalog = make(catalog, len(AvailableLanguages))

//go:embed translations/*.json
var translationFiles embed.FS

func getTranslationDict(language string) (translationDict, error) {
	if _, ok := defaultCatalog[language]; !ok {
		var err error
		if defaultCatalog[language], err = loadTranslationFile(language); err != nil {
			return nil, err
		}
	}
	return defaultCatalog[language], nil
}

func loadTranslationFile(language string) (translationDict, error) {
	translationFileData, err := translationFiles.ReadFile("translations/" + language + ".json")
	if err != nil {
		return nil, err
	}

	translationMessages, err := parseTranslationMessages(translationFileData)
	if err != nil {
		return nil, err
	}

	return translationMessages, nil
}

func parseTranslationMessages(data []byte) (translationDict, error) {
	var translationMessages translationDict
	if err := json.Unmarshal(data, &translationMessages); err != nil {
		return nil, fmt.Errorf(`invalid translation file: %w`, err)
	}
	return translationMessages, nil
}
