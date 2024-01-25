# GraphQL Rapid Application Generator - Schema Annotations

The primary method of creating your project is through annotations in the GraphQL schema.

A basic understanding of GraphQL is expected, and this document does not explain GraphQL concepts except where they are extended or modified.

## aws_*
These are defined by AWS AppSync and are used for permissions and other metadata within AWS.

## go_ignore
On: Field, Input Field, Type, Input Type

Parameters:
- None

This directive makes the backend code generators ignore a field, type, or input. Used for objects that will never be exposed to backend, or fields that are not needed (such as resolved foreign keys on types) that should not be directly persisted to a database.

## DynamoDB related
It is strongly recommended to read the AWS DynamoDB documentation if you are not familiar, particularly around design and choice of Hash and Sort keys, and the use of GSIs and LSIs.

### dynamodb
On: Type

Parameters:
- name: String, Required. The DynamoDB table name
- hash_key: String, Required. The DynamoDB Hash (or Partition) Key name
- sort_key: String, Optional. The DynamoDB Sort (or Range) Key name

This directive defines that a Type is backed by a DynamoDB table and instructions various plugins to generate infrastructure and data layer code around this.

### dynamodb_gsi
On: Type

Parameters:
- name: String, Required. The index name
- hash_key: String, Required. The DynamoDB Hash (or Partition) Key name
- sort_key: String, Optional. The DynamoDB Sort (or Range) Key name

This directive defines a Global Secondary Index on an existing DynamoDB Table. It should be applied only to a Type which also has a @dynamodb directive.
The `name` field should almost always take the form of either:
- The hash key, an underscore, and the sort key - if you regularly query both of these attributes together
- The hash key only, if it's primarily queried by this attribute and sorted by the other.
The latter naming convention will be detected by the generator and provide more useful query functions automatically.

### dynamodb_lsi
On: Type

Parameters:
- name: String, Required. The index name
- sort_key: String, Required. The DynamoDB Sort (or Range) Key name 

This directive defines a Local Secondary Index on an existing DynamoDB Table. It should be applied only to a Type which also has a @dynamodb directive.
Note that LSIs are a lot less flexible than GSIs, particularly for changes, so this should be used with caution.

## AppSync

### appsync_crud
On: Type

Parameters:
- create_hash_type: String, Optional. If Create is enabled, and this field is specified, the Hash Key will be auto generated instead of taken from the input. Valid options are "uuid" or "timestamp".
- create_sort_type: String, Optional. Same as create_hash_type but for the Sort Key.
- disable_create: Boolean, Optional. Disable generation of Create actions.
- disable_read: Boolean, Optional. Disable generation of Read/Get actions.
- disable_update: Boolean, Optional. Disable generation of Update actions.
- disable_delete: Boolean, Optional. Disable generation of Delete actions.

Generates the most commonly used CRUD queries and mutations for a given type.
These include:
- createX
- getX
- updateX
- deleteX

You must create matching Input Types for these actions, called XCreateInput, XUpdateInput, and XDeleteInput.
- XCreateInput should have all fields except any auto generated keys (using create_hash_type or create_sort_type).
- XUpdateInput should have all fields.
- XDeleteInput should have all key fields.

### appsync_scan
On: Type

Parameters:
- plural: String, Required. The Plural form of the object name, such as "Trees" for "Tree"

Generates a listAllX Query for the given Type.

### appsync_list
On: Field

Parameters:
- plural: String, Required. The Plural form of the object name, such as "Trees" for "Tree"
- using: String, Optional. The Database index to perform the query on.
- forward: Boolean, Optional. Whether to return results in ascending index order.
- name: String, Optional. Overwrites automatic Query name if specified.

Generates listXByY Queries for the given Type, using the field on which this directive is applied.

### appsync_foreign_key
On: Field

Parameters:
- query: Boolean, Optional. This resolver returns multiple values (e.g. the Field type is an Array)
- query_single: Boolean, Optional. This resolver needs to perform a multiple value retrieval, however it should only return the first entry (and never return more than one).
- batch: Boolean, Optional. This resolver is fed an array of values and should retrieve them all by ID rather than performing a query.
- field_source: String, Optional. Field in the current Type to use for comparison
- field_foreign: String, Optional. Field in the foreign table to use for comparison
- table: String, Required. The foreign table to read. IMPORTANT NOTE: This is the raw table name, NOT the Type name.
- index: String, Optional. The index to use on the foreign table, if specified.
- additional_field_source: String, Optional. Same as field_source but for an optional second comparison for composite keys or indexes.
- additional_field_foreign: String, Optional. Same as field_foreign but for an optional second comparison for composite keys or indexes.

Generates AppSync resolvers to allow retrieval of foreign referenced data from other tables (or potentially the same table).

This has the following four lookup patterns:
- query, query_single, batch all False: Lookup a single item in a table by it's key exactly. For example, this is used when you have the ID of another entity stored in another field in this table, such as as parent_id or other_entity_id. The Field type must be an Object matching the schema of the table referenced.
- query_single True: The same as above, except where the database layer does not allow a direct retrieval and we must query (for example, DynamoDB if you need to match a Global Secondary Index does not allow GetItem). The Field type must be an Object matching the schema of the table referenced.
- batch True: Similar to the first case, except where we have an Array of IDs. For example, if you had a parent-child architecture and this table contained ChildrenIDs \[String!]!. The Field type must be an Array of Objects matching the schema of the table referenced.
- query True: Performs a Query returning multiple entries into an Array, filtering by the field(s) specified. The Field type must be an Array of Objects matching the schema of the table referenced.

### appsync_sensitive_data
On: Field

Parameters:
- match_attribute: String, Required. The Field on this Type that must be matched for sensitive data to be released.
- user_claim: String, Required. The User attribute from the authentication JWT that must match the specified attribute.
- override_groups: String Array, Optional. JWT User Groups that can bypass this check (for example, Admins).

This directive adds additional constraints to `appsync_foreign_key` resolvers to ensure that private data is not leaked.

A common example would be: `match_attribute="user_email", user_claim="email", override_groups=["Admins"]`. This would ensure that no entries can be returned unless they have a user_email matching the current users email address, or the current user is an Admin.

### appsync_lambda
On: Query or Mutation

Parameters:
- language: String, Required. The backend programming language to be used. Supported is "go", "python", "nodejs"
- path: String, Required. The folder name within the "backend/lambda/" directory that this Lambda will be contained in. For cross-platform compatibility this should always be lowercase and avoid special characters.
- timeout: String, Required. A numeric string containing the number of seconds for the lambda timeout.
- memory: String, Required. A numeric string containing the amount of Lambda memory to allocate, in MB.

Generates AppSync resolvers and code skeletons for a Lambda function. Permissions will have to be assigned once the skeleton is generated.

## Data

### crud_type
On: Field

Parameters:
- render: String, Optional. How to display this value in text pages, grids, etc. Valid values: datetime | date | time | currency | fkey | cms_image | phone | markdown
- input: String, Optional. What kind of input element to use for forms. Value values: select | cms_image | datetime | date | time | multiselect | markdown | currency | phone
- non_empty: Boolean, Optional. Whether a zero-value is permitted (empty string, 0, etc)
- default: String, Optional. Default value to be used on Create if none specified.
- help_text: String. Optional. Additional info to render next to input boxes.
- validation_regex: String, Optional. Regex for additional validation as required. Extended regex features should not be used.
- min: Int, Optional. For Integer values, minimum value. For String values, minimum length.
- max: Int, Optional. For Integer values, maximum value. For String values, maximum length.
- readonly: Boolean, Optional. Does not allow frontend updates of this field. This does not enforce extra security on the backend or within code gen, so custom backend code or API clients may still modify it.
- fkey_resolver: String, Optional. For render type fkey, or input types select / multiselect.
- fkey_name: String, Optional. For render type fkey, or input types select / multiselect.

This directive primarily impacts the rendering of forms on the frontend (web or mobile), but also is used for validation rules in the backend on a best-effort basis.

### normalise
On: Field

Parameters:
- force_lower: Boolean, Optional. If True, the backend and frontend should convert this value to lowercase.
- trim: Boolean, Optional. If True, the backend and frontend should trim whitespace from both ends.

Normalise modifies frontend and backend generation to clean up values every time they are ingested, stored, read, or otherwise handled.
This is strongly recommended for key user-entered values, such as email addresses or phone numbers.






