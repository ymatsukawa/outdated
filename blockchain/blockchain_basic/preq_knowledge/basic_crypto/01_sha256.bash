cat <<-EOL
sha-256 encrypt
echo "when input is same, output is also same
echo "when input is different, output is also different
EOL

echo "abc" | sha256sum
echo "abc" | sha256sum # same result value of above
echo "abcd" | sha256sum # different from result of above