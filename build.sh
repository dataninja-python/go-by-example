#!/usr/bin/env bash

read -p "Enter directory path: " dirPath

go build $dirPath/*
