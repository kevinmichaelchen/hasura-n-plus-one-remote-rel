# hasura-n-plus-one-remote-rel

This demo shows the N+1 problem in Hasura when using Remote Relationships
(“remote joins”).

## Getting started

### Step 0: Prerequisites

You will need Docker.

You will also need [pkgx](https://pkgx.sh/):

```shell
sudo rm -rf $(which pkgx) ; curl -fsS https://pkgx.sh | sh
```

### Step 1: Run everything

```shell
pkgx task start
```

### Step 2: Create some data

```graphql
mutation SeedData {
  owners: insertOwner(
    objects: [
      { name: "Gravy", pets: { data: { name: "Porkchop" } } }
      { name: "Blueberry", pets: { data: { name: "Oatmeal" } } }
    ]
  ) {
    returning {
      id
      name
      pets {
        id
        name
      }
    }
  }
}
```

### Step 3: Query remote schemas

```graphql
query GetStuff {
  owners: owner {
    id
    nickname
    pets {
      id
      nickname
    }
  }
}
```

## What's the problem?

Here, we are requesting multiple owners, and for each owner, we are requesting
their pet. Each `nickname` field is an HTTP call to our Nickname service. These
calls aren't batched.

This is the `N + 1` problem (or `1 + N` problem, if you prefer). Hasura will
make one call for the owners, and then N calls for each of their nicknames.

**There is no batching happening!**

## How to solve?

Hasura needs to support batching, similar to [how Tailcall does][tailcall].

[tailcall]: https://tailcall.run/docs/graphql-n-plus-one-problem-solved-tailcall/#batch-apis