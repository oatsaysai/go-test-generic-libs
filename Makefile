run-lambda:
	DB_IP=127.0.0.1 \
	DB_PORT=3306 \
	DB_USER=mariauser \
	DB_PASS=password \
	DB_NAME=report \
		go run ./ -lambda=true

run-test-libs:
	DB_IP=127.0.0.1 \
	DB_PORT=3306 \
	DB_USER=mariauser \
	DB_PASS=password \
	DB_NAME=report \
		go run ./ -lambda=false
