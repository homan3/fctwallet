setup

newaddress fct fct01
newaddress fct fct02
newaddress fct fct03
newaddress fct fct04
newaddress fct fct05
newaddress fct fct06
newaddress fct fct07
newaddress fct fct08
newaddress fct fct09
newaddress fct fct10

newaddress ec ec01
newaddress ec ec02
newaddress ec ec03
newaddress ec ec04
newaddress ec ec05
newaddress ec ec06
newaddress ec ec07
newaddress ec ec08
newaddress ec ec09
newaddress ec ec10

newtransaction t
addinput t 01-Fountain 1001
addinput t 02-Fountain 1000
addoutput t   fct01 1000
addecoutput t ec01 1000

sign t
submit t
