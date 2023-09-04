package data

// This file is derivative of gql-rapid-gen templates.
// Template Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD.

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type MultiTableProvider interface {
}

type multiTableImpl struct {
	db *dynamodb.Client
}

var multiTableProviderSingleton MultiTableProvider

func initMultiTableProvider(db *dynamodb.Client) {
	multiTableProviderSingleton = &multiTableImpl{
		db: db,
	}
}

func GetMultiTableProvider() MultiTableProvider {
	return multiTableProviderSingleton
}
