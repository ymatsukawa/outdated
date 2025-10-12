cat <<-EOL
Create private key by openssl.
Generation method is secp256k1.
-----------------
EOL

ROOT_DIR=`pwd`
WORKING_DIR=`dirname $0`
CURRENT_DIR=$ROOT_DIR/$WORKING_DIR

echo "1) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"
echo "create private key."
openssl ecparam -genkey -name secp256k1 -out $CURRENT_DIR/secp256k1-private.pem
PRIVATE_KEY=$CURRENT_DIR/secp256k1-private.pem
cat $PRIVATE_KEY
echo "1) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"


echo "2) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"
echo "regenerate private key as hex value."
# 1st pipeline meaning: create elliptic curve key processing with $PRIVATE_KEY and output binary
openssl ec -in $PRIVATE_KEY -outform DER | tail -c +8 | head -c 32 | xxd -p -c 32 > $CURRENT_DIR/hex_private.pem
cat $CURRENT_DIR/hex_private.pem
echo "2) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"


echo "3) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"
echo "create public key from private key"
openssl ec -in $PRIVATE_KEY -pubout -out $CURRENT_DIR/secp256k1-public.pem
echo "3) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"


echo "4) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"
echo "regenerate public key as hex value."
openssl ec -in $PRIVATE_KEY -pubout -outform DER | tail -c 65 | xxd -p -c 65 > $CURRENT_DIR/hex_pub.pem
cat $CURRENT_DIR/hex_pub.pem
echo "4) >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"