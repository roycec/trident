// Code generated by go-swagger; DO NOT EDIT.

package snaplock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnaplockRetentionOperationDeleteReader is a Reader for the SnaplockRetentionOperationDelete structure.
type SnaplockRetentionOperationDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockRetentionOperationDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockRetentionOperationDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockRetentionOperationDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockRetentionOperationDeleteOK creates a SnaplockRetentionOperationDeleteOK with default headers values
func NewSnaplockRetentionOperationDeleteOK() *SnaplockRetentionOperationDeleteOK {
	return &SnaplockRetentionOperationDeleteOK{}
}

/* SnaplockRetentionOperationDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockRetentionOperationDeleteOK struct {
}

func (o *SnaplockRetentionOperationDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/snaplock/event-retention/operations/{id}][%d] snaplockRetentionOperationDeleteOK ", 200)
}

func (o *SnaplockRetentionOperationDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSnaplockRetentionOperationDeleteDefault creates a SnaplockRetentionOperationDeleteDefault with default headers values
func NewSnaplockRetentionOperationDeleteDefault(code int) *SnaplockRetentionOperationDeleteDefault {
	return &SnaplockRetentionOperationDeleteDefault{
		_statusCode: code,
	}
}

/* SnaplockRetentionOperationDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response codes
| Error code  |  Description |
|-------------|--------------|
| 14090541    | A completed or failed operation cannot be aborted |

*/
type SnaplockRetentionOperationDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock retention operation delete default response
func (o *SnaplockRetentionOperationDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockRetentionOperationDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/snaplock/event-retention/operations/{id}][%d] snaplock_retention_operation_delete default  %+v", o._statusCode, o.Payload)
}
func (o *SnaplockRetentionOperationDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockRetentionOperationDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}