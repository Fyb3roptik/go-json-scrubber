package jsonscrubber_test

import (
	"strings"
	"testing"

	jsonscrubber "github.com/Fyb3roptik/go-json-scrubber"
)

type User struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Address   *Address `json:"address"`
}

type Address struct {
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}

var user *User

func init() {
	user = &User{
		FirstName: "Foo",
		LastName:  "Bar",
		Address: &Address{
			City:  "New York",
			State: "NY",
		},
	}
}

func TestAddOnly(t *testing.T) {
	scrubbedUser := jsonscrubber.AddOnly(user, "first_name", "address").(map[string]interface{})
	scrubbedUser["address"] = jsonscrubber.AddOnly(user.Address, "city")

	expectedScrubbed := [][]string{{"last_name"}, {"address", "address"}, {"address", "state"}, {"address", "zip"}}

	for _, path := range expectedScrubbed {
		if fieldExists(scrubbedUser, path) {
			t.Errorf("unexpected field found %q", strings.Join(path, "."))
		}
	}
}

func TestRemoveOnly(t *testing.T) {
	scrubbedUser := jsonscrubber.RemoveOnly(user, "first_name").(map[string]interface{})
	scrubbedUser["address"] = jsonscrubber.RemoveOnly(user.Address, "city")

	expectedScrubbed := [][]string{{"first_name"}, {"address", "city"}}

	for _, path := range expectedScrubbed {
		if fieldExists(scrubbedUser, path) {
			t.Errorf("unexpected field found %q", strings.Join(path, "."))
		}
	}
}

// test helper, depends on the full object tree being map[string]interface{}
func fieldExists(object map[string]interface{}, path []string) bool {
	key := path[0]
	value, ok := object[key]

	if !ok {
		return false
	}

	if len(path) == 1 {
		return true
	}

	return fieldExists(value.(map[string]interface{}), append([]string{}, path[1:]...))
}
