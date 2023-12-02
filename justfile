[private]
help:
  just -l


[private]
init-scala DAY:
  #!/usr/bin/env sh
  cd day{{DAY}}/scala
  cat <<-EOF > Main.scala
  package aoc.day{{DAY}}
  object Main extends App { println("Hello day{{DAY}}") }
  EOF

[private]
init-rust DAY:
  #!/usr/bin/env sh
  cd day{{DAY}}/rust
  cargo init --name "day{{DAY}}"

[private]
init-zig DAY:
  #!/usr/bin/env sh
  cd day{{DAY}}/zig
  zig init-exe

[private]
init-flix DAY:
  #!/usr/bin/env sh
  cd day{{DAY}}/flix
  flix init
init-go DAY:
  #!/usr/bin/env sh
  cd day{{DAY}}/go
  go mod init "day{{DAY}}"
  cat <<-EOF > main.go
  package main
  import "fmt"
  func main() {
      fmt.Println("Hello day{{DAY}}!")
  }
  EOF

init LANG DAY:
  #!/usr/bin/env sh
  mkdir -p day{{DAY}}/{{LANG}}
  aoc download --day {{DAY}} --overwrite --input-file day{{DAY}}/{{LANG}}/input.txt --puzzle-file day{{DAY}}/README.md
  cd day{{DAY}}/{{LANG}}
  just -q init-{{LANG}} {{DAY}}

[private]
run-scala DAY:
  #!/usr/bin/env sh
  cd day{{DAY}}/scala
  scala-cli run .

[private]
run-rust DAY:
  #!/usr/bin/env sh
  cd day{{DAY}}/rust
  cargo run

[private]
run-zig DAY:
  #!/usr/bin/env sh
  cd day{{DAY}}/zig
  zig run

[private]
run-flix DAY:
  #!/usr/bin/env sh
  cd day{{DAY}}/flix
  flix run

[private]
run-go DAY:
  #!/usr/bin/env sh
  cd day{{DAY}}/go
  go run .

run LANG DAY:
  just -q run-{{LANG}} {{DAY}}

watch LANG DAY:
  watchexec --watch day{{DAY}}/{{LANG}} --workdir day{{DAY}}/{{LANG}} --restart --clear reset just -q run-{{LANG}} {{DAY}}
  

