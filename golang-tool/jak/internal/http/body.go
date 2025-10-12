package http

import (
	"encoding/json"

	se "github.com/ymatsukawa/jak/internal/sys_error"
)

// RequestBody defines an interface for handling different types of HTTP request bodies.
// Implementations of this interface represent different content types and formats.
type RequestBody interface {
	// ContentType returns the MIME type of the body.
	ContentType() string

	// Content returns the body content as a string.
	Content() string

	// Validate checks if the body is valid according to its content type rules.
	Validate() error

	// IsEmpty checks if the body contains any content.
	IsEmpty() bool
}

// BaseBody provides common functionality for all request body types.
// It serves as a base implementation that specific body types can embed.
type BaseBody struct {
	// BodyContent contains the actual content of the request body as a string.
	BodyContent string
}

// NewBaseBody creates a new BaseBody with the specified content.
//
// Parameters:
//   - content: String content for the body
//
// Returns:
//   - BaseBody: Initialized base body structure
func NewBaseBody(content string) BaseBody {
	return BaseBody{
		BodyContent: content,
	}
}

// Content returns the body content as a string.
//
// Returns:
//   - string: The body content
func (b *BaseBody) Content() string {
	return b.BodyContent
}

// IsEmpty checks if the body is empty (contains no content).
//
// Returns:
//   - bool: True if the body is empty, false otherwise
func (b *BaseBody) IsEmpty() bool {
	return b.BodyContent == ""
}

// validateBase performs basic validation common to all body types.
// It checks if the body content is not empty.
//
// Returns:
//   - error: se.ErrBodyEmpty if the body is empty, nil otherwise
func (b *BaseBody) validateBase() error {
	if b.BodyContent == "" {
		return se.ErrBodyEmpty
	}
	return nil
}

// RawBody represents a raw request body with a custom content type.
// It can be used for any content type not covered by more specific implementations.
type RawBody struct {
	BaseBody
	// Type specifies the content type (MIME type) of the raw body.
	Type string
}

// NewRawBody creates a new raw body with the specified content and content type.
//
// Parameters:
//   - content: String content for the body
//   - contentType: MIME type for the content
//
// Returns:
//   - *RawBody: Initialized raw body structure
func NewRawBody(content, contentType string) *RawBody {
	return &RawBody{
		BaseBody: NewBaseBody(content),
		Type:     contentType,
	}
}

// ContentType returns the MIME type of the raw body.
//
// Returns:
//   - string: The content type
func (b *RawBody) ContentType() string {
	return b.Type
}

// IsEmpty checks if the raw body is empty.
//
// Returns:
//   - bool: True if the body is empty, false otherwise
func (b *RawBody) IsEmpty() bool {
	return b.BaseBody.IsEmpty()
}

// Validate checks if the raw body is valid.
// A valid raw body has non-empty content and a specified content type.
//
// Returns:
//   - error: se.Error if validation fails, nil otherwise
func (b *RawBody) Validate() error {
	if err := b.validateBase(); err != nil {
		return err
	}
	if b.Type == "" {
		return se.ErrContentTypeEmpty
	}

	return nil
}

// FormBody represents a form-urlencoded request body.
// It is used for sending form data in the format "key1=value1&key2=value2".
type FormBody struct {
	BaseBody
}

// NewFormBody creates a new form-urlencoded body with the specified content.
//
// Parameters:
//   - content: Form content in the format "key1=value1&key2=value2"
//
// Returns:
//   - *FormBody: Initialized form body structure
func NewFormBody(content string) *FormBody {
	return &FormBody{
		BaseBody: NewBaseBody(content),
	}
}

// ContentType returns the MIME type for form-urlencoded content.
//
// Returns:
//   - string: "application/x-www-form-urlencoded"
func (b *FormBody) ContentType() string {
	return ContentTypeFormURLEncoded
}

// Validate checks if the form body is valid.
// A valid form body has non-empty content.
//
// Returns:
//   - error: se.Error if validation fails, nil otherwise
func (b *FormBody) Validate() error {
	return b.validateBase()
}

// JsonBody represents a JSON request body.
// It is used for sending data in JSON format.
type JsonBody struct {
	BaseBody
}

// NewJsonBody creates a new JSON body with the specified content.
//
// Parameters:
//   - content: JSON string content
//
// Returns:
//   - *JsonBody: Initialized JSON body structure
func NewJsonBody(content string) *JsonBody {
	return &JsonBody{
		BaseBody: NewBaseBody(content),
	}
}

// ContentType returns the MIME type for JSON content.
//
// Returns:
//   - string: "application/json"
func (b *JsonBody) ContentType() string {
	return ContentTypeJSON
}

// Validate checks if the JSON body is valid.
// A valid JSON body has non-empty content and is parseable as JSON.
//
// Returns:
//   - error: se.Error if validation fails, nil otherwise
func (b *JsonBody) Validate() error {
	if err := b.validateBase(); err != nil {
		return err
	}
	if !json.Valid([]byte(b.BodyContent)) {
		return se.ErrInvalidJSONFormat
	}

	return nil
}
