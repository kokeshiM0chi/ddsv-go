


# 使用する構造体の説明
システムはプロセスの集合

## プロセスの構成要素について

### System.goの中身

Location: 状態を指示する番号を文字列でもつ
Rule: プロセスの遷移規則。ある状態（Location）における遷移ルールは複数ある。
これは状態遷移グラフにおいて、各状態の遷移先が複数あることを表している。

#### Process構造体の関数
- EntryPoint() 現在の状態を返却する
- EnterAt(l rule.Location) 最初の状態遷移。開始点の設定。


### rule/rule.goの中身
rule構造体
- 遷移元
- 遷移先
- ラベル（状態を表す人間が読める指示文字列）

### rule/vars/vars.goの中身
マルチプロセスで共有する変数。食事する哲学者問題ではフォークの本数を表す。


### rule/when/Guard.goの中身
Guard
与えられた共有変数を遷移可能かどうか判定する関数

未定義の変数の場合はエラーを返す


### 備考
生成したdotファイルを画像出力
https://github.com/kokeshiM0chi/sample-distributed-system/tree/master/sample-graphviz









# exampleについて
## philosoper.go





## Refs
https://ccvanishing.hateblo.jp/entry/2019/11/30/215950

