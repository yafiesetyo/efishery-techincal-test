package repository

type (
	User struct {
		Phone     string `json:"phone" mapstructure:"phone"`
		Name      string `json:"name" mapstructure:"name"`
		Role      string `json:"role" mapstructure:"role"`
		Password  string `json:"password" mapstructure:"password"`
		CreatedAt string `json:"created_at" mapstructure:"created_at"`
	}
)
