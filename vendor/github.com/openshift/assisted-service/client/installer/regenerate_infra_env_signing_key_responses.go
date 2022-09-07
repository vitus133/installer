// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// RegenerateInfraEnvSigningKeyReader is a Reader for the RegenerateInfraEnvSigningKey structure.
type RegenerateInfraEnvSigningKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegenerateInfraEnvSigningKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRegenerateInfraEnvSigningKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRegenerateInfraEnvSigningKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRegenerateInfraEnvSigningKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRegenerateInfraEnvSigningKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewRegenerateInfraEnvSigningKeyMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRegenerateInfraEnvSigningKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegenerateInfraEnvSigningKeyNoContent creates a RegenerateInfraEnvSigningKeyNoContent with default headers values
func NewRegenerateInfraEnvSigningKeyNoContent() *RegenerateInfraEnvSigningKeyNoContent {
	return &RegenerateInfraEnvSigningKeyNoContent{}
}

/* RegenerateInfraEnvSigningKeyNoContent describes a response with status code 204, with default header values.

Success.
*/
type RegenerateInfraEnvSigningKeyNoContent struct {
}

func (o *RegenerateInfraEnvSigningKeyNoContent) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/regenerate-signing-key][%d] regenerateInfraEnvSigningKeyNoContent ", 204)
}

func (o *RegenerateInfraEnvSigningKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegenerateInfraEnvSigningKeyUnauthorized creates a RegenerateInfraEnvSigningKeyUnauthorized with default headers values
func NewRegenerateInfraEnvSigningKeyUnauthorized() *RegenerateInfraEnvSigningKeyUnauthorized {
	return &RegenerateInfraEnvSigningKeyUnauthorized{}
}

/* RegenerateInfraEnvSigningKeyUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type RegenerateInfraEnvSigningKeyUnauthorized struct {
	Payload *models.InfraError
}

func (o *RegenerateInfraEnvSigningKeyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/regenerate-signing-key][%d] regenerateInfraEnvSigningKeyUnauthorized  %+v", 401, o.Payload)
}
func (o *RegenerateInfraEnvSigningKeyUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *RegenerateInfraEnvSigningKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateInfraEnvSigningKeyForbidden creates a RegenerateInfraEnvSigningKeyForbidden with default headers values
func NewRegenerateInfraEnvSigningKeyForbidden() *RegenerateInfraEnvSigningKeyForbidden {
	return &RegenerateInfraEnvSigningKeyForbidden{}
}

/* RegenerateInfraEnvSigningKeyForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type RegenerateInfraEnvSigningKeyForbidden struct {
	Payload *models.InfraError
}

func (o *RegenerateInfraEnvSigningKeyForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/regenerate-signing-key][%d] regenerateInfraEnvSigningKeyForbidden  %+v", 403, o.Payload)
}
func (o *RegenerateInfraEnvSigningKeyForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *RegenerateInfraEnvSigningKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateInfraEnvSigningKeyNotFound creates a RegenerateInfraEnvSigningKeyNotFound with default headers values
func NewRegenerateInfraEnvSigningKeyNotFound() *RegenerateInfraEnvSigningKeyNotFound {
	return &RegenerateInfraEnvSigningKeyNotFound{}
}

/* RegenerateInfraEnvSigningKeyNotFound describes a response with status code 404, with default header values.

Error.
*/
type RegenerateInfraEnvSigningKeyNotFound struct {
	Payload *models.Error
}

func (o *RegenerateInfraEnvSigningKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/regenerate-signing-key][%d] regenerateInfraEnvSigningKeyNotFound  %+v", 404, o.Payload)
}
func (o *RegenerateInfraEnvSigningKeyNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegenerateInfraEnvSigningKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateInfraEnvSigningKeyMethodNotAllowed creates a RegenerateInfraEnvSigningKeyMethodNotAllowed with default headers values
func NewRegenerateInfraEnvSigningKeyMethodNotAllowed() *RegenerateInfraEnvSigningKeyMethodNotAllowed {
	return &RegenerateInfraEnvSigningKeyMethodNotAllowed{}
}

/* RegenerateInfraEnvSigningKeyMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type RegenerateInfraEnvSigningKeyMethodNotAllowed struct {
	Payload *models.Error
}

func (o *RegenerateInfraEnvSigningKeyMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/regenerate-signing-key][%d] regenerateInfraEnvSigningKeyMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *RegenerateInfraEnvSigningKeyMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegenerateInfraEnvSigningKeyMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateInfraEnvSigningKeyInternalServerError creates a RegenerateInfraEnvSigningKeyInternalServerError with default headers values
func NewRegenerateInfraEnvSigningKeyInternalServerError() *RegenerateInfraEnvSigningKeyInternalServerError {
	return &RegenerateInfraEnvSigningKeyInternalServerError{}
}

/* RegenerateInfraEnvSigningKeyInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type RegenerateInfraEnvSigningKeyInternalServerError struct {
	Payload *models.Error
}

func (o *RegenerateInfraEnvSigningKeyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/regenerate-signing-key][%d] regenerateInfraEnvSigningKeyInternalServerError  %+v", 500, o.Payload)
}
func (o *RegenerateInfraEnvSigningKeyInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegenerateInfraEnvSigningKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
