package errorwrapper

// Application error codes.
const (
	CodeUnknown  Code = ""          // Unclassified or unknown error.
	CodeInternal Code = "internal"  // Internal error or inconsistency.
	CodeInvalid  Code = "invalid"   // Validation failed.
	CodeNotFound Code = "not_found" // Entity does not exist.
)
