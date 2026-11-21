package logtext_test

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"testing"
	"time"

	"github.com/Boalq/CLI-Obsidian-Logger/internal/logtext"
)

func TestLogtxtStruct(t *testing.T) {
	logtxt := logtext.LogText{"", "", ""}
	LogTime := fmt.Sprintf("%d:%d", time.Now().Hour(), time.Now().Minute())

	logtxt.Header = fmt.Sprintf("\n--- \t New Log at: %s \t ---\n\n", LogTime)
	logtxt.Body = "hello today live at 9\n I shitted in my pants"
	logtxt.Footer = "\n..."

	msg := logtxt.Header + logtxt.Body + logtxt.Footer

	t.Log(msg)
}

func TestJSONMarshal(t *testing.T) {
	logtxt := logtext.LogText{"", "", ""}
	LogTime := fmt.Sprintf("%d:%d", time.Now().Hour(), time.Now().Minute())

	logtxt.Header = fmt.Sprintf("\n--- \t New Log at: %s \t ---\n\n", LogTime)
	logtxt.Body = "hello today live at 9\n I shitted in my pants"
	logtxt.Footer = "\n..."

	// MarshalingToJSON
	// Struct -> JSON
	w, err := json.Marshal(&logtxt)
	if err != nil {
		t.Error(err)
	}

	result := jsontext.Value(w)
	t.Log(result)
}

func TestJSONUnmarshal(t *testing.T) {
	w, err := json.Marshal(&logtext.LogText{"\n--Header--\n", "hello", "\n"})

	var new logtext.LogText

	// JSON -> Struct
	err = json.Unmarshal(w, &new)
	if err != nil {
		t.Error(err)
	}

	t.Log(new)
}
