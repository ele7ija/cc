#!/bin/sh
# Apply database migrations
echo "Create database migrations"
python3 mysite/manage.py makemigrations

echo "Apply database migrations"
python3 mysite/manage.py migrate

echo "Starting server"
python3 mysite/manage.py runserver 0.0.0.0:8081