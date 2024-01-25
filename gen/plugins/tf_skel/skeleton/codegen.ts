import { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  schema: './backend/*.graphql',
  documents: [],
  ignoreNoDocuments: true, // for better experience with the watcher
  generates: {
    './frontend/src/gql/backend.ts': {
      plugins: ['typescript', 'typescript-operations', 'typescript-react-apollo'],
      config: {
        avoidOptionals: true,
        enumsAsTypes: true,
        futureProofEnums: true,
      },
    },
    './app/gql/backend.ts': {
      plugins: ['typescript', 'typescript-operations', 'typescript-react-apollo'],
      config: {
        avoidOptionals: true,
        enumsAsTypes: true,
        futureProofEnums: true,
      },
    }
  }
}

export default config