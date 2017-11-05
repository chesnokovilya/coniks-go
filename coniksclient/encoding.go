package coniksclient

import (
	"encoding/json"

	"github.com/coniks-sys/coniks-go/protocol"
)

// CreateRegistrationMsg returns a JSON encoding of
// a protocol.RegistrationRequest for the given (name, key) pair.
func CreateRegistrationMsg(name string, key []byte) ([]byte, error) {
	return json.Marshal(&protocol.Request{
		Type: protocol.RegistrationType,
		Request: &protocol.RegistrationRequest{
			Username: name,
			Key:      key,
		},
	})
}

// CreateKeyLookupMsg returns a JSON encoding of
// a protocol.KeyLookupRequest for the given name.
func CreateKeyLookupMsg(name string) ([]byte, error) {
	return json.Marshal(&protocol.Request{
		Type: protocol.KeyLookupType,
		Request: &protocol.KeyLookupRequest{
			Username: name,
		},
	})
}
