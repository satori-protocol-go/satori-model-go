package message

import (
	"testing"
)

func Test(t *testing.T) {
	for group, messages := range _getRawMessage() {
		t.Logf("start test group: %s", group)
		for _, message := range messages {
			_, err := Parse(message.Raw)
			if err != nil {
				t.Fatalf("%s Parse error: %s", message.Raw, err)
			}
			result, err := Stringify(message.Elements)
			if err != nil {
				t.Fatalf("%s Stringify error: %s", message.Elements, err)
			}
			if message.TargetRaw != "" {
				if result != message.TargetRaw {
					t.Fatalf("%s not eq %s", result, message.TargetRaw)
				}
			} else if result != message.Raw {
				t.Fatalf("%s not eq %s", result, message.Raw)
			}
		}
	}
}
