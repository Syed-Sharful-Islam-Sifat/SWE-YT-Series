#!/usr/bin/env bash
# Calls each endpoint and prints a comparison table.
#   ./bench.sh                     -> default endpoints, limits 10 50 200
#   ./bench.sh "/v1/naive /v3/batch" "50 500"
set -euo pipefail

BASE=${BASE:-http://localhost:8080}
ENDPOINTS=${1:-"/v1/naive /v2/join /v3/batch /v4/json"}
LIMITS=${2:-"10 50 200"}

printf "%-22s %6s %8s %8s %10s\n" endpoint limit queries rows ms
for ep in $ENDPOINTS; do
  for l in $LIMITS; do
    curl -s -o /dev/null "$BASE$ep?limit=$l"   # warm-up
    h=$(curl -s -o /dev/null -D - "$BASE$ep?limit=$l")
    q=$(echo "$h" | grep -i '^x-query-count' | tr -d '\r' | awk '{print $2}')
    r=$(echo "$h" | grep -i '^x-rows'        | tr -d '\r' | awk '{print $2}')
    d=$(echo "$h" | grep -i '^x-duration-ms' | tr -d '\r' | awk '{print $2}')
    printf "%-22s %6s %8s %8s %10s\n" "$ep" "$l" "$q" "$r" "$d"
  done
done
