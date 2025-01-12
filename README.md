## Getting started
To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

## Test and Deploy
Use the built-in continuous integration in Github. 

***

# Editing this README
When you're ready to make this README your own, just edit this file and use the handy template below (or feel free to structure it however you want - this is just a starting point!). Thank you to [makeareadme.com](https://www.makeareadme.com/) for this template.

## Suggestions for a good README
Every project is different, so consider which of these sections apply to yours. The sections used in the template are suggestions for most open source projects. Also keep in mind that while a README can be too long and detailed, too long is better than too short. If you think your README is too long, consider utilizing another form of documentation rather than cutting out information.

## Name
Choose a lets-code-grpc name for your project.
 
## Badges
On some READMEs, you may see small images that convey metadata, such as whether or not all the tests are passing for the project. You can use Shields to add some to your README. Many services also have instructions for adding a badge.

## Installation
Within a particular ecosystem, there may be a common way of installing things, such as using Yarn, NuGet, or Homebrew. However, consider the possibility that whoever is reading your README is a novice and would like more guidance. Listing specific steps helps remove ambiguity and gets people to using your project as quickly as possible. If it only runs in a specific context like a particular programming language version or operating system or has dependencies that have to be installed manually, also add a Requirements subsection.

## Usage
Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

## Support
Tell people where they can go to for help. It can be any combination of an issue tracker, a chat room, an email address, etc.

## Roadmap
If you have ideas for releases in the future, it is a good idea to list them in the README. 

## Authors and acknowledgment
Show your appreciation to those who have contributed to the project.

## License
For project technical test, say how it is licensed.

## Project status 

## Please refer to the reading materials for further details
1. https://go.dev

## Documentation Articles
1. https://www.notion.so/gRPC-1752ffdf28e680a99253fdcb6e3682bb?pvs=4

## Documentation Api
1. https://app.swaggerhub.com/apis/lets-code-grpc/1.0.0

## Documentation ERD
1. https://drive.google.com/file/d/1H2Vwpg__PzfhTgXo_ayiQLWSoKfoSEHI/view?usp=sharing

## Go version 
1. $go version > go1.21.0

## Third-party Libraries 
Execute the following commands to install the libraries: 
```shellscript
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get gorm.io/driver/mysql
go get github.com/gin-gonic/gin
go get github.com/spf13/viper  
go get github.com/go-playground/validator/v10

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get google.golang.org/grpc

```

## Setup gRPC
1. Instal brew install protobuf
2. protoc --version
3. libprotoc 3.19.0
4. generate rpc -> protoc --go_out=. --go-grpc_out=. proto/user.proto
5. protoc --go_out=. --go-grpc_out=. proto/product.proto
5. add line grpc compiler option go_package = "./proto;proto"; 

   
## Setup
1. **Create the `.env` file**  
   In the `project_name/configuration/` directory, create a `.env` file to configure your environment variables.

## File Locations
1. .env file: /file/env
2. JSON files: /file/json
3. Product images for editing: /file/images/product_edit
4. Product images: /file/images/product
5. Dummy database data: /file/data_dummy

## Data Insertion Order
To ensure proper data relationships, insert data into the tables in the following order:

```shellscript
-- 1. users
INSERT INTO DOT_TEMP.users
SELECT * FROM DOT.users;

-- 2. balances
INSERT INTO DOT_TEMP.balances
SELECT * FROM DOT.balances;

-- 3. currencies
INSERT INTO DOT_TEMP.currencies
SELECT * FROM DOT.currencies;

-- 4. companies
INSERT INTO DOT_TEMP.companies
SELECT * FROM DOT.companies;


-- 5. products
INSERT INTO DOT_TEMP.products 
SELECT *
FROM DOT.products;

-- 6. product_prices
INSERT INTO DOT_TEMP.product_prices
SELECT * FROM DOT.product_prices;

-- 7. transactions
INSERT INTO DOT_TEMP.transactions
SELECT * FROM DOT.transactions;
```

## Running the Project
1. **Run the Project**  
   Execute the following command to run the project:
   ```bash
   go run main.go 

3. In the database.go file, there's an auto-migrate function that will automatically create the tables and their schemas in the database.
2. SQL Data Files
    Move all SQL files containing the data for each table to the appropriate folder.
3. JSON Export Files
    Move or export JSON files (e.g., file.json) to the appropriate tool (Postman, Swagger) for testing.

## Testing Project 
1. Token Requirement
    Every request will require a token.
2. Login to Get Your Token
    To obtain the token, simply run the exported JSON file at the endpoint: POST http://localhost:8089/users/login
