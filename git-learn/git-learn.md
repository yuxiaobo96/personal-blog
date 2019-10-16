---
title: git-learn
---
### 回退版本

1. 当在工作区修改文件，还没`git add .`时：
git checkout .
(与git add .相对)或者git reset --hard
2. 当已经 `git add .`,保存到暂存区时：
git reset(退回到第1）步骤) + git checkout .（回退到未修改状态）
或者 git reset --hard
3. 当已经 `git commit`，保存到本地仓库时：
git reset --hard origin/master 从远程仓库重新拉代码（远程仓库主分支：origin/master）
4. 当已经 `git push`，推送到远程仓库时：
先回退版本（即回退到第1）步骤之前，当时还未修改文件）
git reset --hard HEAD^（HEAD代表当前版本，^代表上一版本，也可用数字进行代替，如:HEAD~1）
也可以通过指定版本进行回退：git reset --hard +版本号（前7位就好）
5. 再重新进行修改文件
工作区里update-git add到暂存区-git commit到本地仓库
最后需强制推送到远程仓库
git push -f origin +分支名

{{< note >}}

1. git log用来记录每一次的commit提交记录
2. git reflog用来记录每一次git命令记录（包括checkout,merge,commit等）

{{< note >}}

### 合并分之及解决冲突

1. 合并分之
可在当前分之上使用：git merge +其他分之
2. 解决冲突
在当前分之上解决冲突；
去到提示的文件中git所标注的版本冲突处进行修改；
把修改后的冲突文件进行git add、git commit 操作；
最后再git push -f origin +当前分之（也可以不强推）

{{< note >}}

git log --graph用来记录分之的合并图

{{< note >}}

### 在clone下来的其他分之上进行操作（如release-1.14）

1. 当clone下来一个项目时，只能看到一个master主分支；
2. 使用`git branch -a` 可以看到所有分之；
3. 此时需要创建远程origin的release-1.14分支到本地仓库上：
         git checkout -b release-1.14 origin/release-1.14
（或者先切换到要进行操作的远程分支上git checkout remotes/origin/release-1.14；然后在此分之上创建与远程分之对应的本地分之 git checkout -b release-1.14）

### 将同一分之上已经提交的两条或多条commit信息合并（commit2和commit3）

1. 利用`git log`查看commit版本（最新的版本在最上面），
找到要合并commit版本的前一条的commit信息(commit1)；
2. 使用git rebase -i + commit1版本号  (-i参数指不需要合并commit的hash值)；
进入vi的编辑模式，上方未注释的部分是要执行的指令，下方注释的部分是指令的提示说明；
上方指令部分由三部分组成：命令名称、commit hash、commit message，
其中命令名称有两种：pick-执行该条commit；
                                             squash-把当前commit合并到前一个commit;
3. 将后一个commit的前方命令`pick`改成`sqash`或`s`,保存退出；
进入到commit message的编辑界面，将提示说明前（即#Pleasesa）的几行文字删除，增添新的commit信息即可，完成commit信息合并。

{{< note >}}

1. git rebase --continue 继续运行（用于解决冲突后，使操作继续执行）
2. git rebase --abort 回退到变基前，即恢复所有的修改

{{< note >}}

### 保存当前分之修改，然后去修改其他分之（如当前分之修改到一半，急需去优先修改另一分之）

1. 保存当前分之修改：
git slash
2. 切换到另一分之进行搞作
3. 切换到1）步骤的分之，恢复并删除slash信息：
git stash pop
4. 继续该分之的剩余修改

### 在当前分之上复制其他分之上的部分修改（比如某一模块中的bug）

1. 首先使用`git log`找到其他分之此bug修改的版本（commit信息）
2. 在当前分之：git cherry-pick +需要复制的commit版本号（前7位就好）

### 为git中常用的命令增加别名

1. 使用命令：
git config --global alias.别名 + 要修改的命令
如：git config --global alias.st status
         git config --global alias.co checkout
2. 配置Git的时候，加上`--global`是针对当前用户起作用的，如果不加，那只针对当前的仓库起作用。
    每个仓库的Git配置文件都放在`.git/config`文件中：

```bash
$personal-blog git:(master) cat .git/config
[core]
	repositoryformatversion = 0
	filemode = true
	bare = false
	logallrefupdates = true
[remote "origin"]
	url = https://github.com/yuxiaobo96/personal-blog.git
	fetch = +refs/heads/*:refs/remotes/origin/*
[branch "master"]
	remote = origin
	merge = refs/heads/master

```

而当前用户的Git配置文件放在用户主目录下的一个隐藏文件`.gitconfig`中：

```bash
$ ~ cat .gitconfig
[user]
	name = yuxiaobo
	email = yuxiaobogo@163.com
[alias]
	st = status
```

{{< note >}}

若配置错误，可在这两个文件下编辑修改或删除。

{{< note >}}

增加别名
配置~/.gitconfig