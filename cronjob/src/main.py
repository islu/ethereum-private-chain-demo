import os
import schedule
import requests
import time

host = os.getenv('HOST', 'localhost')
port = os.getenv('PORT', '8080')
target_address = os.getenv('TARGET_ADDRESS', '0x8De0c53FC169BA09F111aA4170697e8CF42CCbBe')

def job():
    print("執行任務...")
    url = f"http://{host}:{port}/api/v1/chain/tx/{target_address}/sync"

    # print(url)

    r = requests.post(url)

    print(r.text)

# 設定每 30 秒執行一次
schedule.every(30).seconds.do(job)

while True:
    schedule.run_pending()
    time.sleep(1)