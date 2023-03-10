# article

I use existing libs :

 - Ozzo Validation, for input request validation
 - Godotenv, for env loader
 - jmoiron/sqlx for postgres driver


# For setup after cloning the repo:
> cd article
> go mod tidy

# to do a endpoint unit test :
> go to the article handler (article/src/interface/rest/handlers/article then run a command "go test"
> you can see the coverage testing by open the project with vscode, choose the testing file, right click then choose "Go:Toogle Test Coverage in Current Package"
>the result : 
```
Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-goR19VYP/go-code-cover article/src/interface/rest/handlers/article

ok  	article/src/interface/rest/handlers/article	0.444s	coverage: 100.0% of statements
```

# for db table :
> in folder db, there is a .sql file with the create table command. I use postgresql for this case. you can run the command in your sql editor page.

# redis
> you can use existing redis in your device, but if you dont have it, you can install it in docker by "docker-compose up" in terminal of this root directory