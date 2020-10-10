package version

import (
	v1 "test_db/version/v1"
	v2 "test_db/version/v2"
	v3 "test_db/version/v3"

	"github.com/go-gormigrate/gormigrate/v2"
)

// ModelSchemaList Model Structs
var ModelSchemaList = []*gormigrate.Migration{
	&v1.AddCreditCardTable,
	&v2.AddAgeColumnForUserTable,
	// Case3. 新版本的 Schema 依次添加，注：顺序严格。
	&v3.AddEmailColumnForUserTable,
}
