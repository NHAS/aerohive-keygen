# Aerohive Key Generator
Aerohive access points enable users to configure devices using a restricted command line shell.  
There is an undocumented restricted shell command `_shell` which spawns `/bin/sh` if given the correct password. 

This tool generates that password for the `AP230` and `AP130`, as those are the devices I have on hand.

# Usage

```
go build .
./aerohive-keygen --serial <your devices serial number> --version <your firmware version>
```
Enter into the restricted shell, and type `_shell`, enter in the generated password.  

## Todo
- ~~Add support for generating AP130 keys~~
- ~~Detect platform from serial number~~
- ~~Add version parameter for AP230~~




