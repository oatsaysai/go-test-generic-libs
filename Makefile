run:
	DB_IP=127.0.0.1 \
	DB_PORT=3306 \
	DB_USER=mariauser \
	DB_PASS=password \
	DB_NAME=report \
	S3_ENDPOINT=127.0.0.1:9000 \
	S3_ACCESS_KEY=minio \
	S3_SECRET_KEY=TpD46gwKTZ7mzcuEw5voVR7 \
	S3_USE_SSL=false \
	S3_BUCKET=test-report \
	S3_LOCATION=ap-southeast-1 \
		go run ./
