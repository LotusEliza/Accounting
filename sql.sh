#!/bin/bash
PROJECT=account-dev
docker exec -it $PROJECT ./cockroach sql --insecure
