// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DBConnection d b connection
// swagger:model DBConnection
type DBConnection struct {

	// SQL statements (semicolon separated) to issue after connecting to the database. Requires `custom_after_connect_statements` license feature
	AfterConnectStatements string `json:"after_connect_statements,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// (Write-Only) Base64 encoded Certificate body for server authentication (when appropriate for dialect).
	Certificate string `json:"certificate,omitempty"`

	// Creation date for this connection
	// Read Only: true
	CreatedAt string `json:"created_at,omitempty"`

	// Database name
	Database string `json:"database,omitempty"`

	// Time zone of database
	DbTimezone string `json:"db_timezone,omitempty"`

	// (Read-only) SQL Dialect details
	// Read Only: true
	Dialect *Dialect `json:"dialect,omitempty"`

	// (Read/Write) SQL Dialect name
	DialectName string `json:"dialect_name,omitempty"`

	// Is this an example connection
	// Read Only: true
	Example *bool `json:"example,omitempty"`

	// (Write-Only) Certificate keyfile type - .json or .p12
	FileType string `json:"file_type,omitempty"`

	// Host name/address of server
	Host string `json:"host,omitempty"`

	// Additional params to add to JDBC connection string
	JdbcAdditionalParams string `json:"jdbc_additional_params,omitempty"`

	// Unix timestamp at start of last completed PDT reap process
	// Read Only: true
	LastReapAt string `json:"last_reap_at,omitempty"`

	// Unix timestamp at start of last completed PDT trigger check process
	// Read Only: true
	LastRegenAt string `json:"last_regen_at,omitempty"`

	// Cron string specifying when maintenance such as PDT trigger checks and drops should be performed
	MaintenanceCron string `json:"maintenance_cron,omitempty"`

	// Maximum size of query in GBs (BigQuery only, can be a user_attribute name)
	MaxBillingGigabytes string `json:"max_billing_gigabytes,omitempty"`

	// Maximum number of concurrent connection to use
	MaxConnections int64 `json:"max_connections,omitempty"`

	// Name of the connection. Also used as the unique identifier
	Name string `json:"name,omitempty"`

	// (Write-Only) Password for server authentication
	Password string `json:"password,omitempty"`

	// db_connection_override for this connection in the `pdt` maintenance context
	PdtContextOverride *DBConnectionOverride `json:"pdt_context_override,omitempty"`

	// Pool Timeout
	PoolTimeout int64 `json:"pool_timeout,omitempty"`

	// Port number on server
	Port string `json:"port,omitempty"`

	// Timezone to use in queries
	QueryTimezone string `json:"query_timezone,omitempty"`

	// Scheme name
	Schema string `json:"schema,omitempty"`

	// SQL Runner snippets for this connection
	// Read Only: true
	Snippets []*Snippet `json:"snippets"`

	// Precache tables in the SQL Runner
	SQLRunnerPrecacheTables bool `json:"sql_runner_precache_tables,omitempty"`

	// Use SSL/TLS when connecting to server
	Ssl bool `json:"ssl,omitempty"`

	// Name of temporary database (if used)
	TmpDbName string `json:"tmp_db_name,omitempty"`

	// Fields whose values map to user attribute names
	UserAttributeFields []string `json:"user_attribute_fields"`

	// (Limited access feature) Are per user db credentials enabled. Enabling will remove previously set username and password
	UserDbCredentials bool `json:"user_db_credentials,omitempty"`

	// Id of user who last modified this connection configuration
	// Read Only: true
	UserID string `json:"user_id,omitempty"`

	// Username for server authentication
	Username string `json:"username,omitempty"`

	// Whether the connection uses OAuth for authentication.
	// Read Only: true
	UsesOauth *bool `json:"uses_oauth,omitempty"`

	// Verify the SSL
	VerifySsl bool `json:"verify_ssl,omitempty"`
}

// Validate validates this d b connection
func (m *DBConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDialect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePdtContextOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnippets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DBConnection) validateDialect(formats strfmt.Registry) error {

	if swag.IsZero(m.Dialect) { // not required
		return nil
	}

	if m.Dialect != nil {
		if err := m.Dialect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dialect")
			}
			return err
		}
	}

	return nil
}

func (m *DBConnection) validatePdtContextOverride(formats strfmt.Registry) error {

	if swag.IsZero(m.PdtContextOverride) { // not required
		return nil
	}

	if m.PdtContextOverride != nil {
		if err := m.PdtContextOverride.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pdt_context_override")
			}
			return err
		}
	}

	return nil
}

func (m *DBConnection) validateSnippets(formats strfmt.Registry) error {

	if swag.IsZero(m.Snippets) { // not required
		return nil
	}

	for i := 0; i < len(m.Snippets); i++ {
		if swag.IsZero(m.Snippets[i]) { // not required
			continue
		}

		if m.Snippets[i] != nil {
			if err := m.Snippets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snippets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DBConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DBConnection) UnmarshalBinary(b []byte) error {
	var res DBConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
