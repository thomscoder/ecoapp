import socket

hostpresent = False
react_app_host_present = False
hostname = socket.gethostname()
ip_address = socket.gethostbyname(hostname)
server_ws_port = input("Digit server port: ")
path="/"

with open('../.env','w') as file:
    with open('../.env','r') as read_file:
        for line in read_file:
            if "HOST" in line:
                hostpresent = True
        if hostpresent == False:
            file.write("HOST={ip} \nPATH={path} \nPORT={port}".format(ip=ip_address, path=path, port=server_ws_port))

with open('../client/.env','w') as file:
    with open('../client/.env','r') as read_var:
        for line in read_var:
            if "REACT_APP_HOST" in line:
                react_app_host_present = True
        if react_app_host_present == False:
            file.write("REACT_APP_ENDPOINT=http://{ip}: \nREACT_APP_HOST={ip} \nREACT_APP_P_PORT={p_port} \nREACT_APP_PORT={port} \n REACT_APP_PATH={path}".format(ip=ip_address, path=path, port=server_ws_port))
            
            
