# ANWB APP
This application showcases use of the ANWB api with a simple example.

## How it works
Using a cron-like method the API fetches the latest version of the ANWB information and stores this document as-is in the database. Since the source material is essentially one big JSON document, MongoDB has been chosen as database.

Using a REST interface the API exposes an index of the stored documents and a few convenience endpoints for fetching various road events like traffic jams and roadworks.

Finally the UI has two simple features. One feature is a slider that allows the user to select a point in time for which the ANWB data will be fetched. This data is shown in the second feature, which is a set of 2 tables with the traffic jams and road works of that time period.

## Running the app
To run the app you can either use the docker compose file, or start each component natively.

### API key
This app requires an API key for the ANWB API. Reach out to a contributor for a working key. See the section running natively below on where to insert the key.

### Using docker compose
Run `docker compose up` in the root directory.

### Running natively
- Run `docker compose up mongodb` to get a running instance of the database.
- Navigate to the api directory. Copy `.env.DIST` to `.env`, add the API key, run `go mod download` and then `go run .`.
- Navigate to the app direction. Copy `.env.local.DIST` to `.env.local`, run `yarn install` and `yarn dev`.

## Considerations
As this is a sample app it was not made with the intention to be production ready. To see what would probably need to resolved first (at the very least), search for open TODO's in the codebase.