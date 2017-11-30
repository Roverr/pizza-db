# pizza-db
Generic pizza database structure with minimal CRUD

# Requires
- Go build environment
- Docker
- NPM
- Make

# Environment variables
 - `PIZZA_DB_URL` - Url in the following form: `username:password@tcp(host:port)/pizzeria?parseTime=true`
 - `PIZZA_DB_PASSWORD` - Root password of the database
 - `PIZZA_LISTEN_ADDRESS` - Address to listen to in `host:port` form

# Quick steps to start
- Set env variables (use [DIRENV](https://github.com/direnv/direnv) for example)
- `npm run db` - Creates DB container and starts it on port `3306`
- `npm run generate` - Generates data into the database (requires <b> pizza-generation </b> to be compiled)
- `npm start` - Creates a new container from the application and starts it, directed to connect to previously started db


# Other Commands
- `npm run dev` - Start webpack developer server
- `npm run clean-db` - Basically drops the database (requires `pizza-generation` to be compiled)
# UI
After building of the frontend and backend, the application have a UI which is located at
`$URL/ui`
