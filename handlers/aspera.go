package handlers

type UploadEventResponse struct {
	FileName        string `json:"file_name"`
	SourcePath      string `json:"source_path"`
	DestinationPath string `json:"destination_path"`
	OutputResult    string `json:"output_result"`
}

type Credentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password"`
	KeyS3URI string `json:"keyS3URI"`
}
