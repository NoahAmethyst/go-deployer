# solc --abi ./contract/bcat/*.sol -o  ./contract/bcat/
# abigen --abi=./contract/bcat/BCAT.abi --pkg=store --out=./contract/bcat/bcat.go
# solc --bin ./contract/bcat/*.sol -o ./contract/bcat
# abigen --bin=./contract/bcat/BCAT.bin --abi=./contract/bcat/BCAT.abi --pkg=store --out=./contract/bcat/bcat.go
# abigen --bin=./contract/bcat/Minter.bin --abi=./contract/bcat/BCAT.abi --pkg=store --out=./contract/bcat/bcat.go

# minter

# solc --abi ./contract/bcat_minter/*.sol -o  ./contract/bcat_minter/
# abigen --abi=./contract/bcat_minter/BCatMinter.abi --pkg=store --out=./contract/bcat_minter/minter.go
# solc --bin ./contract/bcat_minter/*.sol -o ./contract/bcat_minter
# abigen --bin=./contract/bcat_minter/BCatMinter.bin --abi=./contract/bcat/BCatMinter.abi --pkg=store --out=./contract/bcat_minter/minter.go
# abigen --bin=./contract/bcat_minter/BCatMinter.bin --abi=./contract/bcat_minter/BCatMinter.abi --pkg=store --out=./contract/bcat_minter/Minter.go

# origin bcat
 solc --abi ./contract/bcatv4/*.sol -o  ./contract/bcatv4/
 abigen --abi=./contract/bcatv4/BCATV4.abi --pkg=store --out=./contract/bcatv4/bcat.go
 solc --bin ./contract/bcatv4/*.sol -o ./contract/bcatv4
 abigen --bin=./contract/bcatv4/BCATV4.bin --abi=./contract/bcatv4/BCATV4.abi --pkg=store --out=./contract/bcatv4/bcat.go


# data storage
# solc --abi ./contract/data/*.sol -o  ./contract/data/
# abigen --abi=./contract/data/DataStorage.abi --pkg=store --out=./contract/data/DataStorage.go
# solc --bin ./contract/data/*.sol -o ./contract/data
# abigen --bin=./contract/data/DataStorage.bin --abi=./contract/data/DataStorage.abi --pkg=store --out=./contract/data/DataStorage.go
