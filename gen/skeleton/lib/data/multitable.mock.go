package data

// This file is derivative of gql-rapid-gen templates.
// Template Copyright (c) 2023 under the MIT license per gql-rapid-gen/LICENSE.MD.

type multiTableMock struct {
}

func initMultiTableProviderMock() {
	multiTableProviderSingleton = &multiTableMock{}
}
