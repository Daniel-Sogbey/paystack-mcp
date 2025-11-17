package types

type Manifest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tools       []struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Method      string `json:"method"`
		Path        string `json:"path"`
		Description string `json:"description"`
		InputSchema struct {
			Type       string   `json:"type"`
			Required   []string `json:"required"`
			Properties struct {
				Email struct {
					Type string `json:"type"`
				} `json:"email"`
				Amount struct {
					Type string `json:"type"`
				} `json:"amount"`
				Currency struct {
					Type string `json:"type"`
				} `json:"currency"`
				Reference struct {
					Type string `json:"type"`
				} `json:"reference"`
				Metadata struct {
					Type string `json:"type"`
				} `json:"metadata"`
			} `json:"properties"`
		} `json:"input_schema"`
		OutputSchema struct {
			Type       string `json:"type"`
			Properties struct {
				Status struct {
					Type string `json:"type"`
				} `json:"status"`
				Message struct {
					Type string `json:"type"`
				} `json:"message"`
				Data struct {
					Type       string `json:"type"`
					Properties struct {
						AuthorizationURL struct {
							Type string `json:"type"`
						} `json:"authorization_url"`
						AccessCode struct {
							Type string `json:"type"`
						} `json:"access_code"`
						Reference struct {
							Type string `json:"type"`
						} `json:"reference"`
					} `json:"properties"`
				} `json:"data"`
			} `json:"properties"`
		} `json:"output_schema"`
	} `json:"tools"`
}
