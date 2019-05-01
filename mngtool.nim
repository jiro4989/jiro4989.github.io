## このプロジェクトのビルドツール

import os, sequtils, ospaths, osproc, terminal, times
from strutils import split, join, replace, startsWith
from strformat import `&`
from algorithm import sorted
import posix ## Unix依存 Windowsだと問題おきそう

let indexAdocTemplate = readFile("page/index.adoc.tmpl")

proc echoTaskTitle(title: string) =
    echo &"""
=====================================================================
  {title}
====================================================================="""

proc info(f: string) =
  styledEcho fgGreen, bgDefault, &"[OK] Generated {f}", resetStyle

proc err(f: string, prefix = "[NG] Failed generating from ") =
  styledEcho fgRed, bgDefault, prefix & f & ".", resetStyle

proc readPageTitle(path: string): string =
  for line in path.readFile.split("\n"):
    if line.startsWith("= "):
      return line.replace("= ", "")

proc buildIndexAdoc(dir: string, depth: int) =
  ## index.htmlの元になるindex.adocを生成する。
  if depth == 0:
    echoTaskTitle "Build index adoc"
  for k, f in walkDir(dir):
    if k == pcDir:
      try:
        var links: string
        # 先にディレクトリの一覧をセット
        links.add "* サブカテゴリ\n"
        for k2, f2 in walkDir(f):
          if k2 == pcDir:
            # 相対パス指定にするため最後のディレクトリ名だけ取得
            let f2fp = splitPath(f2)
            let title = f2fp[1]
            let nf = title / "index.html"
            links.add &"** link:./{nf}[{title}]\n"

        # ファイル一覧をセット
        links.add "* ページ\n"
        for k2, f2 in walkDir(f):
          if k2 == pcFile and f2.splitFile.name != "index":
            # 相対パス指定にするため最後のファイル名だけ取得
            let nf = f2.changeFileExt(".html").splitPath[1]
            var title = f2.readPageTitle
            links.add &"** link:./{nf}[{title}]\n"

        # テンプレートファイルにリンクなどを埋め込む
        let outFile = f / "index.adoc"
        var tmpl = indexAdocTemplate
        tmpl = tmpl.replace("{title}", f.split(AltSep)[1..^1].join($AltSep))
        tmpl = tmpl.replace("{metadataPath}", "../" & "..".repeat(depth).join("/") & "/metadata.txt[]")
        tmpl = tmpl.replace("{parentCategory}", "link:../index.html[こちら]")
        tmpl = tmpl.replace("{links}", links)
        writeFile(outFile, tmpl)

        info outFile
      except:
        err f
        err getCurrentExceptionMsg(), prefix="     "
      buildIndexAdoc(f, depth + 1)

proc buildHTML(fromDir, toDir: string) =
  ## AsciidocからHTMLを生成する。
  echoTaskTitle "Build HTML"
  # ビルドしたHTMLの配置先を最初に全部消す
  removeDir(toDir)
  createDir(toDir)
  for f in walkDirRec(fromDir):
    try:
      # asciidoc以外は無視
      let fp = splitFile(f)
      if fp.ext != ".adoc":
        continue

      # dockerでHTMLを生成
      # 生成されるファイルの所有権がrootにならないように指定
      let uid = getuid()
      let gid = getgid()
      discard execProcess(&"docker run --rm -u {uid}:{gid} -v {getCurrentDir()}:/documents/ asciidoctor/docker-asciidoctor asciidoctor -r asciidoctor-diagram {f}")

      # 生成したファイルと対応する配置先ディレクトリの生成
      let genedFile = f.changeFileExt(".html")
      let genedDir = toDir / fp.dir.split(AltSep)[1..^1].join($AltSep)
      createDir(genedDir)

      # 生成したディレクトリにファイルを移動
      let genedFp = genedFile.splitFile
      let movedFile = genedDir / genedFp.name & genedFp.ext
      moveFile(genedFile, movedFile)
      
      info movedFile
    except:
      err f
      err getCurrentExceptionMsg(), prefix="     "
  
proc getNewerWrittenPages(dir: string, pageCount: int): seq[tuple[path, lastWriteTime: string]] =
  ## 更新日時最新のものを指定数取得する。
  ## index.htmlはカウント対象から除く。
  var paths: seq[tuple[path: string, t: times.Time]]
  var pathsCount: int
  for fp in walkDirRec(dir, yieldFilter={pcFile}):
    if fp.splitFile.ext != ".adoc":
      continue
    if fp.splitFile.name == "index":
      continue
    var f = open(fp)
    let wt = fp.getFileInfo.lastWriteTime
    paths.add (fp, wt)
    f.close
    pathsCount.inc
  let pc = if pathsCount < pageCount: pathsCount else: pageCount
  result = paths.sorted(proc (x, y: tuple[path: string, t: times.Time]): int =
                          cmp(y.t.toSeconds, x.t.toSeconds))[0..<pc]
                .mapIt((it.path, it.t.format("yyyy/MM/dd HH:mm:ss")))

proc buildNewerWritternPages(dir: string, pageCount: int) =
  ## 更新日時最新のものを指定数取得し、一覧ファイルに出力する。
  echoTaskTitle "Build newer written pages"
  var s = &"""
最新の更新された記事一覧 ({pageCount}件)
-----------------------
"""
  # 最新のページ一覧を取得
  # 更新日時とともにAsciidoc記法のリストとして追加
  for f in getNewerWrittenPages(dir, pageCount):
    let url = f.path.split(AltSep)[1..^1].join($AltSep)
    let linkTitle = f.path.readPageTitle
    let writeTime = f.lastWriteTime
    s.add &"* link:./{url}[{linkTitle}] {writeTime} 更新\n"
  let outFile = "page/new-pages.txt"
  writeFile(outFile, s)
  info outFile

when false:
  buildIndexAdoc("page", 0)
  buildHTML("page", "docs")
else:
  buildNewerWritternPages("page", 10)