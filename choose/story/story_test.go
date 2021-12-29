package story

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestStoryLoad(t *testing.T) {

	jsn := `
	{
		"intro": {
			"title": "title",
			"story": ["aa", "bb"],
			"options" :
			[
				{"text": "opt 1", "arc": "arc_1"},
				{"text": "opt 2", "arc": "arc_2"}
			]
		}
	}`

	got, err := load([]byte(jsn))
	if err != nil {
		t.Error(err)
	}

	want := Story{
		"intro": {
			Title: "title",
			Story: []string{"aa", "bb"},
			Options: []Option{
				{Text: "opt 1", Arc: "arc_1"},
				{Text: "opt 2", Arc: "arc_2"},
			},
		},
	}

	if !reflect.DeepEqual(want, got) {
		w, _ := json.Marshal(want)
		g, _ := json.Marshal(got)
		t.Errorf("Want: %s", w)
		t.Errorf("Got:  %s", g)
	}

}
