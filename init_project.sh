#!/bin/bash

# Création des répertoires
mkdir -p controllers models routes config scripts

# Création des fichiers principaux
touch main.go
touch controllers/urlController.go
touch models/url.go
touch routes/routes.go
touch config/database.go
touch scripts/init_db.sh

echo "Projet Go url-shortener initialisé avec succès."
