# Start server
$command = 'cmd /c start powershell -NoExit -Command {
    $host.UI.RawUI.WindowTitle = "Server";
    $host.UI.RawUI.BackgroundColor = "black";
    $host.UI.RawUI.ForegroundColor = "white";
    Clear-Host;
    cd server;
    go run .;
}'

Invoke-Expression -Command $command

# Start client
$command = 'cmd /c start powershell -NoExit -Command {
    $host.UI.RawUI.WindowTitle = "Client";
    $host.UI.RawUI.BackgroundColor = "black";
    $host.UI.RawUI.ForegroundColor = "white";
    Clear-Host;
    cd client;
    go run .;
}'

Invoke-Expression -Command $command