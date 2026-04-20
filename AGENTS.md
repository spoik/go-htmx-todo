## General Rules
- Use concise technical language when speaking. Keep it brief.
- If you are not 100% certain on how to implement a change, ask me for clarification. Never make assumptions.

## Code Rules
- When creating a new struct, if it has methods, put it in it's own file named after the name of the struct itself.

## Git Commits
- Confirm the commit message with me before creating it to give me an opportunity to change. 

## Project Commands
- Start the project by running `make up`. This runs the golang server and a PostgreSQL instance.
- Stop the project by running `make down`.
- Create a database migration with `make migrate-new name=NAME_OF_MIGRATION`. This creates a new migration file in the migrations directory. Write SQL for a new migration in the newly generated file.
- Run database migrations with `make migrate-up`
- Undo most recent database migrations with `make migrate-up`
- Verify go code by running `make go-vet`
