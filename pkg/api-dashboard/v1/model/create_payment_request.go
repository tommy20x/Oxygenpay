// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreatePaymentRequest create payment request
//
// swagger:model createPaymentRequest
type CreatePaymentRequest struct {

	// Fiat currency
	// Required: true
	// Enum: [USD EUR]
	Currency string `json:"currency"`

	// Optional payment description
	// Example: White T-shirt size M
	// Max Length: 128
	Description *string `json:"description,omitempty"`

	// To provide request idempotency order UUID should be generated on your side.
	// Should be unique for each payment
	//
	// Example: 123e4567-e89b-12d3-a456-426655440000
	// Required: true
	// Format: uuid
	ID strfmt.UUID `json:"id"`

	// Indicates that this is a test payment.
	// Test payments are processed in testnets (e.g. Ethereum Goerli)
	//
	IsTest bool `json:"isTest,omitempty"`

	// Optional order ID from your internal system
	// Example: customer#123#order#456
	OrderID *string `json:"orderId"`

	// Price in fiat currency
	// Example: 29.9
	// Required: true
	// Minimum: 0.01
	Price float64 `json:"price"`

	// A "back to store" button URL. Visible to a customer after the system receives unconfirmed transaction.
	// If not provided, defaults to merchant's website url.
	//
	// Example: https://my-store.com/success
	RedirectURL *string `json:"redirectUrl"`
}

// Validate validates this create payment request
func (m *CreatePaymentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createPaymentRequestTypeCurrencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["USD","EUR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createPaymentRequestTypeCurrencyPropEnum = append(createPaymentRequestTypeCurrencyPropEnum, v)
	}
}

const (

	// CreatePaymentRequestCurrencyUSD captures enum value "USD"
	CreatePaymentRequestCurrencyUSD string = "USD"

	// CreatePaymentRequestCurrencyEUR captures enum value "EUR"
	CreatePaymentRequestCurrencyEUR string = "EUR"
)

// prop value enum
func (m *CreatePaymentRequest) validateCurrencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createPaymentRequestTypeCurrencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreatePaymentRequest) validateCurrency(formats strfmt.Registry) error {

	if err := validate.RequiredString("currency", "body", m.Currency); err != nil {
		return err
	}

	// value enum
	if err := m.validateCurrencyEnum("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *CreatePaymentRequest) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", *m.Description, 128); err != nil {
		return err
	}

	return nil
}

func (m *CreatePaymentRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreatePaymentRequest) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", float64(m.Price)); err != nil {
		return err
	}

	if err := validate.Minimum("price", "body", m.Price, 0.01, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create payment request based on context it is used
func (m *CreatePaymentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreatePaymentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePaymentRequest) UnmarshalBinary(b []byte) error {
	var res CreatePaymentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
