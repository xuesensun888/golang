import os
import subprocess
import time
import yaml

# Error Return Code
PING_TIMEOUT_ERROR = -1
SUBP_TIMEOUT_ERROR = -2

# LOG RECORD PERIOD, UNIT: SECOND
LOG_PERIOD = 604800  # 7days

# ORIN IP LIST
IP_LIST_1A = [
    '192.168.5.16',
    '192.168.5.32',
    '192.168.5.64',
    '192.168.6.16',
    '192.168.6.32',
    '192.168.6.64'
]
IP_LIST_1B = [
    '192.168.5.16',
    '192.168.5.32',
    '192.168.5.48',
    '192.168.6.16',
    '192.168.6.32',
    '192.168.6.48'
]
IP_LIST_2A = [
    '192.168.5.16',
    '192.168.5.48',
    '192.168.5.64',
    '192.168.6.16',
    '192.168.6.48',
    '192.168.6.64'
]
IP_LIST_2B = [
    '192.168.5.32',
    '192.168.5.48',
    '192.168.5.64',
    '192.168.6.32',
    '192.168.6.48',
    '192.168.6.64'
]

# Switch IP LIST
SWITCH_IP_LIST = [
    '192.168.14.201',
    '192.168.14.202',
    '192.168.15.201',
    '192.168.15.202'
]

# V2X IP LIST
V2X_IP_LIST = [
    '192.168.10.123',
    '192.168.11.123'
]

# IMU IP LIST
IMU_IP_LIST = [
    '192.168.9.201',
    '192.168.9.202'
]

# Ping Command
PING_TEST_CMD = 'ping {} -c {} -W {}'
# Ping Test Parameters
PING_COUNT = 1
PING_TIMEOUT = 1
SUBPROCESS_TIMEOUT = 1.5

def ping_test(ip, ping_count, ping_timeout, subp_timeout):
    ping_cmd = PING_TEST_CMD.format(ip, ping_count, ping_timeout)
    try:
        ping_ret = subprocess.run(ping_cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, timeout=subp_timeout)
    except subprocess.TimeoutExpired as e:
        return SUBP_TIMEOUT_ERROR
    else:
        if ping_ret.returncode != 0:
            return PING_TIMEOUT_ERROR
        else:
            return 0

def get_log_ts(log_name: str):
    log_time_list = log_name.split('.')[0].split('_')
    log_time = log_time_list[1] + "_" + log_time_list[2]
    log_time_struct = time.strptime(log_time, "%y-%m-%d_%H:%M:%S")
    log_ts = time.mktime(log_time_struct)
    return log_ts

def get_orin_id():
    get_orin_id_cmd = "cat /proc/device-tree/orin-id"
    try:
        orin_id_ret = subprocess.run(get_orin_id_cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, timeout=1)
    except subprocess.TimeoutExpired as e:
        return -3
    else:
        if orin_id_ret.returncode != 0:
            print(orin_id_ret.stderr.decode('utf-8'))
            return -4
        else:
            print(orin_id_ret.stdout.decode('utf-8'))
            orin_id = orin_id_ret.stdout.decode('utf-8')[0]
    return orin_id

def log_delete():
    file_list = os.listdir('./')
    file_list.sort()
    sys_ts = time.time()
    for file_name in file_list:
        if (file_name=="check.py") or (file_name=="result.log"):
            continue
        log_ts = get_log_ts(file_name)
        if (sys_ts - log_ts)>LOG_PERIOD:
            os.remove('./'+file_name)
        else:
            pass

test_result = {
    'orinID': '',
    'module': 'net',
    'check_result': 'OK',
    'project_list': {
        'public network': {
            'result': 'OK',
            'reason': ''
        },
        'DNS check': {
            'result': 'OK',
            'reason': ''
        },
        'Orin ethernet': {
            'result': 'OK',
            'reason': ''
        },
        'Switch ethernet': {
            'result': 'OK',
            'reason': ''
        },
        'MCU ethernet': {
            'result': 'OK',
            'reason': ''
        },
        'PAD ethernet': {
            'result': 'OK',
            'reason': ''
        },
        'V2X ethernet': {
            'result': 'OK',
            'reason': ''
        },
        'Tbox ethernet': {
            'result': 'OK',
            'reason': ''
        },
        'IMU ethernet': {
            'result': 'OK',
            'reason': ''
        }
    }
}

def check_public_network():
    public_network = ping_test("baidu.com", PING_COUNT, PING_TIMEOUT, SUBPROCESS_TIMEOUT)
    if public_network == -1:
        test_result['project_list']['public network']['result'] = "NG"
        test_result['project_list']['public network']['reason'] = "ping baidu.com timeout, cannot connect to the public network"
        return -1
    elif public_network == -2:
        test_result['project_list']['public network']['result'] = "NG"
        test_result['project_list']['public network']['reason'] = "subprocess ping command timeout"
        return -1
    else:
        return 0
    
def check_dns():
    DNS_check = ping_test("xiaojukeji.com", PING_COUNT, PING_TIMEOUT, SUBPROCESS_TIMEOUT)
    if DNS_check == -1:
        test_result['project_list']['DNS check']['result'] = "NG"
        test_result['project_list']['DNS check']['reason'] = "ping xiaojukeji.com timeout, cannot parser xiaojukeji.com"
        return -1
    elif DNS_check == -2:
        test_result['project_list']['DNS check']['result'] = "NG"
        test_result['project_list']['DNS check']['reason'] = "subprocess ping command timeout"
        return -1
    else:
        return 0

def check_orin_ethernet():
    orin_check = 0
    orin_reason = ''
    ping_orin_timeout_ip = []
    ping_orin_cmd_timeout_ip = []
    orin_id = get_orin_id()
    if orin_id == "1":
        ip_list = IP_LIST_1A
    elif orin_id == "2":
        ip_list = IP_LIST_1B
    elif orin_id == "3":
        ip_list = IP_LIST_2A
    elif orin_id == "4":
        ip_list = IP_LIST_2B
    else:
        print("Error: Invalid Orin ID")
        test_result['project_list']['Orin ethernet']['result'] = "NG"
        test_result['project_list']['Orin ethernet']['reason'] = "Invalid Orin ID"
        return -1       
    for ip in ip_list:
        orin_ethernet = ping_test(ip, PING_COUNT, PING_TIMEOUT, SUBPROCESS_TIMEOUT)
        if orin_ethernet == -1:
            orin_check = -1
            ping_orin_timeout_ip.append(ip)
        elif orin_ethernet == -2:
            orin_check = -1
            ping_orin_cmd_timeout_ip.append(ip)
        else:
            pass 
    if orin_check == -1:
        test_result['project_list']['Orin ethernet']['result'] = "NG"
        if len(ping_orin_timeout_ip) != 0:
            orin_reason = orin_reason + "ping {} timeout".format(ping_orin_timeout_ip)
        if len(ping_orin_cmd_timeout_ip) != 0:
            orin_reason = orin_reason + ", subprocess ping {} cmd timeout".format(ping_orin_cmd_timeout_ip)
        test_result['project_list']['Orin ethernet']['reason'] = orin_reason
        return -1
    else:
        return 0

def check_switch_ethernet():
    switch_check = 0
    switch_reason = ''
    ping_switch_timeout_ip = []
    ping_switch_cmd_timeout_ip = []
    for ip in SWITCH_IP_LIST:
        switch_ethernet = ping_test(ip, PING_COUNT, PING_TIMEOUT, SUBPROCESS_TIMEOUT)
        if switch_ethernet == -1:
            switch_check = -1
            ping_switch_timeout_ip.append(ip)
        elif switch_ethernet == -2:
            switch_check = -1
            ping_switch_cmd_timeout_ip.append(ip)
        else:
            pass
    if switch_check == -1:
        test_result['project_list']['Switch ethernet']['result'] = "NG"
        if len(ping_switch_timeout_ip) != 0:
            switch_reason = switch_reason + "ping {} timeout".format(ping_switch_timeout_ip)
        if len(ping_switch_cmd_timeout_ip) != 0:
            switch_reason = switch_reason + ", subprocess ping {} cmd timeout".format(ping_switch_cmd_timeout_ip)
        test_result['project_list']['Switch ethernet']['reason'] = switch_reason
        return -1
    else:
        return 0

def check_mcu_ethernet():
    mcu_ethernet = ping_test("192.168.0.200", PING_COUNT, PING_TIMEOUT, SUBPROCESS_TIMEOUT)
    if mcu_ethernet == -1:
        test_result['project_list']['MCU ethernet']['result'] = "NG"
        test_result['project_list']['MCU ethernet']['reason'] = "ping MCU timeout, cannot connect to the MCU"
        return -1
    elif mcu_ethernet == -2:
        test_result['project_list']['MCU ethernet']['result'] = "NG"
        test_result['project_list']['MCU ethernet']['reason'] = "subprocess ping command timeout"
        return -1
    else:
        return 0

def check_pad_ethernet():
    pad_ethernet = ping_test("192.168.5.10", PING_COUNT, PING_TIMEOUT, SUBPROCESS_TIMEOUT)
    if pad_ethernet == -1:
        test_result['project_list']['PAD ethernet']['result'] = "NG"
        test_result['project_list']['PAD ethernet']['reason'] = "ping PAD timeout, cannot connect to the PAD"
        return -1
    elif pad_ethernet == -2:
        test_result['project_list']['PAD ethernet']['result'] = "NG"
        test_result['project_list']['PAD ethernet']['reason'] = "subprocess ping command timeout"
        return -1
    else:
        return 0

def check_v2x_ethernet():
    v2x_check = 0
    v2x_reason = ''
    ping_v2x_timeout_ip = []
    ping_v2x_cmd_timeout_ip = []
    for ip in V2X_IP_LIST:
        v2x_ethernet = ping_test(ip, PING_COUNT, PING_TIMEOUT, SUBPROCESS_TIMEOUT)
        if v2x_ethernet == -1:
            v2x_check = -1
            ping_v2x_timeout_ip.append(ip)
        elif v2x_ethernet == -2:
            v2x_check = -1
            ping_v2x_cmd_timeout_ip.append(ip)
        else:
            pass
    if v2x_check == -1:
        test_result['project_list']['V2X ethernet']['result'] = "NG"
        if len(ping_v2x_timeout_ip) != 0:
            v2x_reason = v2x_reason + "ping {} timeout".format(ping_v2x_timeout_ip)
        if len(ping_v2x_cmd_timeout_ip) != 0:
            v2x_reason = v2x_reason + ", subprocess ping {} cmd timeout".format(ping_v2x_cmd_timeout_ip)
        test_result['project_list']['V2X ethernet']['reason'] = v2x_reason
        return -1
    else:
        return 0

def check_tbox_ethernet():
    tbox_ethernet = ping_test("192.168.5.200", PING_COUNT, PING_TIMEOUT, SUBPROCESS_TIMEOUT)
    if tbox_ethernet == -1:
        test_result['project_list']['Tbox ethernet']['result'] = "NG"
        test_result['project_list']['Tbox ethernet']['reason'] = "ping TBOX timeout, cannot connect to the TBOX"
        return -1
    elif tbox_ethernet == -2:
        test_result['project_list']['Tbox ethernet']['result'] = "NG"
        test_result['project_list']['Tbox ethernet']['reason'] = "subprocess ping command timeout"
        return -1
    else:
        return 0 

def check_imu_ethernet():
    imu_check = 0
    imu_reason = ''
    ping_imu_timeout_ip = []
    ping_imu_cmd_timeout_ip = []
    for ip in IMU_IP_LIST:
        imu_ethernet = ping_test(ip, PING_COUNT, PING_TIMEOUT, SUBPROCESS_TIMEOUT)
        if imu_ethernet == -1:
            imu_check = -1
            ping_imu_timeout_ip.append(ip)
        elif imu_ethernet == -2:
            imu_check = -1
            ping_imu_cmd_timeout_ip.append(ip)
        else:
            pass
    if imu_check == -1:
        test_result['project_list']['IMU ethernet']['result'] = "NG"
        if len(ping_imu_timeout_ip) != 0:
            imu_reason = imu_reason + "ping {} timeout".format(ping_imu_timeout_ip)
        if len(ping_imu_cmd_timeout_ip) != 0:
            imu_reason = imu_reason + ", subprocess ping {} cmd timeout".format(ping_imu_cmd_timeout_ip)
        test_result['project_list']['IMU ethernet']['reason'] = imu_reason
        return -1
    else:
        return 0 

def check_all_net():
    check_result_list = []
    # Check public network
    check_result = check_public_network()
    check_result_list.append(check_result)
    # Check DNS
    check_result = check_dns()
    check_result_list.append(check_result)
    # Check Orin ethernet connect
    if get_orin_id() == "5":
        pass
    else:
        check_result = check_orin_ethernet()
        check_result_list.append(check_result)
    # Switch ethernet connect
    if get_orin_id() == "5":
        pass
    else:
        check_result = check_switch_ethernet()
        check_result_list.append(check_result)
    # MCU ethernet connect, NOT CHECK NOW
    # check_mcu_ethernet()
    # check_result_list.append(check_result)
    
    # PAD ethernet connect
    check_result = check_pad_ethernet()
    check_result_list.append(check_result)
    # V2X ethernet connect
    check_result = check_v2x_ethernet()
    check_result_list.append(check_result)
    # TBOX ethernet connect
    check_result = check_tbox_ethernet()
    check_result_list.append(check_result)
    # IMU ethernet connect, NOT CHECK NOW
    # check_imu_ethernet()
    # check_result_list.append(check_result)
    
    # All Check Result Summary 
    if -1 in check_result_list:
        test_result['check_result'] = "NG"
    else: pass
    # Get Orin ID
    orin_id = get_orin_id()
    if orin_id == "1":
        test_result['orinID'] = "1A"
    elif orin_id == "2":
        test_result['orinID'] = "1B"
    elif orin_id == "3":
        test_result['orinID'] = "2A"
    elif orin_id == "4":
        test_result['orinID'] = "2B"
    elif orin_id == "5":
        test_result['orinID'] = "SL"
    else:
        test_result['orinID'] = "Invalid OrinID"
    return 0

check_all_net()
local_time = time.strftime("%y-%m-%d_%H:%M:%S")
log_name = 'result_{}.yaml'.format(local_time)
with open(log_name, 'w') as file:
    yaml.dump(test_result, file)
# Create Software Link
subprocess.run("rm result.log", shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, timeout=1)
subprocess.run("ln -s {} result.log".format(log_name), shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, timeout=1)
log_delete()