Output assets/demo.gif
Require gh
Require ghs

Set FontSize 46
Set Width 2800
Set Height 1000
Set WindowBar Rings

Set Margin 20
Set MarginFill "#6B50FF"
Set BorderRadius 10

Type "gh gist list --filter test.sh"
Sleep 700ms
Enter
Sleep 2s

Type "gh gist view b9a1e66ac9dfcd2ffe4ce1e0a5f8ab46 -f test.sh -r"
Sleep 700ms
Enter
Sleep 2s

Type "ghs test.sh 'Hello World'"
Sleep 700ms
Enter
Sleep 7s