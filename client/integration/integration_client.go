// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new integration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for integration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AcceptIntegrationHubLegalAgreement accepts integration hub legal agreement

Accepts the legal agreement for a given integration hub. This only works for integration hubs that have legal_agreement_required set to true and legal_agreement_signed set to false.
*/
func (a *Client) AcceptIntegrationHubLegalAgreement(params *AcceptIntegrationHubLegalAgreementParams) (*AcceptIntegrationHubLegalAgreementOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAcceptIntegrationHubLegalAgreementParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "accept_integration_hub_legal_agreement",
		Method:             "POST",
		PathPattern:        "/integration_hubs/{integration_hub_id}/accept_legal_agreement",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcceptIntegrationHubLegalAgreementReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AcceptIntegrationHubLegalAgreementOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for accept_integration_hub_legal_agreement: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AllIntegrationHubs gets all integration hubs

### Get information about all Integration Hubs.

*/
func (a *Client) AllIntegrationHubs(params *AllIntegrationHubsParams) (*AllIntegrationHubsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllIntegrationHubsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "all_integration_hubs",
		Method:             "GET",
		PathPattern:        "/integration_hubs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllIntegrationHubsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AllIntegrationHubsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for all_integration_hubs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AllIntegrations gets all integrations

### Get information about all Integrations.

*/
func (a *Client) AllIntegrations(params *AllIntegrationsParams) (*AllIntegrationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllIntegrationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "all_integrations",
		Method:             "GET",
		PathPattern:        "/integrations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllIntegrationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AllIntegrationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for all_integrations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateIntegrationHub creates integration hub

### Create a new Integration Hub.

This API is rate limited to prevent it from being used for SSRF attacks

*/
func (a *Client) CreateIntegrationHub(params *CreateIntegrationHubParams) (*CreateIntegrationHubOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateIntegrationHubParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create_integration_hub",
		Method:             "POST",
		PathPattern:        "/integration_hubs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateIntegrationHubReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateIntegrationHubOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create_integration_hub: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteIntegrationHub deletes integration hub

### Delete a Integration Hub.

*/
func (a *Client) DeleteIntegrationHub(params *DeleteIntegrationHubParams) (*DeleteIntegrationHubNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIntegrationHubParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_integration_hub",
		Method:             "DELETE",
		PathPattern:        "/integration_hubs/{integration_hub_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIntegrationHubReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteIntegrationHubNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_integration_hub: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FetchIntegrationForm fetches remote integration form

Returns the Integration form for presentation to the user.
*/
func (a *Client) FetchIntegrationForm(params *FetchIntegrationFormParams) (*FetchIntegrationFormOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchIntegrationFormParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "fetch_integration_form",
		Method:             "POST",
		PathPattern:        "/integrations/{integration_id}/form",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FetchIntegrationFormReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FetchIntegrationFormOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for fetch_integration_form: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Integration gets integration

### Get information about a Integration.

*/
func (a *Client) Integration(params *IntegrationParams) (*IntegrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIntegrationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "integration",
		Method:             "GET",
		PathPattern:        "/integrations/{integration_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IntegrationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IntegrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for integration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IntegrationHub gets integration hub

### Get information about a Integration Hub.

*/
func (a *Client) IntegrationHub(params *IntegrationHubParams) (*IntegrationHubOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIntegrationHubParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "integration_hub",
		Method:             "GET",
		PathPattern:        "/integration_hubs/{integration_hub_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IntegrationHubReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IntegrationHubOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for integration_hub: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TestIntegration tests integration

Tests the integration to make sure all the settings are working.
*/
func (a *Client) TestIntegration(params *TestIntegrationParams) (*TestIntegrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestIntegrationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "test_integration",
		Method:             "POST",
		PathPattern:        "/integrations/{integration_id}/test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TestIntegrationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TestIntegrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for test_integration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateIntegration updates integration

### Update parameters on a Integration.

*/
func (a *Client) UpdateIntegration(params *UpdateIntegrationParams) (*UpdateIntegrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateIntegrationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update_integration",
		Method:             "PATCH",
		PathPattern:        "/integrations/{integration_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateIntegrationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateIntegrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update_integration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateIntegrationHub updates integration hub

### Update a Integration Hub definition.

This API is rate limited to prevent it from being used for SSRF attacks

*/
func (a *Client) UpdateIntegrationHub(params *UpdateIntegrationHubParams) (*UpdateIntegrationHubOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateIntegrationHubParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update_integration_hub",
		Method:             "PATCH",
		PathPattern:        "/integration_hubs/{integration_hub_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateIntegrationHubReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateIntegrationHubOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update_integration_hub: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
