// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new project API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for project API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AllGitBranches gets all git branches

### Get All Git Branches

Returns a list of git branches in the project repository

*/
func (a *Client) AllGitBranches(params *AllGitBranchesParams) (*AllGitBranchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllGitBranchesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "all_git_branches",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}/git_branches",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllGitBranchesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AllGitBranchesOK), nil

}

/*
AllGitConnectionTests gets all git connection tests

### Get All Git Connection Tests

Returns a list of tests which can be run against a project's git connection. Call [Run Git Connection Test](#!/Project/run_git_connection_test) to execute each test in sequence.

Tests are ordered by increasing specificity. Tests should be run in the order returned because later tests require functionality tested by tests earlier in the test list.

For example, a late-stage test for write access is meaningless if connecting to the git server (an early test) is failing.

*/
func (a *Client) AllGitConnectionTests(params *AllGitConnectionTestsParams) (*AllGitConnectionTestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllGitConnectionTestsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "all_git_connection_tests",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}/git_connection_tests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllGitConnectionTestsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AllGitConnectionTestsOK), nil

}

/*
AllProjectFiles gets all project files

### Get All Project Files

Returns a list of the files in the project

*/
func (a *Client) AllProjectFiles(params *AllProjectFilesParams) (*AllProjectFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllProjectFilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "all_project_files",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}/files",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllProjectFilesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AllProjectFilesOK), nil

}

/*
AllProjects gets all projects

### Get All Projects

Returns all projects visible to the current user

*/
func (a *Client) AllProjects(params *AllProjectsParams) (*AllProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "all_projects",
		Method:             "GET",
		PathPattern:        "/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllProjectsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AllProjectsOK), nil

}

/*
CreateGitBranch checkouts new git branch

### Create and Checkout a Git Branch

Creates and checks out a new branch in the given project repository
Only allowed in development mode
  - Call `update_session` to select the 'dev' workspace.

Optionally specify a branch name, tag name or commit SHA as the start point in the ref field.
  If no ref is specified, HEAD of the current branch will be used as the start point for the new branch.


*/
func (a *Client) CreateGitBranch(params *CreateGitBranchParams) (*CreateGitBranchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateGitBranchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create_git_branch",
		Method:             "POST",
		PathPattern:        "/projects/{project_id}/git_branch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateGitBranchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateGitBranchOK), nil

}

/*
CreateGitDeployKey creates deploy key

### Create Git Deploy Key

Create a public/private key pair for authenticating ssh git requests from Looker to a remote git repository
for a particular Looker project.

Returns the public key of the generated ssh key pair.

Copy this public key to your remote git repository's ssh keys configuration so that the remote git service can
validate and accept git requests from the Looker server.

*/
func (a *Client) CreateGitDeployKey(params *CreateGitDeployKeyParams) (*CreateGitDeployKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateGitDeployKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create_git_deploy_key",
		Method:             "POST",
		PathPattern:        "/projects/{project_id}/git/deploy_key",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateGitDeployKeyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateGitDeployKeyOK), nil

}

/*
CreateProject creates project

### Create A Project

dev mode required.
- Call `update_session` to select the 'dev' workspace.

`name` is required.
`git_remote_url` is not allowed. To configure Git for the newly created project, follow the instructions in `update_project`.


*/
func (a *Client) CreateProject(params *CreateProjectParams) (*CreateProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create_project",
		Method:             "POST",
		PathPattern:        "/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateProjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateProjectOK), nil

}

/*
DeleteGitBranch deletes a git branch

### Delete the specified Git Branch

Delete git branch specified in branch_name path param from local and remote of specified project repository

*/
func (a *Client) DeleteGitBranch(params *DeleteGitBranchParams) (*DeleteGitBranchNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGitBranchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_git_branch",
		Method:             "DELETE",
		PathPattern:        "/projects/{project_id}/git_branch/{branch_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteGitBranchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteGitBranchNoContent), nil

}

/*
DeployToProduction deploys to production

### Deploy LookML from this Development Mode Project to Production

Git must have been configured, must be in dev mode and deploy permission required

Deploy is a two / three step process
1. Push commits in current branch of dev mode project to the production branch (origin/master).
   Note a. This step is skipped in read-only projects.
   Note b. If this step is unsuccessful for any reason (e.g. rejected non-fastforward because production branch has
             commits not in current branch), subsequent steps will be skipped.
2. If this is the first deploy of this project, create the production project with git repository.
3. Pull the production branch into the production project.


*/
func (a *Client) DeployToProduction(params *DeployToProductionParams) (*DeployToProductionOK, *DeployToProductionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeployToProductionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deploy_to_production",
		Method:             "POST",
		PathPattern:        "/projects/{project_id}/deploy_to_production",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeployToProductionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeployToProductionOK:
		return value, nil, nil
	case *DeployToProductionNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
FindGitBranch finds a git branch

### Get the specified Git Branch

Returns the git branch specified in branch_name path param if it exists in the given project repository

*/
func (a *Client) FindGitBranch(params *FindGitBranchParams) (*FindGitBranchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindGitBranchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "find_git_branch",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}/git_branch/{branch_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindGitBranchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindGitBranchOK), nil

}

/*
GitBranch gets active git branch

### Get the Current Git Branch

Returns the git branch currently checked out in the given project repository

*/
func (a *Client) GitBranch(params *GitBranchParams) (*GitBranchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGitBranchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "git_branch",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}/git_branch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GitBranchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GitBranchOK), nil

}

/*
GitDeployKey gits deploy key

### Git Deploy Key

Returns the ssh public key previously created for a project's git repository.

*/
func (a *Client) GitDeployKey(params *GitDeployKeyParams) (*GitDeployKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGitDeployKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "git_deploy_key",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}/git/deploy_key",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GitDeployKeyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GitDeployKeyOK), nil

}

/*
Project gets project

### Get A Project

Returns the project with the given project id

*/
func (a *Client) Project(params *ProjectParams) (*ProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "project",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ProjectOK), nil

}

/*
ProjectFile gets project file

### Get Project File Info

Returns information about a file in the project

*/
func (a *Client) ProjectFile(params *ProjectFileParams) (*ProjectFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "project_file",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}/files/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectFileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ProjectFileOK), nil

}

/*
ProjectValidationResults cacheds project validation results

### Get Cached Project Validation Results

Returns the cached results of a previous project validation calculation, if any.
Returns http status 204 No Content if no validation results exist.

Validating the content of all the files in a project can be computationally intensive
for large projects. Use this API to simply fetch the results of the most recent
project validation rather than revalidating the entire project from scratch.

A value of `"stale": true` in the response indicates that the project has changed since
the cached validation results were computed. The cached validation results may no longer
reflect the current state of the project.

*/
func (a *Client) ProjectValidationResults(params *ProjectValidationResultsParams) (*ProjectValidationResultsOK, *ProjectValidationResultsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectValidationResultsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "project_validation_results",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectValidationResultsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ProjectValidationResultsOK:
		return value, nil, nil
	case *ProjectValidationResultsNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
ProjectWorkspace gets project workspace

### Get Project Workspace

Returns information about the state of the project files in the currently selected workspace

*/
func (a *Client) ProjectWorkspace(params *ProjectWorkspaceParams) (*ProjectWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "project_workspace",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}/current_workspace",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ProjectWorkspaceOK), nil

}

/*
ResetProjectToProduction resets to production

### Reset a project to the revision of the project that is in production.

**DANGER** this will delete any changes that have not been pushed to a remote repository.

*/
func (a *Client) ResetProjectToProduction(params *ResetProjectToProductionParams) (*ResetProjectToProductionOK, *ResetProjectToProductionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetProjectToProductionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reset_project_to_production",
		Method:             "POST",
		PathPattern:        "/projects/{project_id}/reset_to_production",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResetProjectToProductionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ResetProjectToProductionOK:
		return value, nil, nil
	case *ResetProjectToProductionNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
ResetProjectToRemote resets to remote

### Reset a project development branch to the revision of the project that is on the remote.

**DANGER** this will delete any changes that have not been pushed to a remote repository.

*/
func (a *Client) ResetProjectToRemote(params *ResetProjectToRemoteParams) (*ResetProjectToRemoteOK, *ResetProjectToRemoteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetProjectToRemoteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reset_project_to_remote",
		Method:             "POST",
		PathPattern:        "/projects/{project_id}/reset_to_remote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResetProjectToRemoteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ResetProjectToRemoteOK:
		return value, nil, nil
	case *ResetProjectToRemoteNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
RunGitConnectionTest runs git connection test

### Run a git connection test

Run the named test on the git service used by this project and return the result. This
is intended to help debug git connections when things do not work properly, to give
more helpful information about why a git url is not working with Looker.

Tests should be run in the order they are returned by [Get All Git Connection Tests](#!/Project/all_git_connection_tests).

*/
func (a *Client) RunGitConnectionTest(params *RunGitConnectionTestParams) (*RunGitConnectionTestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunGitConnectionTestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "run_git_connection_test",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}/git_connection_tests/{test_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RunGitConnectionTestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunGitConnectionTestOK), nil

}

/*
UpdateGitBranch updates project git branch

### Checkout and/or reset --hard an existing Git Branch

Only allowed in development mode
  - Call `update_session` to select the 'dev' workspace.

Checkout an existing branch if name field is different from the name of the currently checked out branch.

Optionally specify a branch name, tag name or commit SHA to which the branch should be reset.
  **DANGER** hard reset will be force pushed to the remote. Unsaved changes and commits may be permanently lost.


*/
func (a *Client) UpdateGitBranch(params *UpdateGitBranchParams) (*UpdateGitBranchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGitBranchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update_git_branch",
		Method:             "PUT",
		PathPattern:        "/projects/{project_id}/git_branch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateGitBranchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateGitBranchOK), nil

}

/*
UpdateProject updates project

### Update Project Configuration

Apply changes to a project's configuration.


#### Configuring Git for a Project

To set up a Looker project with a remote git repository, follow these steps:

1. Call `update_session` to select the 'dev' workspace.
1. Call `create_git_deploy_key` to create a new deploy key for the project
1. Copy the deploy key text into the remote git repository's ssh key configuration
1. Call `update_project` to set project's `git_remote_url` ()and `git_service_name`, if necessary).

When you modify a project's `git_remote_url`, Looker connects to the remote repository to fetch
metadata. The remote git repository MUST be configured with the Looker-generated deploy
key for this project prior to setting the project's `git_remote_url`.

To set up a Looker project with a git repository residing on the Looker server (a 'bare' git repo):
1. Call `update_session` to select the 'dev' workspace.
1. Call `update_project` setting `git_remote_url` to nil and `git_service_name` to "bare".


*/
func (a *Client) UpdateProject(params *UpdateProjectParams) (*UpdateProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update_project",
		Method:             "PATCH",
		PathPattern:        "/projects/{project_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateProjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateProjectOK), nil

}

/*
ValidateProject validates project

### Validate Project

Performs lint validation of all lookml files in the project.
Returns a list of errors found, if any.

Validating the content of all the files in a project can be computationally intensive
for large projects. For best performance, call `validate_project(project_id)` only
when you really want to recompute project validation. To quickly display the results of
the most recent project validation (without recomputing), use `project_validation_results(project_id)`

*/
func (a *Client) ValidateProject(params *ValidateProjectParams) (*ValidateProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validate_project",
		Method:             "POST",
		PathPattern:        "/projects/{project_id}/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateProjectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ValidateProjectOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}