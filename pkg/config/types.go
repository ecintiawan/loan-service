package config

// Config holds necessary config to run the service.
type Config struct {
	App        App        `json:"app"`
	Vendor     Vendor     `json:"vendor"`
	Credential Credential `json:"credential"`
}

type (
	// App holds config value necessary to run application
	App struct {
		Port string `json:"port"`
	}

	// Vendor holds specific config value
	Vendor struct {
		Upload                 UploadConfig      `json:"upload"`
		Email                  EmailConfig       `json:"email"`
		DefaultApprovalProof   DefaultFileConfig `json:"default_approval_proof"`
		DefaultAgreementLetter DefaultFileConfig `json:"default_agreement_letter"`
	}

	// Credential config
	Credential struct {
		DB    CredentialDB    `json:"db_secret"`
		Email CredentialEmail `json:"email_secret"`
	}
)

type (
	// UploadConfig holds all upload configs
	UploadConfig struct {
		Path string `json:"path"`
		URL  string `json:"url"`
	}

	// DefaultFileConfig holds all local default file configs
	DefaultFileConfig struct {
		FilePath     string `json:"file_path"`
		DestFileName string `json:"destination_file_name"`
	}

	// EmailConfig holds all email configs
	EmailConfig struct {
		SMTPHost   string `json:"smtp_host"`
		SMTPPort   string `json:"smtp_port"`
		SenderName string `json:"sender_name"`
	}

	// CredentialDB holds all database credential
	CredentialDB struct {
		URL string `json:"url"`
	}

	// CredentialEmail holds all email credential
	CredentialEmail struct {
		SenderEmail    string `json:"sender_email"`
		SenderPassword string `json:"sender_password"`
	}
)
