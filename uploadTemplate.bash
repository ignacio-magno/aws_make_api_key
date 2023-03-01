#!/bin/bash

# Uploads a template to AWS SES

aws ses create-template --cli-input-json file://template.json