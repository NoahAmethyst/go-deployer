 solc --abi ./contract/cat/*.sol -o  ./contract/cat/
 abigen --abi=./contract/cat/CAT.abi --pkg=store --out=./contract/cat/cat.go
 solc --bin ./contract/cat/*.sol -o ./contract/cat
 abigen --bin=./contract/cat/CAT.bin --abi=./contract/cat/CAT.abi --pkg=store --out=./contract/cat/cat.go
 abigen --bin=./contract/cat/Minter.bin --abi=./contract/cat/CAT.abi --pkg=store --out=./contract/cat/cat.go

# minter

 solc --abi ./contract/cat_minter/*.sol -o  ./contract/cat_minter/
 abigen --abi=./contract/cat_minter/Minter.abi --pkg=store --out=./contract/cat_minter/minter.go
 solc --bin ./contract/cat_minter/*.sol -o ./contract/cat_minter
 abigen --bin=./contract/cat_minter/Minter.bin --abi=./contract/cat/Minter.abi --pkg=store --out=./contract/cat_minter/minter.go
 abigen --bin=./contract/cat_minter/Minter.bin --abi=./contract/cat_minter/Minter.abi --pkg=store --out=./contract/cat_minter/Minter.go