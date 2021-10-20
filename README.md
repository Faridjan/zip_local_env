# Zip/Unzip local .env file
Transfer your local .ENV files for another developer.

## for Zip:
run in root where .env files
```
./ziper_local_env
```

created "result" file. It's all crypted .evn files

## for UnZip:
Copy crypted "result" file in root dir where .env files

and run:
```
./ziper_local_env
```


### Starting from console:
Add alis in profile ~/.bashrc

For ZSH console ~/.zshrc

```
alias <YOUR_ALIAS>='<YOUR_PATH_TO_THE_BIN_FILE>'
```

and apply alias:
```
source ~/.bashrc
```
