package words

import (
	"testing"
)

func TestNormalizeSentence(t *testing.T) {
	stopWordsMap, err := CreateStopWordsMap("../../stop_words_eng.txt")
	if err != nil {
		t.Fatalf("create stop words map: %s", err)
	}

	for _, test := range []struct {
		Name           string
		InputSentence  string
		ResultSentence []string
		WantErr        bool
		Err            string
	}{
		{
			Name:           "Successful stem sentence",
			InputSentence:  "follower brings bunch of questions",
			ResultSentence: []string{"follow", "bring", "bunch", "question"},
		},
	} {
		t.Run(test.Name, func(t *testing.T) {

			res, err := NormalizeSentence(test.InputSentence, stopWordsMap)
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
			if len(res) != len(test.ResultSentence) {
				t.Errorf("different lens")
			}
			for i := range res {
				if res[i] != test.ResultSentence[i] {
					t.Errorf("wrong result. Expected %q but got %q", test.ResultSentence[i], res[i])
				}
			}
		})
	}
}
