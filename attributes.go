package growthbook

import "github.com/growthbook/growthbook-golang"

type Attributes struct {
	UserID      *string
	Email       *string
	AnonymousID *string
	IsInternal  *bool
}

func getAttributes(attributes Attributes) growthbook.Attributes {
	attrs := growthbook.Attributes{}
	attrs = populateFieldIfSet(attrs, "id", attributes.UserID)
	attrs = populateFieldIfSet(attrs, "anonymousId", attributes.AnonymousID)
	attrs = populateFieldIfSet(attrs, "email", attributes.Email)

	return populateFieldIfSet(attrs, "isInternal", attributes.IsInternal)
}

func populateFieldIfSet[TValue any](attributes growthbook.Attributes, key string, value *TValue) growthbook.Attributes {
	if value == nil {
		return attributes
	}
	attributes[key] = *value

	return attributes
}
