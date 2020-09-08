import argparse
import requests

# args
parser = argparse.ArgumentParser()
parser.add_argument('-u', '--username', type=str)
parser.add_argument('-p', '--password', type=str)
args = parser.parse_args()

username = args.username
password = args.password

# request
params = {
    'callback': 'dr1003',
    'DDDDD': username,
    'upass': password,
    '0MKKey': '123456',
    'R1': '0',
    'R2': '',
    'R3': '0',
    'R6': '0',
    'para': '00',
    'v6ip': '',
    'terminal_type': '1',
    'lang': 'zh-cn',
    'jsVersion': '4.1',
    # 'v': '9748'
}

r = requests.get('https://10.81.2.6/drcom/login',
                 params=params,
                 verify=False)
print(r.status_code)
print(r.text)
