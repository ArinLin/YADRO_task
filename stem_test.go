package main

import (
	"testing"
)

func TestNormalizeSentence(t *testing.T) {
	stopWordsMap, err := createStopWordsMap(eng)
	if err != nil {
		t.Fatalf("create stop words map: %s", err)
	}

	for _, test := range []struct {
		Name           string
		InputSentence  string
		ResultSentence string
		Language       string
		WantErr        bool
		Err            string
	}{
		{
			Name:           "Successful stem sentence",
			InputSentence:  "follower brings bunch of questions",
			ResultSentence: "follow bring bunch question",
			Language:       eng,
		},
		{
			Name:          "Use unavaliable language for stemming",
			InputSentence: "follower brings bunch of questions",
			Language:      "polska",
			WantErr:       true,
			Err:           "Unknown language: polska",
		},
	} {
		t.Run(test.Name, func(t *testing.T) {

			res, err := normalizeSentence(test.InputSentence, test.Language, stopWordsMap)
			if err != nil {
				if !test.WantErr {
					t.Errorf("unexpected error: %s", err.Error())
				}
				if err.Error() != test.Err {
					t.Errorf("unexpected error. Expected %q but got %q", test.Err, err.Error())
				}
				return
			}
			if test.WantErr {
				t.Errorf("expected error but nothing got")
			}

			if res != test.ResultSentence {
				t.Errorf("wrong result. Expected %q but got %q", test.ResultSentence, res)
			}
		})
	}
}
