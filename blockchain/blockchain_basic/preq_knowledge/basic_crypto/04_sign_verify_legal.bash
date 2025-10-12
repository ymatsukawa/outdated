ROOT_DIR=`pwd`
WORKING_DIR=`dirname $0`
CURRENT_DIR=$ROOT_DIR/$WORKING_DIR

PRIVATE_KEY=$CURRENT_DIR/secp256k1-private.pem
PUBLIC_KEY=$CURRENT_DIR/secp256k1-public.pem
DATA_FILE=$CURRENT_DIR/message.dat
SIGN_FILE=$CURRENT_DIR/message.sig

# -------------------------------

# create private key and public key from the private key
openssl ecparam -genkey -name secp256k1 -out $PRIVATE_KEY
openssl ec -in $PRIVATE_KEY -pubout -out $PUBLIC_KEY

# create txt file and sign with private key
echo "this is secret message" > $DATA_FILE
openssl dgst -SHA256 -sign $PRIVATE_KEY $DATA_FILE > $SIGN_FILE

# verify data by LEGAL sign(public key)
openssl dgst -SHA256 -verify $PUBLIC_KEY -signature $SIGN_FILE $DATA_FILE

# -------------------------------

PRIVATE_KEY_ATTACKER=$CURRENT_DIR/secp256k1-private_attacker.pem
PUBLIC_KEY_ATTACKER=$CURRENT_DIR/secp256k1-public_attacker.pem
DATA_FILE_MANIPULATED=$CURRENT_DIR/message_manipulated.dat
SIGN_FILE_ATTACKER=$CURRENT_DIR/message_manipulated.sig

# create attacker's private key and public key
openssl ecparam -genkey -name secp256k1 -out $PRIVATE_KEY_ATTACKER
openssl ec -in $PRIVATE_KEY_ATTACKER -pubout -out $PUBLIC_KEY_ATTACKER

# regard this operation is to manipulate $DATA_FILE
echo "I see your secret message" > $DATA_FILE_MANIPULATED
# and create ILLEGAN sign
openssl dgst -SHA256 -sign $PRIVATE_KEY_ATTACKER $DATA_FILE_MANIPULATED > $SIGN_FILE_ATTACKER

# when verify 2 patterns
# * LEGAL sign but ILLEGAL data, verification fails
# data must be manipulated by someone.
openssl dgst -SHA256 -verify $PUBLIC_KEY -signature $SIGN_FILE $DATA_FILE_MANIPULATED
# * LEGAL data but ILLEGAL sign, verification fails
# sender of data must not have LEGAL private key.
openssl dgst -SHA256 -verify $PUBLIC_KEY -signature $SIGN_FILE_ATTACKER $DATA_FILE