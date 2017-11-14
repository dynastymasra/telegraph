package telegraph

import "encoding/json"

// WebHookParseRequest function for parse request from telegram web hook, return struct Update if success
func WebHookParseRequest(r []byte) (*Update, error) {
	update := &Update{}

	if err := json.Unmarshal(r, update); err != nil {
		return nil, err
	}

	return update, nil
}
