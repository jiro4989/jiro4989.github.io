## このプロジェクトのビルドツール

import os, sequtils, ospaths, osproc
from strutils import split, join
from strformat import `&`
import posix ## Unix依存 Windowsだと問題おきそう

proc buildHTML(fromDir, toDir: string) =
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
      
      stderr.writeLine &"[OK] Generated {movedFile}"
    except:
      stderr.writeLine &"[NG] Failed generating from {f}"
      stderr.writeLine getCurrentExceptionMsg()

when isMainModule:
  buildHTML("page", "docs")