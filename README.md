# LRP
ランチリコメンドプロジェクト


## 環境設定
`#  export GOPATH=$HOME/go`
## コンパイルの仕方
`#  /usr/local/go/bin/go install lrp/`


#GAEの使いかた
```
8  TUTORIALDIR=~/src/lunchrecommendproject-175311/go_gae_quickstart-2017-08-01-16-14
9  git clone https://github.com/GoogleCloudPlatform/appengine-guestbook-go.git $TUTORIALDIR
10  cd $TUTORIALDIR
11  ls -ltr
12  git checkout part1-helloworld
13  cat hello.go
14  ls -ltr
15  cat app.yaml
16  goapp serve app.yaml
17  goapp deploy -application lunchrecommendproject-175311 -version 0
```
