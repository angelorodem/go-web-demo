# go-web-demo
Simple go project

migrations are using go-migrate
`migrate -path migrations -database "sqlite3://demo.db" up` to create base db
or
`migrate -path migrations-mock -database "sqlite3://demo.db" up` to create base db with mock data

user operations:

create user
`curl --location 'localhost:8080/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "angelo",
    "email": "angelorodem@gmail.com",
    "password": "VeryNicePassw00rd!"
}'`

delete user (use token from login)
`curl --location --request DELETE 'localhost:8080/user' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer MOCK_VALID_JWT' \
--data-raw '{
    "email": "angelorodem@gmail.com"
}'`

change username (use token from login)
`curl --location --request PATCH 'localhost:8080/user' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer MOCK_VALID_JWT' \
--data-raw '{
    "email": "angelorodem@gmail.com",
    "newUsername": "Angelus IV"
}'`

get user (use token from login)
`curl --location --request GET 'localhost:8080/user' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer MOCK_VALID_JWT' \
--data-raw '{
    "email": "angelorodem@gmail.com"
}'`

login
`curl --location 'localhost:8080/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "angelorodem@gmail.com",
    "password": "VeryNicePassw00rd!"
}'`

