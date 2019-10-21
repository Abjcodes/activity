package vocab

// A public key represents a public cryptographical key for a user
type W3IDSecurityV1PublicKey interface {
	// GetJSONLDId returns the "id" property if it exists, and nil otherwise.
	GetJSONLDId() JSONLDIdProperty
	// GetTypeName returns the name of this type.
	GetTypeName() string
	// GetUnknownProperties returns the unknown properties for the PublicKey
	// type. Note that this should not be used by app developers. It is
	// only used to help determine which implementation is LessThan the
	// other. Developers who are creating a different implementation of
	// this type's interface can use this method in their LessThan
	// implementation, but routine ActivityPub applications should not use
	// this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// GetW3IDSecurityV1Owner returns the "owner" property if it exists, and
	// nil otherwise.
	GetW3IDSecurityV1Owner() W3IDSecurityV1OwnerProperty
	// GetW3IDSecurityV1PublicKeyPem returns the "publicKeyPem" property if it
	// exists, and nil otherwise.
	GetW3IDSecurityV1PublicKeyPem() W3IDSecurityV1PublicKeyPemProperty
	// IsExtending returns true if the PublicKey type extends from the other
	// type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this PublicKey is lesser, with an arbitrary but
	// stable determination.
	LessThan(o W3IDSecurityV1PublicKey) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetJSONLDId sets the "id" property.
	SetJSONLDId(i JSONLDIdProperty)
	// SetW3IDSecurityV1Owner sets the "owner" property.
	SetW3IDSecurityV1Owner(i W3IDSecurityV1OwnerProperty)
	// SetW3IDSecurityV1PublicKeyPem sets the "publicKeyPem" property.
	SetW3IDSecurityV1PublicKeyPem(i W3IDSecurityV1PublicKeyPemProperty)
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}
