My adventofcode 2023 solutions in golang to learn and mess around with the language

Error handling is mostly ignored for the purpose of aoc. This should have proper error handling, but for aoc a panic is
good enough debugging

The one thing I dislike about adventofcode is that there are plenty of ways to write valid code that passes examples but fails larger/real input sets

Random Snippets that were used here but don't need another file

Generating the files

Powershell
```powershell
mkdir inputs
for ($i = 1; $i -lt 26; $i++) { 
    mkdir day$1
    cd day$i
    $null > day$i.go
    Set-Content -Path .\day$i.go -Value "package day$i"
    $null > day$($i)_test.go
    Set-Content -Path .\day$($i)_test.go -Value "package day$($i)_test"
    cd ..
    cd inputs
    $null > day$i.txt
    cd ..
}
```

Fish (needs testing)
```fish
mkdir inputs
for i in (seq 1 25)
    cd day$i
    echo "package day$i" > day$i.go
    echo "package day$i_test" > day$i_test.go
    cd ..
    cd inputs
    touch day$i.txt
    cd ..
end
```