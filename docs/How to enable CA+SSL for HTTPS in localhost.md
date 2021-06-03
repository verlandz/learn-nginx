> **[DISCLAIMER]**\
> I got the guide from [here](https://stackoverflow.com/a/60516812) and modify a bit for this project.\
> Credit to respective owner.

# Files
list of files that you will get
```
/etc/ssl/private/mynginx_ca.key
/etc/ssl/certs/mynginx_ca.pem
/etc/ssl/private/mynginx.key
/etc/ssl/certs/mynginx.crt
/etc/ssl/certs/mynginx.pem
/etc/ssl/certs/mynginx.csr
/etc/ssl/ext/mynginx.ext
```

# Step
in terminal
```
####################################
## Become a Certificate Authority ##
####################################
# Generate private key
openssl genrsa -des3 -out /etc/ssl/private/mynginx_ca.key 2048

# Generate root certificate
openssl req -x509 -new -nodes -key /etc/ssl/private/mynginx_ca.key -sha256 -days 825 -out /etc/ssl/certs/mynginx_ca.pem

##########################
# Create CA-signed certs #
##########################
# Generate a private key
openssl genrsa -out /etc/ssl/private/mynginx.key 2048

# Create a certificate-signing request
openssl req -new -key /etc/ssl/private/mynginx.key -out /etc/ssl/certs/mynginx.csr

# Create a config file for the extensions
mkdir /etc/ssl/ext; >/etc/ssl/ext/mynginx.ext cat <<-EOF
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names
[alt_names]
DNS.1 = mynginx.com
DNS.2 = www.mynginx.com
EOF

# Create the signed certificate
openssl x509 -req -in /etc/ssl/certs/mynginx.csr -CA /etc/ssl/certs/mynginx_ca.pem -CAkey /etc/ssl/private/mynginx_ca.key -CAcreateserial \
-out /etc/ssl/certs/mynginx.crt -days 825 -sha256 -extfile /etc/ssl/ext/mynginx.ext
```

in browser, ex: Google Chrome v86
```
Setting > Security > Manage Certificates > Authorities > Import
choose file "/etc/ssl/certs/mynginx_ca.pem"
```

in nginx file
```
ssl_certificate     /etc/ssl/certs/mynginx.crt;
ssl_certificate_key /etc/ssl/private/mynginx.key;
```

then `service nginx restart`


# Sample
this is sample of additional required information. You can modify by yourself.

### Password
I used `admin` to all password.

### Generate root certificate
```
Country Name (2 letter code) [AU]:ID
State or Province Name (full name) [Some-State]:My State
Locality Name (eg, city) []:
Organization Name (eg, company) [Internet Widgits Pty Ltd]:My Company
Organizational Unit Name (eg, section) []:
Common Name (e.g. server FQDN or YOUR name) []:*.mynginx.com
Email Address []:
```

### Create a certificate-signing request
```
Country Name (2 letter code) [AU]:ID
State or Province Name (full name) [Some-State]:My State
Locality Name (eg, city) []:
Organization Name (eg, company) [Internet Widgits Pty Ltd]:My Company
Organizational Unit Name (eg, section) []:
Common Name (e.g. server FQDN or YOUR name) []:*.mynginx.com
Email Address []:

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:admin
An optional company name []: 
```
