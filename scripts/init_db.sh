#!/bin/bash


# Load environment variables from the .env file
if [ -f .env ]; then
    export $(cat .env | sed 's/#.*//g' | xargs)
else
    echo "file .env don't exists."
    exit 1
fi


# Use environment variables with default values if not set
DB_USER=${DB_USER:-root} 
DB_PASSWORD=${DB_PASSWORD:-""}  
DB_NAME=${DB_NAME:-"url_shortener"}  

# Créer la base de données
mysql -u$DB_USER -p$DB_PASSWORD -e "CREATE DATABASE IF NOT EXISTS $DB_NAME;"
