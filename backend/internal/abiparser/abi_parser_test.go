package abiparser_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/msquared-io/etherbase/backend/internal/abiparser"
)

func TestParseEventSignature(t *testing.T) {
	input := "event Transfer(address indexed from, address indexed to, uint256 value)"
	expectedJSON := `{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "from",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "to",
				"type": "address"
			},
			{
				"indexed": false,
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Transfer",
		"type": "event"
	}`

	eventABI, err := abiparser.ParseEventSignature(input)
	require.NoError(t, err)

	bytes, err := json.MarshalIndent(eventABI, "", "  ")
	require.NoError(t, err)

	require.JSONEq(t, expectedJSON, string(bytes))
}
