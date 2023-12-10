default := 'help'

help:
  just -l

[private]
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

# Fetch AdventOfCode README, input and examples for DAY.
get DAY:
    #!/usr/bin/env sh
    set -e
    mkdir -p day{{DAY}}
    cd day{{DAY}}
    aoc download --day {{DAY}} --overwrite --input-file input.txt --puzzle-file README.md --session ../.aoc
    awk -f ../parts-from-md.awk README.md
    touch part1.txt part2.txt

# Initialize a new LANG project for solving DAY.
init LANG DAY:
    #!/usr/bin/env sh
    set -e
    just -q get {{DAY}}
    mkdir -p day{{DAY}}/{{LANG}}
    cd day{{DAY}}/{{LANG}}
    curl -sSL https://www.toptal.com/developers/gitignore/api/{{LANG}} -o .gitignore
    just -q init-{{LANG}} {{DAY}}

run-go DAY INPUT:
  #!/usr/bin/env sh
  cd day{{DAY}}/go
  go run main.go {{INPUT}}

# Run LANG project with the personalized input for DAY.
run LANG DAY:
    just run-{{LANG}} {{DAY}} $(pwd)/day{{DAY}}/input.txt
    # Run LANG project with the first example input from README.
run-part1 LANG DAY:
    just run-{{LANG}} {{DAY}} $(pwd)/day{{DAY}}/part1.txt
    # Run LANG project with the second example input from README.
run-part2 LANG DAY:
    just run-{{LANG}} {{DAY}} $(pwd)/day{{DAY}}/part2.txt

watch LANG DAY:
  watchexec --watch day{{DAY}}/{{LANG}} --workdir day{{DAY}}/{{LANG}} --restart --clear reset just -q run-{{LANG}} {{DAY}}


submit DAY PART VALUE:
    aoc submit --day {{DAY}} --session .aoc {{PART}} {{VALUE}}
