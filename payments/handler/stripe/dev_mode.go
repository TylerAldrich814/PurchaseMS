package handler

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"log"
	"os/exec"
	"strings"
)

func startStripeForwardPort(
  ctx     context.Context,
  url     string,
  execCmd *exec.Cmd,
) error {
  log.Printf("startStripeForwardPort")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	execCmd.Stdout = &stdout
	execCmd.Stderr = &stderr

	// ->> Start the Stripe CLI process in the background
	if err := execCmd.Start(); err != nil {
		return err
	}

  log.Print(execCmd.Args)
	// ->> Read command output asynchronously
	go func() {
		scanner := bufio.NewScanner(&stdout)
		for scanner.Scan() {
      log.Printf("[Stripe CLI]: %s",scanner.Text()) // Log the Stripe CLI output
		}
	}()
  // ->> Read for any potential Erros
  go func() {
    scanner := bufio.NewScanner(&stderr)
    for scanner.Scan() {
      log.Printf("[Stripe CLI Error]: %s", scanner.Text())
    }
  }()

  log.Printf("->> Stripe Server Started @ %s", url)
  return nil
}

// Helper function for obtaining Stripe forward Secret by 
// running the command
// $ stripe listen --forward-to <URL> --print-secret
func extractStripeSecret(
  url string,
)( string,error ){
  cmd := exec.Command(
    "stripe", 
    "listen", 
    "--forward-to",
    url,
    "--print-secret",
  )
  var stdout bytes.Buffer
  var stderr bytes.Buffer
  cmd.Stdout = &stdout
  cmd.Stderr = &stderr

  if err := cmd.Start(); err != nil {
    return "", err
  }
  
  go func(){
    scanner := bufio.NewScanner(&stdout)
    for scanner.Scan(){
      log.Println(scanner.Text())
    }
  }()
  if err := cmd.Wait(); err != nil {
    return "", err
  }

  output := stdout.String()

  lines := strings.Split(output, "\n")
  for _, line := range lines {
    if strings.HasPrefix(line, "whsec_"){
      return line, nil
    }
  }
  return "", errors.New("Stripe Webhook Secret not found in output") 
}
