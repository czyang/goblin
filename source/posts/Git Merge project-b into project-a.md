{
    "title": "Git Merge Multiple Projects into One",
    "author": "Chengzhi Yang",
    "createDate": "2016-12-01",
    "modifyDate": "2016-12-01",
    "permanent":"git-merge-multiple-projects-into-one"
}

# Git Merge Multiple Projects into One

The steps list below is merge project-b into project-a, follow and repeat these steps you can merge up projects into one repository.

```sh
cd path/to/project-a
git remote add project-b https://url/to/project-b.git
git fetch project-b
git merge —allow-unrelated-histories project-b/master
git remote remove project-b
```

Now project-b's files will place at project-a's folder, if you want move all project-b's files into a independent folder you need extra work.

Assume project-a is a empty project before we merge it with project-b. We can simply move all files exclude .git folder to project-b's folder.

```sh
mkdir project-b-folder
git mv !(.git|project-b-folder) project-b-folder
git commit -m "Move project-b files to its folder."
```

Reference:
http://stackoverflow.com/a/10548919/522915
