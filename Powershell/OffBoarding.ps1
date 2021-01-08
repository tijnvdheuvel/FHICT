#Setting defaults
$emailFinance = "[REDACTED]"
#Setting mail settings
$From = "[REDACTED]"; $SMTPServer = "smtp.gmail.com"; $SMTPPort = "465";
###Email
$username   = '[REDACTED]'
$password   = '[REDACTED]'
$secstr     = New-Object -TypeName System.Security.SecureString
$password.ToCharArray() | ForEach-Object {$secstr.AppendChar($_)}

Import-Csv "C:\Users\Administrator\Documents\OffBoarding.csv" -Delimiter ';' |
ForEach-Object {

#region DissableUser
Set-ADUser `
-Identity $_.samAccountNaam `
-Enabled $false
Get-ADUser -Identity $_.samAccountNaam | Move-ADObject -TargetPath "OU=Dissabled,OU=medewerkers,OU=challenge,DC=challenge,DC=local"
#Remove each user from the group
Remove-ADGroupMember -Identity "Medewerker" -Members $_.samAccountNaam -Confirm:$false
Remove-ADGroupMember -Identity "Users" -Members $_.samAccountNaam -Confirm:$false
#endregion AddUser

#region Mail
$Manager = (Get-ADUser -Identity $_.samAccountNaam).Manager
#Create mail message
$mailFinance = @{
    from       = $From
    to         = $emailFinance
    subject    = ("Verzoek papierwerk uitschrijving " + $_.Naam)
    smtpserver = "smtp.gmail.com"
    port       = "587"
    body       = ("Beste heer/mevrouw, `
Zou u het papier werk op orde willen maken voor  de uitschrijving van: " + $_.Naam + ". `
Het adress is: " + $_."address" + ", " + $_.plaatsnaam)
    credential = New-Object -typename System.Management.Automation.PSCredential -argumentlist $username, $secstr
    usessl     = $true
    verbose    = $true
}
Send-MailMessage @mailFinance; echo "mailFinance succesvol verzonden"
#endregion Mail

#Create network share
#region LockHome
$fullPath = "\\DC01\Homedirectories\" + $_.Naam
$User = Get-ADUser -Identity $_.samAccountNaam
    $homeShare = Get-Item -path $fullPath

    $acl = Get-Acl $homeShare

    $FileSystemRights = [System.Security.AccessControl.FileSystemRights]"Write, ReadAndExecute, Synchronize"
    $AccessControlType = [System.Security.AccessControl.AccessControlType]::Allow
    $InheritanceFlags = [System.Security.AccessControl.InheritanceFlags]"ContainerInherit, ObjectInherit"
    $PropagationFlags = [System.Security.AccessControl.PropagationFlags]"None"
    $AccessRule = New-Object System.Security.AccessControl.FileSystemAccessRule ($User.SID, $FileSystemRights, $InheritanceFlags, $PropagationFlags, $AccessControlType)
    $acl.RemoveAccessRule($AccessRule)

    Set-Acl -Path $homeShare -AclObject $acl -ea Stop

    Write-Host ("Removed permissions at {0}" -f $fullPath)
    #endregion LockHome
}

