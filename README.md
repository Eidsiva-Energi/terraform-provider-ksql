# KSQL provider

## Test

To test the provider you run:

1. Start the kafka environment using docker compose.

```
docker-compose up
```

2. Run the tests.

```
cd ksql
TF_ACC=1 go test -count=1 -run='TestAccStreamResource' -v
```