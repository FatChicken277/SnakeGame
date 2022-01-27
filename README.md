# SNAKEE :snake:

![printf](https://user-images.githubusercontent.com/60364301/151388149-2726d77e-df6c-40a2-be41-7fe162ae540f.png)

SNAKE is a snake game based on the phaser3 graphics engine and on vuejs for the frontend and a Golang API as the backend.

## Requirements :pencil:

- Golang
- Npm
- Vuejs
- CockroachDB database
- Postgresql

## Installation :wrench:

- Clone the repository
   ```bash
   $ git clone https://github.com/FatChicken277/snake_game.git
   $ cd snake_game/
   ```
- Configure connection to database:
   ```bash
   $ cd backend/
   ```
   After you enter the backend folder in the main.go file you will find a variable called [DatabaseSource](https://github.com/FatChicken277/snake_game/blob/20a15656994b9eb7d5a26ebabf88d587b6a17bb3/backend/main.go#L17) which is the path of the database, change it to your liking for the path of your database

- Install frontend:

   ```bash
   $ cd frontend/
   $ npm install
   ```


## Usage :trophy:

- Start the database
- Define env variables: 
  ``` 
  CLUSER= User name
  CLPWD= Password of database
  CLHOST= localhost
  CLPORT= 26257
  CLDB= Name of the database
  CAPATH= Path of the cockroach certificates $HOME/.postgresql/  root.crt
  DB= Database name
  SECRET= JWT Secret
  ALLOWEDORIGINS= *
  ```
- Start the API:

  ```bash
  $ cd backend/
  $ ENV VARIABLES go run main.go
  ```
- Start the Frontend:

  ```bash
  $ cd frontend/
  $ npm run serve
  ```

Now you can enjoy SNAKEE at [https://www.snakee.xyz/](https://www.snakee.xyz/) or play local at [http://localhost:8081/](http://localhost:8081/) :video_game:

## Contributing :family:
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
