# Task 02: Prologue – The Logbook of the Grand Line
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
### Writeup :
I pretty much had to tell gemini to strip away the one piece references. Tbh i'm pretty sure even a one piece fan would find it hard to decipher whatever was written on the task readme. "Somewhere among the cargo lies an object that still possesses the power to awaken itself." I mean seriously? What did that even mean 😭. Anyway once I got to know it was about finding file permissions the rest was easy.
![](pics/1.0.png)
![](pics/1.1.png)

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
### Writeup :
This one was pretty easy. Found the shell script and tried to run it. Hint was enough to solve the rest of it.
![](pics/2.0.png)
![](pics/2.1.png)
![](pics/2.2.png)
![](pics/2.3.png)
![](pics/2.4.png)

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
### Writeup :
Gemini told me it was prolly some encrypted thing. Luckily it was base64 and I got it right my first try. Opened the file the grep command returned and voila.
![](pics/3.0.png)

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
### Writeup :
Didn't need gemini for this. Level description was enough. Tried opening the only file there was. Got a bunch of random shit. Searched up how to retrieve data from corrupted files and found that the linux file command can do it regardless of extension. It worked. Renamed the file to the name the file command returned. Used the tar command to extract the tar file. Used the unzip module to unzip the file inside the tar file. Searched through the extracted zip file and got the second glyph.
![](pics/4.0.png)
![](pics/4.1.png)

## Level 5
### A brief summary :
I couldnt run the python file for some reason so i used cat to read the file and decode the glyph myself
### Tools used :
https://www.online-python.com
### Result :
Prize : 
https://github.com/rogueone-x/Laugh-Tale-Merge-War
### Writeup :
listed through directories until i found the python file. Couldnt run it for some reason so i ```cat``` the file and ran it through an online compiler. Modern problems require modern solutions. Found the last level link.
![](pics/5.0.png)
![](pics/5.1.png)

## Level 6
### Writeup :
(accidently) Saw that there were 2 branches while cloning the repo. listed through all directories. Thought key_part_1 and 2 looked suspicious. Entered password as Linebers 😂. Then switched branches and saw the rest of it. Entered the password and won the challenge. Also disclaimer. I forgot to take screenshots the first time I solved it so i had to do it all over again 🥲.
![](pics/6.0.png)
![](pics/6.1.png)