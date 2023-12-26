wctl xlsx ./data --wpb-flag=1 --wpb-type=./wpb --go-flag=1 --go-type=./xlsx/configs/ --go-import-pb="github.com/walleframe/svc_xlsx/examples/xlsx/wpb" 
wctl gen -c wpb -o ./xlsx 
