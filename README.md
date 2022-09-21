# Example Federated GQLGEN Subgraph

This is an example application template that can be used to create Federated GraphQL subgraph using [gqlgen](https://gqlgen.com/getting-started/).

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

Example integration test is provided. It starts up the the example server and executes `foo` query against it. Run
`test` command to execute provided tests.

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

Once the app has started you can explore the example schema by opening the GraphQL Playground endpoint at http://localhost:8080.

## Additional Resources

* [gqlgen documentation](https://gqlgen.com/getting-started/)
* [Golang documentation](https://go.dev/doc/)
