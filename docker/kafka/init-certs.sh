#!/bin/bash
set -e

FOLDER="/etc/kafka/secrets"

mkdir -p "$FOLDER"

# if the keystore already exists, do not recreate
if [[ -f "$FOLDER/kafka.keystore.jks" ]]; then
  echo "🔒 Certificates already exist, skipping generation..."
  exit 0
fi

echo "🔑 Generating certificates and keystore..."

# --- CA ---
openssl genrsa -out "$FOLDER/ca.key" 4096
openssl req -x509 -new -nodes -key "$FOLDER/ca.key" -sha256 -days 3650 \
  -out "$FOLDER/ca.crt" -subj "/CN=Local Dev CA"

# --- Broker key e CSR ---
openssl genrsa -out "$FOLDER/kafka.localhost.key" 2048
openssl req -new -key "$FOLDER/kafka.localhost.key" -out "$FOLDER/kafka.localhost.csr" \
  -subj "/CN=*.localhost"

# SAN wildcard
echo "subjectAltName=DNS:localhost,DNS:*.localhost,DNS:kafka" > "$FOLDER/san.cnf"

# Sign CSR with CA
openssl x509 -req -in "$FOLDER/kafka.localhost.csr" -CA "$FOLDER/ca.crt" -CAkey "$FOLDER/ca.key" \
  -CAcreateserial -out "$FOLDER/kafka.localhost.crt" -days 365 -sha256 -extfile "$FOLDER/san.cnf"

# --- Keystore ---
keytool -genkey -alias kafka -keyalg RSA -keystore "$FOLDER/kafka.keystore.jks" \
  -storepass changeit -keypass changeit -dname "CN=*.localhost" -noprompt

# Export certificate + key to PKCS12
openssl pkcs12 -export -in "$FOLDER/kafka.localhost.crt" -inkey "$FOLDER/kafka.localhost.key" \
  -name kafka -password pass:changeit -out "$FOLDER/kafka.p12"

# Import PKCS12 to keystore
keytool -importkeystore -deststorepass changeit -destkeypass changeit \
  -destkeystore "$FOLDER/kafka.keystore.jks" -srckeystore "$FOLDER/kafka.p12" \
  -srcstoretype PKCS12 -srcstorepass changeit -alias kafka -noprompt

# --- Truststore ---
keytool -import -trustcacerts -alias CARoot -file "$FOLDER/ca.crt" \
  -keystore "$FOLDER/kafka.truststore.jks" -storepass changeit -noprompt

# --- Credentials ---
echo "changeit" > "$FOLDER/keystore_creds"
echo "changeit" > "$FOLDER/truststore_creds"
echo "changeit" > "$FOLDER/key_creds"

# --- JAAS ---
cat > "$FOLDER/kafka_server_jaas.conf" <<'EOF'
KafkaServer {
  org.apache.kafka.common.security.scram.ScramLoginModule required;
};
EOF

echo "✅ Certificates generated in $FOLDER"
