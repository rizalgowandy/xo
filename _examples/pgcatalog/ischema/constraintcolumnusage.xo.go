package ischema

// Code generated by xo. DO NOT EDIT.

// ConstraintColumnUsage represents a row from 'information_schema.constraint_column_usage'.
type ConstraintColumnUsage struct {
	TableCatalog      SQLIdentifier `json:"table_catalog"`      // table_catalog
	TableSchema       SQLIdentifier `json:"table_schema"`       // table_schema
	TableName         SQLIdentifier `json:"table_name"`         // table_name
	ColumnName        SQLIdentifier `json:"column_name"`        // column_name
	ConstraintCatalog SQLIdentifier `json:"constraint_catalog"` // constraint_catalog
	ConstraintSchema  SQLIdentifier `json:"constraint_schema"`  // constraint_schema
	ConstraintName    SQLIdentifier `json:"constraint_name"`    // constraint_name
}
