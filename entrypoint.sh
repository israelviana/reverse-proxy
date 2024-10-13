#!/bin/sh

sleep 10

migrate -path=/app/migrations -database=postgres://postgres:postgres@db:5432/reverse_proxy?sslmode=disable up

exec ./myapp
