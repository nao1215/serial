% SERIAL(1)
% Naohiro CHIKAMATSU <n.chika156@gmail.com>
% 2020年11月

# 名前

serial –  シリアル番号付きのファイル名にリネームする。

# 書式

**serial** [OPTIONS] DIRECTORY_PATH

# 説明
**serial**は、任意のディレクトリ以下にあるファイルの名前をユーザ指定の名前に連番を付与してリネームするCLIコマンドです。serialは、リネームしたファイルの格納先ディレクトリを指定できます。また、オリジナルファイルを保持したい場合、リネームではなくファイルコピーができます。

# 例
**カレントディレクトリにあるファイルの名前をシリアル番号付きのファイル名にリネームする。**

    $ ls
      a.txt  b.txt  c.txt
    $ serial --name demo  .
      Rename a.txt to demo_1.txt
      Rename b.txt to demo_2.txt
      Rename c.txt to demo_3.txt

**指定のディレクトリへファイルをコピー&リネーム**

    $ serial -p -k -n ../../dir/demo .
      Copy a.txt to ../../dir/0_demo.txt
      Copy b.txt to ../../dir/1_demo.txt
      Copy c.txt to ../../dir/2_demo.txt


# OPTIONS
**-d**, **--dry-run**
:   標準出力にファイル名のリネーム結果を表示します（ファイル更新はしません）。

**-f**, **--force**
:   同名のファイルが存在する場合であっても、強制的に上書き保存します。

**-h**, **--help**
:   ヘルプメッセージを表示します。

**-k**, **--keep**
:   リネーム前のファイルを保持します（リネームはせず、コピーします）。

**-n new_filename**, **--name=new_filename**
:   格納先のディレクトリ名を含んだ／含まないベースファイル名（このファイル名に連番を付与します）。

**-p**, **--prefix**
:   連番をファイル名の先頭に付与します。

**-s**, **--suffix**
:   連番をファイル名の末尾に付与します（デフォルト）。

**-v**, **--version**
:   serialコマンドのバージョンを表示します。

# 終了ステータス
**0**
:   成功

**1**
:   serialコマンドの引数指定でエラー

# バグ
GitHub Issuesを参照してください。URL：https://github.com/nao1215/serial/issues

# LICENSE
serialコマンドプロジェクトは、MITライセンス条文の下でライセンスされています。