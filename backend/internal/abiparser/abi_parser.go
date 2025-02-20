package abiparser

import (
	"fmt"
	"regexp"
	"strings"
)

// ABIElement represents a single element in a typical Ethereum ABI JSON.
type ABIElement struct {
	Anonymous bool       `json:"anonymous"`
	Inputs    []ABIInput `json:"inputs"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
}

// ABIInput represents an input parameter in an ABI element.
type ABIInput struct {
	Indexed bool   `json:"indexed"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

// ParseEventSignature parses a string like:
//   "event Transfer(address indexed from, address indexed to, uint256 value)"
// and returns a corresponding ABIElement.
//
// This is a naive parser that assumes the input is well-formed for an event.
func ParseEventSignature(signature string) (*ABIElement, error) {
	// Regex to capture:
	// 1) the word "event"
	// 2) the event name (Transfer)
	// 3) the parameter list (address indexed from, address indexed to, uint256 value)
	re := regexp.MustCompile(`^event\s+(\w+)\s*\(([^)]*)\)$`)
	matches := re.FindStringSubmatch(signature)
	if len(matches) != 3 {
		return nil, fmt.Errorf("invalid event signature: %s", signature)
	}

	eventName := matches[1]
	params := matches[2]

	// Split the params by comma
	parts := splitByCommaRespectingParentheses(params) // or just do strings.Split(params, ",") for naive approach

	var inputs []ABIInput
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		// Each part could look like:
		//   "address indexed from"
		//   "address from"
		//   "uint256 value"
		// We'll parse out whether "indexed" is present, the type, and the name.
		isIndexed := false
		if strings.Contains(part, " indexed ") {
			isIndexed = true
		}

		fields := strings.Fields(part)
		if len(fields) < 2 {
			return nil, fmt.Errorf("invalid parameter definition: %s", part)
		}

		// Potential patterns:
		//   [0] = address, [1] = indexed, [2] = from
		//   [0] = address, [1] = from
		//   [0] = uint256, [1] = value
		var typ, name string
		if isIndexed {
			// e.g. ["address", "indexed", "from"]
			if len(fields) < 3 {
				return nil, fmt.Errorf("invalid indexed parameter: %s", part)
			}
			typ = fields[0]
			name = fields[2]
		} else {
			// e.g. ["address", "from"] or ["uint256", "value"]
			typ = fields[0]
			name = fields[1]
		}

		inputs = append(inputs, ABIInput{
			Indexed: isIndexed,
			Name:    name,
			Type:    typ,
		})
	}

	abiElement := &ABIElement{
		Anonymous: false,
		Inputs:    inputs,
		Name:      eventName,
		Type:      "event",
	}
	return abiElement, nil
}

// splitByCommaRespectingParentheses is a helper if you have more complex type expressions
// like arrays e.g. "uint256[] memory" etc. For this simple case, you can just do strings.Split.
// We'll keep it simple here:
func splitByCommaRespectingParentheses(s string) []string {
	// A naive approach for demonstration; only safe for well-formed single-level param lists.
	// If you need to handle nested parentheses, you'll need a stack-based approach.
	return strings.Split(s, ",")
}
