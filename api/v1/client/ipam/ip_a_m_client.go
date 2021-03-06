// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new ipam API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ipam API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteIPAMIP releases an allocated IP address
*/
func (a *Client) DeleteIPAMIP(params *DeleteIPAMIPParams) (*DeleteIPAMIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIPAMIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteIPAMIP",
		Method:             "DELETE",
		PathPattern:        "/ipam/{ip}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteIPAMIPReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteIPAMIPOK), nil

}

/*
PostIPAM allocates an IP address
*/
func (a *Client) PostIPAM(params *PostIPAMParams) (*PostIPAMCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIPAMParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIPAM",
		Method:             "POST",
		PathPattern:        "/ipam",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostIPAMReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostIPAMCreated), nil

}

/*
PostIPAMIP allocates an IP address
*/
func (a *Client) PostIPAMIP(params *PostIPAMIPParams) (*PostIPAMIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIPAMIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIPAMIP",
		Method:             "POST",
		PathPattern:        "/ipam/{ip}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostIPAMIPReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostIPAMIPOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
