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

### Step 1: Run Go service

```shell
go run main.go
```

### Step 2: Run everything else

```shell
pkgx task start
```

### Step 3: Create some data

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

### Step 4: Query remote schemas

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