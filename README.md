# hasura-n-plus-one-remote-rel

This demo shows the N+1 problem in Hasura when using Remote Relationships (“remote joins”).

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

## Using the API

### Creating some data

```graphql
mutation CreateOwnerAndPet {
  owner: insertOwnerOne(
    object: {
      name: "Kevin"
      pets: {
        data: {
          name: "Porkchop"
        }
      }
    }
  ) {
    id
    name
    pets {
      id
      name
    }
  }
}
```

[atlas]: https://atlasgo.io/
