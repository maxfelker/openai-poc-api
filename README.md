# Open API POC

Simple Golang API using mux to support openai chat completions. 

## Getting started with local development 

First, clone the repo and go into the `openai-poc-api/` directory:

```bash
git clone git@github.com:maxfelker/openai-poc-api.git
cd openai-poc-api/
```

### Creating environment variable file

In the root of the project, you'll need to create the `.env` file which contains OpenAI keys. Here is what the file should look like:

```
OPENAI_API_KEY=
OPENAI_ORGANIZATION_ID=
```

You can get the API Key from your [platform dashboard here](https://platform.openai.com/api-keys) and [your organization ID here](https://platform.openai.com/account/organization). 

### Building and running the dev service 

Now you are ready to run the `dev` Docker service and start the API:

```bash
docker compose up --build dev
```

This will run the API locally at [http://localhost:8000](http://localhost:8000). The dev service has hot reloading support using [air](https://github.com/cosmtrek/air).

### Building and running the release service

If you want to build a production version, build and run the `release` Docker service:

```bash
docker compose up --build release
```

This will run the API locally at [http://localhost](http://localhost) (port 80). The `release` service is a compiled version of the Go src and a lightweigh distro intended for production environments.
