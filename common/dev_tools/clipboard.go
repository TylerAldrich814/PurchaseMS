package devtools

import (
	"fmt"
	"github.com/TylerAldrich814/common"
	_ "github.com/joho/godotenv/autoload"
	"golang.design/x/clipboard"
)

var (
  enableClipboard = common.EnvString("DEV_CLIPBOARD", "")
)

// Helper function for taking a stirng and placing it into 
// the clipboard.
// This Method will require the service package that is using
// it to include 'DEV_CLIPBOARD="TRUE"' in it's .env file.
func ClipboardCopy(data string) error {
  if enableClipboard != "TRUE" { return nil }

  err := clipboard.Init()
  if err != nil {
    return fmt.Errorf(" ->> Failed to initialize Clipboard: %v\n", err)
  }
  clipboard.Write(clipboard.FmtText, []byte(data))

  return nil
}
