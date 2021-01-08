#Setting defaults
$emailFinance = "[REMOVED]"
$emailInkoop = "[REMOVED]"
#Setting mail settings
$From = "[REMOVED]"; $SMTPServer = "smtp.gmail.com"; $SMTPPort = "465";
###Email
$username   = '[REMOVED]'
$password   = '[REMOVED]'
$secstr     = New-Object -TypeName System.Security.SecureString
$password.ToCharArray() | ForEach-Object {$secstr.AppendChar($_)}

Import-Csv "C:\Users\Administrator\Documents\OnBoarding.csv" -Delimiter ';' |
ForEach-Object {
#Create home directory
#region AddUser
New-ADUser `
-Name $_.Naam `
-GivenName $_.voornaam `
-Surname $_.achternaam `
-Path $_."ParentOU" `
-SamAccountName $_.samAccountNaam `
-UserPrincipalName ($_.samAccountNaam + '@' + $env:userdnsdomain) `
-AccountPassword (ConvertTo-SecureString "123user!!!!!" -AsPlainText -Force) `
-EmailAddress $_."E-Mail" `
-Enabled $true `
-ChangePasswordAtLogon $true `
-StreetAddress $_."address" `
-City $_.plaatsnaam
echo ("User: " + $_.Naam + " toegevoegd." )
#Add each user to the group specified in the file
Add-ADGroupMember `
-Identity $_.GroupNaam `
-Members $_.SamAccountNaam
#endregion AddUser

#region Mail
#Create mail message
$mailFinance = @{
    from       = $From
    to         = $emailFinance
    subject    = ("Verzoek papierwerk " + $_.Naam)
    smtpserver = "smtp.gmail.com"
    port       = "587"
    body       = ("Beste heer/mevrouw, `
Zou u het papier werk op orde willen maken voor: " + $_.Naam + ". `
Het adress is: " + $_."address" + ", " + $_.plaatsnaam)
    credential = New-Object -typename System.Management.Automation.PSCredential -argumentlist $username, $secstr
    usessl     = $true
    verbose    = $true
}
$mailInkoop = @{
    from       = $From
    to         = $emailInkoop
    subject    = ("Verzoek werkkleding " + $_.Naam)
    smtpserver = "smtp.gmail.com"
    port       = "587"
    body       = ("Beste heer/mevrouw, `
Zou u een bestelling voor een nieuwe set werkkleding willen plaatsen voor: " + $_.Naam + ". `
Het adress is: " + $_."address" + ", " + $_.plaatsnaam + ". `
Met vriendelijke groet, `
Het IT beheer")
    credential = New-Object -typename System.Management.Automation.PSCredential -argumentlist $username, $secstr
    usessl     = $true
    verbose    = $true
}
#Send-MailMessage @mailFinance; echo "mailFinance succesvol verzonden"
#Send-MailMessage @mailInkoop; echo "mailinkoop succesvol verzonden"
#endregion Mail

#Create network share
#region SetHome
$fullPath = "\\DC01\Homedirectories\" + $_.Naam
$User = Get-ADUser -Identity $_.samAccountNaam
    Set-ADUser $User -HomeDrive "Z:" -HomeDirectory $fullPath -ea Stop
    $homeShare = New-Item -path $fullPath -ItemType Directory -force -ea Stop

    $acl = Get-Acl $homeShare

    $FileSystemRights = [System.Security.AccessControl.FileSystemRights]"Write, ReadAndExecute, Synchronize"
    $AccessControlType = [System.Security.AccessControl.AccessControlType]::Allow
    $InheritanceFlags = [System.Security.AccessControl.InheritanceFlags]"ContainerInherit, ObjectInherit"
    $PropagationFlags = [System.Security.AccessControl.PropagationFlags]"None"

    $AccessRule = New-Object System.Security.AccessControl.FileSystemAccessRule ($User.SID, $FileSystemRights, $InheritanceFlags, $PropagationFlags, $AccessControlType)
    $acl.AddAccessRule($AccessRule)

    Set-Acl -Path $homeShare -AclObject $acl -ea Stop

    Write-Host ("HomeDirectory created at {0}" -f $fullPath)
}
#endregion SetHome
