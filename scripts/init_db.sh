#!/bin/bash

DB_USER="root"
DB_PASSWORD="dbfilmon"
DB_NAME="url_shortener"

# Créer la base de données
mysql -u$DB_USER -p$DB_PASSWORD -e "CREATE DATABASE IF NOT EXISTS $DB_NAME;"
