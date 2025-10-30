# git_helper (CLI Tool)

## 概要

`git_helper` は、Git操作を補助するCLIツールです。  
以下の2つの実行方式に対応しています：

- **ランチャーモード**：対話式の選択UIで操作
- **ショートカットモード**：引数で直接操作を指定

現在対応しているコマンド一覧：

	•	git_pull           : 現在のブランチを pull
	•	git_fork           : 現在のブランチから新しいブランチを作成
	•	git_pr             : プルリクエストを作成
	•	git_commit         : コミット（プレフィックス補完付き）
	•	git_commit_redo    : コミットの取り消し
	•	git_push           : 現在のブランチを push
	•	git_fetch          : リモートブランチをフェッチ
	•	git_com            : コミットに使う説明文を事前登録

本リポジトリにはビルド済みバイナリは含まれておりません。  
以下のインストール手順に従ってビルドしてください。  
Go 1.24.4 以降が必要です。

## インストール方法

```bash
go build -o /usr/local/bin/git_helper main.go
```

上記のように PATH が通っているディレクトリへ出力することで、
どこからでも git_helper コマンドが使用可能になります。

## 使用方法

Gitリポジトリ内で以下のように実行します。

### ランチャーモード（対話式）

```bash
git_helper
```

矢印キーでコマンドを選択し、Enterキーで実行します。

### ショートカットモード（直接指定）

```bash
git_helper --mode=git_pull
```

コマンド名を直接指定することで、選択UIをスキップして即実行できます。

## より便利な使い方（alias設定）

.zshrc や .bashrc に以下のように記述することで、
コマンド名そのままで実行可能になります。

```bash
alias git_pull='git_helper --mode=git_pull'
alias git_fork='git_helper --mode=git_fork'
alias git_pr='git_helper --mode=git_pr'
alias git_commit='git_helper --mode=git_commit'
alias git_commit_redo='git_helper --mode=git_commit_redo'
alias git_push='git_helper --mode=git_push'
alias git_fetch='git_helper --mode=git_fetch'
alias git_com='git_helper --mode=git_com'
```

こちらを記述いただくことで、
モード名がそのままコマンドとなり大変便利です

## 謝辞

本ツールは、Go言語と多くの優れたOSSライブラリに支えられて開発されました。
ご自由にご利用・改変・再配布いただいて構いません。

なお、本ツールを使用したことによる損害・事故等について、作者は一切の責任を負いかねますのでご了承ください。
