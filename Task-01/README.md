# Task - 01 : Git Exercises
- ```git add file.txt```: stages file.txt for commit. ```add .``` stages every changed file.
- ```git commit -m "hello"```: A sort of "checkpoint" for your files. -m defines the message(in this case "hello"). Add file name to commit only one file.
- ```git push```: Basically commit but to a remote network. Commit is for our machine, push is for the git repository on the internet.
- ```touch .gitignore```: touch command creates file. I used it to make the .gitignore file for "ignore-them"
- ```nano .gitignore```: opens in-terminal editor for file.
- ```git merge branch_name```:merges ```branch_name``` to current branch
- ```git status```: tells which branch we are on. And pending commits. Not really used with a specific level in mind just to know where i was as i was getting a few errors with too many commits/files.
- ```git stash```: used to create a temporary save point. As fracz said "You can think of stash as an intelligent Git clipboard".
- ```git stash pop```: takes your temporary save point and puts it right back
- ```git rebase branch_name```:pretty much a linear merge. normal merge makes it look like the 2 branches were worked on simultaneously. Rebase rewrites everything and makes it look linear from the common commit. Merge is still more accurate but rebase looks cleaner i guess.
- ```git rm file.txt```:removes file from repo.
- ```git mv old.txt new.txt```: renames old.txt to new.txt
- ```git commit --amend```:edits last commit
- ```git commit --amend --no-edit```: change only commited files but no edit message
- ```git commit --amend --no-edit --date="1987-08-03"```: overrides timestamp of commit too
- ```git rebase -i HEAD~3^```: Enables editing last 3 commits. basically a more powerful ```commit --amend```.
- ```git rebase --continue```: I guess thats for continuing the commit edit chain. Idk git told me to run it.
- ```hhh```



git clone https://gitexercises.fracz.com/git/exercises.git
cd exercises
git config user.name "Rajiv"
git config user.email "rajiv.a.r@outlook.com"
./configure.sh
git start