# LogBook
## Level 1
### command used : 
ls -la sector_A sector_B sector_C sector_D
### what it does : 
- ls = lists file
- -a = includes hidden files
- -l = lists details of files too
- -la = get all the details of the hidden files (basically a combo of -a and -l)
### Resource Used:
https://docs.rackspace.com/docs/checking-linux-file-permissions-with-ls
### AWAKENING_SIGNATURE:

ONE_PIECE{GITO_GITO_NO_AWAKENING}
## Level 2
### commands used :
- git checkout whiskey_peak_investigation
- export AWAKENING_SIGNATURE="ONE_PIECE{GITO_GITO_NO_AWAKENING}"
- ./unlock_vault.sh
### what they do :
- checkout : changes branches
- export : sets variable
### Result :
BAROQUE_DIAL{SPLIT_TIMELINE_MISDIRECTION}
## Level 3
### commands used :
- grep -r "QkFST1FVRV9ESUFMe1NQTElUX1RJTUVMSU5FX01JU0RJUkVDVElPTn0K"
### what it does :
- grep : searches through files
- -r : does it recursively
- QkFST1FVRV9ESUFMe1NQTElUX1RJTUVMSU5FX01JU0RJUkVDVElPTn0K" : base64 encode of BAROQUE_DIAL{SPLIT_TIMELINE_MISDIRECTION}
### Result :
```SECURITY LOG ACCESS // LEVEL 3 CLEARANCE REQUIRED
-------------------------------------------------
STATUS: METALLIC WAX SUIT ACTIVE

SECURITY_TAG:
QkFST1FVRV9ESUFMe1NQTElUX1RJTUVMSU5FX01JU0RJUkVDVElPTn0K

-------------------------------------------------

BAROQUE WORKS EXECUTIVE REPORT

PONEGLYPH_FRAGMENT_I = "KjY2MjF4bW0lKzYqNyBsIS0vbTAtJTcnL"

-------------------------------------------------
```
## Level 4
### commands used :
- file puffing_tom_blueprints
- mv puffing_tom_blueprints step2_blueprints.tar
- tar -xvf step2_blueprints.tar
- unzip step1_blueprints.zip
### what they do :
- file : get the file type even without any extensions
- mv : used to rename files
- tar -xvf: used to unzip tar files (-x : Extract a tar ball, -v : Verbose output or show progress while extracting files. -f : Specify an archive or a tarball filename.)
- unzip : unzip .zip files (have to be installed using sudo apt)
### Resources Used :
- https://phoenixnap.com/kb/linux-file-command
- https://www.cyberciti.biz/faq/tar-extract-linux/
### Result :
PONEGLYPH_FRAGMENT_II="SwnbzptDiM3JSpvFiMuJ28PJzAlJ28VIzA="
## Level 5
### A brief summary :
I couldnt run the python file for some reason so i used cat to read the file and decode the glyph myself
### Tools used :
https://www.online-python.com
### Result :
Prize : 
https://github.com/rogueone-x/Laugh-Tale-Merge-War
## Level 6
### A brief summary :
Ahh finally. At first i thought the password was Linebers. That didnt work so i looked at the other branch and got the rest of it. the Pirate King's Password: TheGrandLineRemembers
FLAG{The_Grand_Line_Remembers_Your_Commit}