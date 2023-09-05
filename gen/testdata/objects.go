package testdata

import "github.com/mjdrgn/gql-rapid-gen/parser"

func DynamoDBSimple(name string) *parser.ParsedObject {
	return &parser.ParsedObject{
		Name: name,
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     name,
						"hash_key": "myhash",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
		},
		Interfaces: nil,
	}
}

func DynamoDBComposite(name string) *parser.ParsedObject {
	return &parser.ParsedObject{
		Name: name,
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     name,
						"hash_key": "myhash",
						"sort_key": "mysort",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "mysort",
				Directives:  nil,
				Description: "Sort Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
		},
		Interfaces: nil,
	}
}

func DynamoDBCompositeNumbers(name string) *parser.ParsedObject {
	return &parser.ParsedObject{
		Name: name,
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     name,
						"hash_key": "myhash",
						"sort_key": "mysort",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeIntReq,
			},
			{
				Name:        "mysort",
				Directives:  nil,
				Description: "Sort Key",
				Arguments:   nil,
				Type:        parser.TypeIntReq,
			},
		},
		Interfaces: nil,
	}
}

func DynamoDBGSI(name string) *parser.ParsedObject {
	return &parser.ParsedObject{
		Name: name,
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     name,
						"hash_key": "myhash",
					}),
				},
			},
			"dynamodb_gsi": {
				&parser.ParsedDirective{
					Name: "dynamodb_gsi",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "test1",
						"hash_key": "test1",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "test1",
				Directives:  nil,
				Description: "GSI Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
		},
		Interfaces: nil,
	}
}

func DynamoDBGSIOrdered(name string) *parser.ParsedObject {
	return &parser.ParsedObject{
		Name: name,
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     name,
						"hash_key": "myhash",
					}),
				},
			},
			"dynamodb_gsi": {
				&parser.ParsedDirective{
					Name: "dynamodb_gsi",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "test1",
						"hash_key": "test1",
						"sort_key": "test2",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "test1",
				Directives:  nil,
				Description: "GSI Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "test2",
				Directives:  nil,
				Description: "GSI Sort Key",
				Arguments:   nil,
				Type:        parser.TypeIntReq,
			},
		},
		Interfaces: nil,
	}
}

func DynamoDBGSIComposite(name string) *parser.ParsedObject {
	return &parser.ParsedObject{
		Name: name,
		Directives: map[string][]*parser.ParsedDirective{
			"dynamodb": {
				&parser.ParsedDirective{
					Name: "dynamodb",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     name,
						"hash_key": "myhash",
					}),
				},
			},
			"dynamodb_gsi": {
				&parser.ParsedDirective{
					Name: "dynamodb_gsi",
					Arguments: parser.ArgumentsFromMap(map[string]any{
						"name":     "test1_test2",
						"hash_key": "test1",
						"sort_key": "test2",
					}),
				},
			},
		},
		Description: "Description",
		Fields: []*parser.ParsedField{
			{
				Name:        "myhash",
				Directives:  nil,
				Description: "Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "test1",
				Directives:  nil,
				Description: "GSI Hash Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
			{
				Name:        "test2",
				Directives:  nil,
				Description: "GSI Sort Key",
				Arguments:   nil,
				Type:        parser.TypeStringReq,
			},
		},
		Interfaces: nil,
	}
}
