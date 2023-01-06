Simple CRUD API to write to mySQL db

To run the mysql container with docker run:
**NOTE
remember to set correct ports under -p flag
plus the db name, user and password should match the ones from config constants

$ docker run -d \
--name go-crud-api \
--mount type=volume,src=mysql_vol,target=/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=root_password  \
-e MYSQL_DATABSE=database \
-e MYSQL_USER=user \
-e MYSQL_PASSWORD=password  \
-p 8080:8080 \
mysql:8.0.31

Connect to running container:

$ docker exec -it go-crud-api bash

Connect o mysql DB:

$ mysql -h 127.0.0.1 -u root -p
