package main

import (
	"log/slog"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
  if err := godotenv.Load(); err != nil {
    panic(err)
  }

  cmd := exec.Command(
    "tern",
    "migrate",
    "--migrations",
    "./internal/database/pgdb/migrations",
    "--config",
    "./configs/tern.conf",
  )
  if err := cmd.Run(); err != nil {
    slog.Error("Failed to run tern:")
    panic(err)
  }
}
