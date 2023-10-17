#!/bin/bash

chmod +x ./tools/executables/load-envs.sh
./tools/executables/load-envs.sh

if [[ schema != "" && migrations_path != "" ]]; then
  echo "Searching for migrations..."
  migrate -source file://$migrations_path -database $uri down
fi
