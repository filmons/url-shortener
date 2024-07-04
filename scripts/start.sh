#!/bin/bash

# Initialiser la base de données
bash scripts/init_db.sh

# Lancer le serveur
go run main.go
