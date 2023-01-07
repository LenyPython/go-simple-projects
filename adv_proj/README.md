Simple CRUD API to write to mySQL db

To run the mysql container with docker run:
**NOTE
remember to set correct ports under -p flag
plus the db name, user and password should match the ones from config constants

** IMPORTANT
If you mount a volume second or more time the database won't be created with
env variables given to docker. If you require to create a DB and user 
MOUNT NEW volumes only or creat user and DB manually


$ docker run -d \
--name go-crud-api \
--mount type=volume,src=mysql_vol,target=/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=root_password  \
-e MYSQL_DATABASE=go_db \
-e MYSQL_USER=user \
-e MYSQL_PASSWORD=password  \
-p 3306:3306 \
mysql:8.0.31

Connect to running container:

$ docker exec -it go-crud-api bash

Connect o mysql DB:

$ mysql -h 127.0.0.1 -u root -p
