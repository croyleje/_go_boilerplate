package main

type PLACEHOLDER struct {
	Entry    string
	Value    int
	Name     []Name
	Database []Database
}

type Name struct {
	Value int
	Field bool
}

type Database struct {
	Value  int
	Field  bool
	String string
}

func nested_struct_literal() PLACEHOLDER {

	variable := PLACEHOLDER{
		Entry: "Entry",
		Value: 0,
		Name: []Name{
			{Value: 1, Field: true},
		},
		Database: []Database{
			{Value: 1, Field: true, String: "Example"},
		},
	}

	return variable

}
