// Defines methods/functions to encode/decode messages between client
// and server. Currently this module supports JSON marshal/unmarshal only.
// Protobuf will be supported in the future.

package coniksserver

import (
	"encoding/json"

	"github.com/coniks-sys/coniks-go/protocol"
)

// UnmarshalRequest parses a JSON-encoded request msg and
// creates the corresponding protocol.Request, which will be handled
// by the server.
func UnmarshalRequest(msg []byte) (*protocol.Request, error) {
	var content json.RawMessage
	req := protocol.Request{
		Request: &content,
	}
	if err := json.Unmarshal(msg, &req); err != nil {
		return nil, err
	}
	var request interface{}
	switch req.Type {
	case protocol.RegistrationType:
		request = new(protocol.RegistrationRequest)
	case protocol.KeyLookupType:
		request = new(protocol.KeyLookupRequest)
	case protocol.KeyLookupInEpochType:
		request = new(protocol.KeyLookupInEpochRequest)
	case protocol.MonitoringType:
		request = new(protocol.MonitoringRequest)
	case protocol.STRType:
		request = new(protocol.STRHistoryRequest)
	}
	if err := json.Unmarshal(content, &request); err != nil {
		return nil, err
	}
	req.Request = request
	return &req, nil
}
