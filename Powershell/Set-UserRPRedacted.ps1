Connect-VIServer -Server  '192.168.128.200' -User 'administrator@vsphere.local' -Password '[Redacted]'
$ADGroup = "DG_Infrastructure-Semester 3"
$VspherePool = "MainPoolInfra"
$counter = 00
$UserList = Get-ADGroupmember -Identity $ADGroup | Select "SamAccountName"
$AantalPools = 10
$PoolPrefix = "RP-INF2-"
$i = 0

foreach ($user in $UserList){
if ($i > $AantalPools) {
exit
}

$i++
    [string]$number = $i
    if ($number.length -eq 1) {
        $number = "00" + $number
    } elseif ($number.length -eq 2) {
        $number = "0" + $number
    }
$VIAcc = Get-VIAccount -Id $user.SamAccountName -Domain "NETLAB"
$Pool = Get-ResourcePool -Name ($PoolPrefix + $number)
Get-VIPermission -Entity $Pool -Principal ("NETLAB\*")| Remove-VIPermission -Confirm:$false
New-VIPermission -role "ResourcePoolManager" -Principal ("NETLAB\" + $user.SamAccountName) -Entity $Pool > $null

}
Write-Host ("Er zijn: " + $i + " resource pools aangepast van het type: " + $PoolPrefix)
Disconnect-VIServer -Server '192.168.128.200' -Confirm:$false