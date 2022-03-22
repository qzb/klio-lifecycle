package placeholders

import (
	"fmt"
	"strings"
)

// DuplicatedPlaceholderError reports that some placeholder name is
// duplicated. Since placeholder names are case-insensitive actual names may be
// different.
type DuplicatedPlaceholderError struct {
	Name1 string
	Name2 string
}

func (e *DuplicatedPlaceholderError) Error() string {
	return fmt.Sprintf("duplicated placeholders: %q and %q", e.Name1, e.Name2)
}

// CyclicPlaceholderError reports that replacement values contain placeholders
// which form a cycle.
type CyclicPlaceholderError struct {
	Cycle []string
}

func (e *CyclicPlaceholderError) Error() string {
	return fmt.Sprintf("cyclic placeholders:\n  %s", strings.Join(e.Cycle, " -> "))
}

// MissingPlaceholderError reports that some placeholder marker cannot be
// replaced because it is using a name which is not defined in values.
type MissingPlaceholderError struct {
	MissingName string
	ValidNames  []string
}

func (e *MissingPlaceholderError) Error() string {
	return fmt.Sprintf("value for {{ %s }} placeholder is not specified, available placeholders:\n  %s", e.MissingName, strings.Join(e.ValidNames, ", "))
}

// InvalidPlaceholderNameError records value name which contains unallowed
// characters.
type InvalidPlaceholderNameError struct {
	Name string
}

func (e *InvalidPlaceholderNameError) Error() string {
	return fmt.Sprintf(`invalid placeholder name: %q, placeholder name segments cannot be empty and must contain only letters, digits and "_"`, e.Name)
}
