# Check if the required environment variable is set
ifndef HASURA_GRAPHQL_PRO_KEY
$(error Environment variable HASURA_GRAPHQL_PRO_KEY is not set. Please set it before running make.)
endif

.PHONY: all
all:
	pkgx task start