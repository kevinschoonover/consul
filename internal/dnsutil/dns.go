// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package dnsutil

import (
	"errors"
	"regexp"
)

// MaxLabelLength is the maximum length for a name that can be used in DNS.
const MaxLabelLength = 63

// InvalidNameRe is a regex that matches characters which can not be included in
// a DNS name.
var InvalidNameRe = regexp.MustCompile(`[^A-Za-z0-9\\-]+`)

// matches valid DNS labels according to RFC 1123, should be at most 63
// characters according to the RFC
var validLabel = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?$`)

// IsValidLabel returns true if the string given is a valid DNS label (RFC 1123).
// Note: the only difference between RFC 1035 and RFC 1123 labels is that in
// RFC 1123 labels can begin with a number.
func IsValidLabel(name string) bool {
	return validLabel.MatchString(name)
}

// ValidateLabel is similar to IsValidLabel except it returns an error
// instead of false when name is not a valid DNS label. The error will contain
// reference to what constitutes a valid DNS label.
func ValidateLabel(name string) error {
	if !IsValidLabel(name) {
		return errors.New("a valid DNS label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character")
	}
	return nil
}
