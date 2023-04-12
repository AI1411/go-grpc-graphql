#!/bin/sh

# LocalStack entrypoint
/docker-entrypoint.sh &

# Wait for LocalStack services to be up
echo "Waiting for LocalStack to be ready..."
awslocal s3 wait bucket-exists --bucket star-user-images

# Create default S3 bucket
echo "Creating S3 bucket..."
awslocal s3api create-bucket --bucket star-user-images --region us-east-1

# Keep container running
tail -f /dev/null