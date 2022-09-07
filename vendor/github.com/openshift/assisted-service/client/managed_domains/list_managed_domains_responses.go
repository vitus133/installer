// Code generated by go-swagger; DO NOT EDIT.

package managed_domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// ListManagedDomainsReader is a Reader for the ListManagedDomains structure.
type ListManagedDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListManagedDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListManagedDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListManagedDomainsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListManagedDomainsOK creates a ListManagedDomainsOK with default headers values
func NewListManagedDomainsOK() *ListManagedDomainsOK {
	return &ListManagedDomainsOK{}
}

/* ListManagedDomainsOK describes a response with status code 200, with default header values.

Success.
*/
type ListManagedDomainsOK struct {
	Payload models.ListManagedDomains
}

func (o *ListManagedDomainsOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains][%d] listManagedDomainsOK  %+v", 200, o.Payload)
}
func (o *ListManagedDomainsOK) GetPayload() models.ListManagedDomains {
	return o.Payload
}

func (o *ListManagedDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListManagedDomainsInternalServerError creates a ListManagedDomainsInternalServerError with default headers values
func NewListManagedDomainsInternalServerError() *ListManagedDomainsInternalServerError {
	return &ListManagedDomainsInternalServerError{}
}

/* ListManagedDomainsInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type ListManagedDomainsInternalServerError struct {
	Payload *models.Error
}

func (o *ListManagedDomainsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains][%d] listManagedDomainsInternalServerError  %+v", 500, o.Payload)
}
func (o *ListManagedDomainsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListManagedDomainsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
