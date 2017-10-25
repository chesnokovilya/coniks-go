package coniksclient

import (
	"encoding/json"

	"github.com/coniks-sys/coniks-go/protocol"
)

// UnmarshalResponse decodes the given message into a protocol.Response
// according to the given request type t. The request types are integer
// constants defined in the protocol package.
func UnmarshalResponse(t int, msg []byte) *protocol.Response {
	type Response struct {
		Error             protocol.ErrorCode
		DirectoryResponse json.RawMessage
	}
	var resp Response
	if err := json.Unmarshal(msg, &resp); err != nil {
		return &protocol.Response{
			Error: protocol.ErrMalformedMessage,
		}
	}

	// DirectoryResponse is omitempty for the places
	// where Error is in Errors
	if resp.DirectoryResponse == nil {
		response := &protocol.Response{
			Error: resp.Error,
		}
		if err := response.Validate(); err != nil {
			return &protocol.Response{
				Error: protocol.ErrMalformedMessage,
			}
		}
		return response
	}

	switch t {
	case protocol.RegistrationType, protocol.KeyLookupInEpochType, protocol.MonitoringType:
		response := new(protocol.DirectoryProof)
		if err := json.Unmarshal(resp.DirectoryResponse, &response); err != nil {
			return &protocol.Response{
				Error: protocol.ErrMalformedMessage,
			}
		}
		return &protocol.Response{
			Error:             resp.Error,
			DirectoryResponse: response,
		}
	case protocol.STRType:
		response := new(protocol.STRHistoryRange)
		if err := json.Unmarshal(resp.DirectoryResponse, &response); err != nil {
			return &protocol.Response{
				Error: protocol.ErrMalformedMessage,
			}
		}
		return &protocol.Response{
			Error:             resp.Error,
			DirectoryResponse: response,
		}
	default:
		panic("Unknown request type")
	}
}

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
