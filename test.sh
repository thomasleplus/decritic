#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

bin="$(go env GOPATH)/bin/decritic"
rc=0
while IFS=$'\n' read -r line; do
	IFS=',' read -ra array <<<"${line}"
	input="${array[0]}"
	expected="${array[1]}"
	actual="$(echo "${input}" | ${bin})"
	if [ "${actual}" != "${expected}" ]; then
		>&2 echo "input=${input}, expected=${expected}, actual=${actual}"
		rc=$((rc + 1))
	fi
done
exit ${rc}
