package models

// Log - custom logger interface{}
type Log interface {
	CustomLog(level, errorMessage string)
}

// LogS struct to implement Log interface
type LogS struct{}

// CustomLog - function to handle logs
func (c *LogS) CustomLog(level, errorMessage string) {
	// your code
}
