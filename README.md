# Dovetail CLI

A Command line interface for the Dovetail API

## Getting Started


### Dependencies

* Git
* Go
* Windows 10 
* Mac OS

### Installing

* Download and install Git 
 - [Windows](https://git-scm.com/download/win)
 - [MacOS](https://git-scm.com/download/mac)
* Download and install Go. https://go.dev/doc/install

### Executing program
* Open a terminal window
* Create a new workspace folder
```
mkdir dovetail
cd dovetail
```
* Clone the repository into the folder
```
git clone https://github.com/zwalker8/dovetailCLI.git
```

* Change the .env.example file to .env
- MacOS / Linux
```
mv .env.example .env
```
- Windows
```
move .env.example .env
```

* Enter your api key
- MacOS / Linux
```
echo "API_KEY=\"YOUR KEY GOES HERE\"" > .env
```
- Windows 
```
echo "API_KEY=""YOUR KEY GOES HERE""" > .env
```
* Run the program 
```
go run $(pwd)/cmd/cli
```

## Authors

Zion Walker (zwalker8@jh.edu)


## Acknowledgments

* [dovetail API DOCS](https://developers.dovetail.com/docs/introduction)
