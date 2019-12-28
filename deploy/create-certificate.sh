if [ -d "./certificates" ]; then
  rm -drf ./certificates
fi

mkdir -p certificates

# see https://stackoverflow.com/questions/10175812/how-to-create-a-self-signed-certificate-with-openssl/27931596#27931596

openssl req -config openssl.cnf \
  -x509 -nodes -days 365 \
  -subj "/C=DE/ST=Germany/L=Munich/O=JABASOFT/CN=JABASOFT" \
  -newkey rsa:2048 -keyout certificates/jabasoft-ds.key -out certificates/jabasoft-ds.crt

