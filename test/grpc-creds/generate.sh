#!/bin/bash
set -e
set -o xtrace

cd "$(realpath -- $(dirname -- "$0"))"

# Check that `minica` is installed
command -v minica >/dev/null 2>&1 || {
  echo >&2 "No 'minica' command available.";
  echo >&2 "Check your GOPATH and run: 'go get github.com/jsha/minica'.";
  exit 1;
}

for SERVICE in admin-revoker expiration-mailer ocsp-responder consul \
  wfe akamai-purger bad-key-revoker crl-updater crl-storer \
  health-checker; do
  minica -domains "${SERVICE}.boulder"
done

for SERVICE in publisher nonce ra ca sa va rva ; do
  minica -domains "${SERVICE}.boulder,${SERVICE}1.boulder,${SERVICE}2.boulder"
done

minica -ip-addresses 10.77.77.77,10.88.88.88

# grpc/creds/creds.go:
minica -domains "creds-test" -ip-addresses "127.0.0.1"

# minica sets restrictive directory permissions, but we don't want that
chmod -R go+rX .
