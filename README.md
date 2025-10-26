# Aerohive Key Generator
Aerohive access points enable users to configure devices using a restricted command line shell.  
There is an undocumented restricted shell command `_shell` which spawns `/bin/sh` if given the correct password.  

This tool generates that password for the `AP130`, `AP150` , `AP230` & `AP630`.  
Other models may apply, which can be tested by extending the case statement in main.go , rebuild and run.  
For the `AP150` and above you will need to add the version with parameter --version  

# Usage
```
go build .
./aerohive-keygen --serial <your devices serial number> --version <your firmware version>
```
Enter into the restricted shell, and type `_shell`, enter in the generated password.  

## Error : Unknown device type
If message returns the device type is unknown, the serial number is not matching a known model.  
Try adding the model in the AP230 case statement in main.go, rebuild and retry.  

## Todo
- ~~Add support for generating AP130 keys~~
- ~~Detect platform from serial number~~
- ~~Add version parameter for AP230~~




