package gojraphql

import (
	"encoding/json"
	"testing"
)

const QUERY = `{
    "users-list-1083746374": {
        "@query": "listUsers",
        "@args": {
            "offset": 100,
            "limit": 300
        }
    },
    "users-28394783423": {
        "@mutate": "saveUser",
        "@args": {
          "id": 9349384,
          "firstName": "John",
          "lastName": "Doe",
          "address": {
            "streetName": "21 Klaana St. Ako Adei Osu. Accra Ghana"
          }
        },
        "@fields": [
            "id",
            {
                "@name": "address",
                "@alias": "userHomeAddress",
                "@fields": ["id", "streetName"]
            },
            {"@name": "status", "@alias": "userActivityStatus"}
        ]
    }
}`

func TestQuery(t *testing.T) {
	r := make(map[string]*Query)
	err := json.Unmarshal([]byte(QUERY), &r)
	if err != nil {
		t.Error(err)
		t.Failed()
	}

	// add more test
}
