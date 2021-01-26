$path = "."
$today = (Get-Date).AddMonths(-1)
$future = (Get-Date 2036-01-01)

Get-ChildItem $path -r | 
ForEach-Object {
    if ((Get-Date _.LastWriteTime) -ge future) {
        _.CreationTime = today
        _.LastAccessTime = today
        _.LastWriteTime = today
    }
}
