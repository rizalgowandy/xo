package ischema

// Code generated by xo. DO NOT EDIT.

// ColumnDomainUsage represents a row from 'information_schema.column_domain_usage'.
type ColumnDomainUsage struct {
	DomainCatalog SQLIdentifier `json:"domain_catalog"` // domain_catalog
	DomainSchema  SQLIdentifier `json:"domain_schema"`  // domain_schema
	DomainName    SQLIdentifier `json:"domain_name"`    // domain_name
	TableCatalog  SQLIdentifier `json:"table_catalog"`  // table_catalog
	TableSchema   SQLIdentifier `json:"table_schema"`   // table_schema
	TableName     SQLIdentifier `json:"table_name"`     // table_name
	ColumnName    SQLIdentifier `json:"column_name"`    // column_name
}
