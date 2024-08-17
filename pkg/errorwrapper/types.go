package errorwrapper

type (
	// Message is a human-readable message.
	Message string

	// Code defines the kind of error this is, mostly for use by systems
	// that must act differently depending on the error.
	Code string

	// Error defines a standard application error.
	Error struct {
		// Underlying error.
		Err error

		// Codes used for Errs to identify known errors in the application.
		// If the error is expected by Errs object, the errors will be shown as listed in Codes.
		Code Code

		// Message is a human-readable message.
		Message Message
	}

	// errorString is a trivial implementation of error.
	errorString struct {
		s string
	}
)
