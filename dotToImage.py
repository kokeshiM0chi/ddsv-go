import os
import pydot
import glob


def getFilePath(folder, name, ext):
    if not os.path.isdir(folder):
        os.makedirs(folder)

    return os.path.join(folder, name + ext)


name_list = []


def createNameList():
    files = glob.glob("./data/*")
    for file in files:
        name_list.append(file)


ext = "png"


def main(name_list, ext):
    createNameList()
    for name in name_list:
        # dot言語で記述されたファイルを読み込む
        print(name)
        (graph,) = pydot.graph_from_dot_file(name)
        # 保存用の名前を抽出する
        name, _ = os.path.splitext(os.path.basename(name))
        # 形式を選択して保存する
        if ext == "png":
            graph.write_png(getFilePath("./data-output/", name, ".png"))
        elif ext == "pdf":
            graph.write_pdf(getFilePath("./data-output/", name, ".pdf"))
        elif ext == "svg":
            graph.write_svg(getFilePath("./data-output/", name, ".svg"))
        else:
            print("[ERROR] ext option miss:", args.ext)


if __name__ == "__main__":
    main(name_list, ext)
