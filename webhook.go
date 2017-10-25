package telegraph

import "encoding/json"

func WebHookParseRequest(r []byte) (*Update, error) {
	update := &Update{}

	if err := json.Unmarshal(r, update); err != nil {
		return nil, err
	}

	return update, nil
}
