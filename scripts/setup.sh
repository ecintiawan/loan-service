#!/usr/bin/env bash

echo -e "\e[32mRunning:\e[33m setup.\e[0m\n"

echo -e "\e[32mInstalling:\e[33m mockgen for mock generator.\e[0m"
command -v mockgen 2>/dev/null || go install github.com/golang/mock/mockgen@v1.6.0
echo ""

echo -e "\e[32mInstalling:\e[33m wire for wire generator.\e[0m"
command -v wire 2>/dev/null || go install github.com/google/wire/cmd/wire@v0.6.0
echo ""

echo -e "\e[32mSetup:\e[33m pre-commit hook.\e[0m"
file=.git/hooks/pre-commit
cp scripts/pre-commit.sh $file
chmod +x $file
test -f $file && echo "$file exists."
echo ""

echo -e "\e[32mSetup:\e[33m copy template credential config.\e[0m"
dir="files/etc/credential/development"
[ ! -d "$dir" ] && mkdir -p "$dir"
cp files/etc/credential/loan-service.secret.json.sample files/etc/credential/development/loan-service.secret.json

echo -e "\e[32mSetup:\e[33m success.\e[0m"
