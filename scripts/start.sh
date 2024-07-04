#!/bin/bash

# Initialiser la base de données
bash scripts/init_db.sh

echo "Current directory: $(pwd)"
# other commands to start the server

# Lancer le serveur
go run main.go
