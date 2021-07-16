package ischema

// Code generated by xo. DO NOT EDIT.

// ColumnOption represents a row from 'information_schema.column_options'.
type ColumnOption struct {
	TableCatalog SQLIdentifier `json:"table_catalog"` // table_catalog
	TableSchema  SQLIdentifier `json:"table_schema"`  // table_schema
	TableName    SQLIdentifier `json:"table_name"`    // table_name
	ColumnName   SQLIdentifier `json:"column_name"`   // column_name
	OptionName   SQLIdentifier `json:"option_name"`   // option_name
	OptionValue  CharacterData `json:"option_value"`  // option_value
}
