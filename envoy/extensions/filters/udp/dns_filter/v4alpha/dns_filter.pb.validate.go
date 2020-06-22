// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/udp/dns_filter/v4alpha/dns_filter.proto

package envoy_extensions_filters_udp_dns_filter_v4alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _dns_filter_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on DnsFilterConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DnsFilterConfig) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		return DnsFilterConfigValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetServerConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsFilterConfigValidationError{
				field:  "ServerConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetClientConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsFilterConfigValidationError{
				field:  "ClientConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DnsFilterConfigValidationError is the validation error returned by
// DnsFilterConfig.Validate if the designated constraints aren't met.
type DnsFilterConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsFilterConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsFilterConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsFilterConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsFilterConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsFilterConfigValidationError) ErrorName() string { return "DnsFilterConfigValidationError" }

// Error satisfies the builtin error interface
func (e DnsFilterConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsFilterConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsFilterConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsFilterConfigValidationError{}

// Validate checks the field values on DnsFilterConfig_ServerContextConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *DnsFilterConfig_ServerContextConfig) Validate() error {
	if m == nil {
		return nil
	}

	switch m.ConfigSource.(type) {

	case *DnsFilterConfig_ServerContextConfig_InlineDnsTable:

		if v, ok := interface{}(m.GetInlineDnsTable()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsFilterConfig_ServerContextConfigValidationError{
					field:  "InlineDnsTable",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DnsFilterConfig_ServerContextConfig_ExternalDnsTable:

		if v, ok := interface{}(m.GetExternalDnsTable()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsFilterConfig_ServerContextConfigValidationError{
					field:  "ExternalDnsTable",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return DnsFilterConfig_ServerContextConfigValidationError{
			field:  "ConfigSource",
			reason: "value is required",
		}

	}

	return nil
}

// DnsFilterConfig_ServerContextConfigValidationError is the validation error
// returned by DnsFilterConfig_ServerContextConfig.Validate if the designated
// constraints aren't met.
type DnsFilterConfig_ServerContextConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsFilterConfig_ServerContextConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsFilterConfig_ServerContextConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsFilterConfig_ServerContextConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsFilterConfig_ServerContextConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsFilterConfig_ServerContextConfigValidationError) ErrorName() string {
	return "DnsFilterConfig_ServerContextConfigValidationError"
}

// Error satisfies the builtin error interface
func (e DnsFilterConfig_ServerContextConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsFilterConfig_ServerContextConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsFilterConfig_ServerContextConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsFilterConfig_ServerContextConfigValidationError{}

// Validate checks the field values on DnsFilterConfig_ClientContextConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *DnsFilterConfig_ClientContextConfig) Validate() error {
	if m == nil {
		return nil
	}

	if d := m.GetResolverTimeout(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return DnsFilterConfig_ClientContextConfigValidationError{
				field:  "ResolverTimeout",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gte := time.Duration(1*time.Second + 0*time.Nanosecond)

		if dur < gte {
			return DnsFilterConfig_ClientContextConfigValidationError{
				field:  "ResolverTimeout",
				reason: "value must be greater than or equal to 1s",
			}
		}

	}

	if len(m.GetUpstreamResolvers()) < 1 {
		return DnsFilterConfig_ClientContextConfigValidationError{
			field:  "UpstreamResolvers",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetUpstreamResolvers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsFilterConfig_ClientContextConfigValidationError{
					field:  fmt.Sprintf("UpstreamResolvers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetMaxPendingLookups() < 1 {
		return DnsFilterConfig_ClientContextConfigValidationError{
			field:  "MaxPendingLookups",
			reason: "value must be greater than or equal to 1",
		}
	}

	return nil
}

// DnsFilterConfig_ClientContextConfigValidationError is the validation error
// returned by DnsFilterConfig_ClientContextConfig.Validate if the designated
// constraints aren't met.
type DnsFilterConfig_ClientContextConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsFilterConfig_ClientContextConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsFilterConfig_ClientContextConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsFilterConfig_ClientContextConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsFilterConfig_ClientContextConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsFilterConfig_ClientContextConfigValidationError) ErrorName() string {
	return "DnsFilterConfig_ClientContextConfigValidationError"
}

// Error satisfies the builtin error interface
func (e DnsFilterConfig_ClientContextConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsFilterConfig_ClientContextConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsFilterConfig_ClientContextConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsFilterConfig_ClientContextConfigValidationError{}
