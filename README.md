# Go `gqlgen` Federated Subgraph

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/new/template/NMBWQ1?referralCode=xsbY2R)

This is an example application template that can be used to create Federated GraphQL subgraph using [gqlgen](https://gqlgen.com/getting-started/). You can use this template from [Rover](https://www.apollographql.com/docs/rover/commands/template/) with `rover template use --template subgraph-go-gqlgen`.

This example application implements following GraphQL schema:

```graphql
directive @contact(
    "Contact title of the subgraph owner"
    name: String!
    "URL where the subgraph's owner can be reached"
    url: String
    "Other relevant notes can be included here; supports markdown links"
    description: String
) on SCHEMA

schema
@contact(
    name: "FooBar Server Team"
    url: "https://myteam.slack.com/archives/teams-chat-room-url"
    description: "send urgent issues to [#oncall](https://yourteam.slack.com/archives/oncall)."
)
@link(
    url: "https://specs.apollo.dev/federation/v2.0",
    import: ["@key"]
) {
    query: Query
}

type Query {
    foo(id: ID!): Foo
}
type Foo @key(fields: "id") {
    id: ID!
    name: String
}
```

## Build

In order to build the project locally run the following go command

```shell
go build
```

### Updating Schema

gqlgen auto generates the code for you based on your schema. If you update schema, regenerate your code by running

```shell
# download and install gqlgen locally, only need to run it once
go get -d github.com/99designs/gqlgen
# regenerate code
go run github.com/99designs/gqlgen generate
```

### Code Quality

Example integration test is provided. It starts up the example server and executes `foo` query against it. Run `test` command to execute provided tests.

```shell
go test
```

### Continuous Integration

This project comes with some example build actions that will trigger on PR requests and commits to the main branch.

## Run

To start the GraphQL server run following go command.

```shell script
go run server.go
```

Once the app has started you can explore the example schema with [Apollo Studio Sandbox](https://www.apollographql.com/docs/graphos/explorer/sandbox) and begin developing your supergraph with `rover dev --url http://localhost:8080 --name my-subgraph`.

## Apollo Studio Integration

1. Set these secrets in GitHub Actions:
    1. APOLLO_KEY: An Apollo Studio API key for the supergraph to enable schema checks and publishing of the subgraph.
    2. APOLLO_GRAPH_REF: The name of the supergraph in Apollo Studio.
    3. PRODUCTION_URL: The URL of the deployed subgraph that the supergraph gateway will route to.
2. Set SUBGRAPH_NAME in .github/workflows/checks.yaml and .github/workflows/deploy.yaml
3. Remove the if: false lines from .github/workflows/checks.yaml and .github/workflows/deploy.yaml to enable schema checks and publishing.
4. Write your custom deploy logic in .github/workflows/deploy.yaml.
5. To secure your subgraph, send the `Router-Authorization` header [from your Cloud router](https://www.apollographql.com/docs/graphos/routing/cloud-configuration#managing-secrets) and set the `ROUTER_SECRET` environment variable wherever you deploy this to.

## Additional Resources

* [gqlgen documentation](https://gqlgen.com/getting-started/)
* [Golang documentation](https://go.dev/doc/)
