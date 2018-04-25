// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new projects API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for projects API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddProject creates a project

Add a new project to an organization
*/
func (a *Client) AddProject(params *AddProjectParams, authInfo runtime.ClientAuthInfoWriter) (*AddProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addProject",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationName}/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddProjectOK), nil

}

/*
AddRelease cuts a release

Add a release to an existing project
*/
func (a *Client) AddRelease(params *AddReleaseParams, authInfo runtime.ClientAuthInfoWriter) (*AddReleaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddReleaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addRelease",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationName}/projects/{projectName}/releases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddReleaseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddReleaseOK), nil

}

/*
EditProject edits a project

Edit a project
*/
func (a *Client) EditProject(params *EditProjectParams, authInfo runtime.ClientAuthInfoWriter) (*EditProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editProject",
		Method:             "POST",
		PathPattern:        "/organizations/{organizationName}/projects/{projectName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EditProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EditProjectOK), nil

}

/*
GetProject organizations project info

Get project details
*/
func (a *Client) GetProject(params *GetProjectParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProject",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationName}/projects/{projectName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProjectOK), nil

}

/*
ListProjects organizations projects

Return a list of projects for an organization
*/
func (a *Client) ListProjects(params *ListProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*ListProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listProjects",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationName}/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListProjectsOK), nil

}

/*
ListReleases projects releases

Return a list of releases for a project
*/
func (a *Client) ListReleases(params *ListReleasesParams, authInfo runtime.ClientAuthInfoWriter) (*ListReleasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListReleasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listReleases",
		Method:             "GET",
		PathPattern:        "/organizations/{organizationName}/projects/{projectName}/releases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListReleasesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListReleasesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
