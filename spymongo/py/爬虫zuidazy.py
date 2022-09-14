# -*- coding: UTF-8 -*-
import sys
reload(sys)
sys.setdefaultencoding('utf-8') 
import requests
from bs4 import BeautifulSoup as bs
import re
import json
import sys
url="http://zuidazy.net/index.php?m=vod-search"
kv={'user-agent':'Chrome/71.0.3578.98'}
name=sys.argv[1]
s=name.encode('utf-8')
params={
    "wd":s}
html=requests.post(url,data=params,headers=kv)
data = html.text
demo=data
soup=bs(demo,"html.parser")
p1=r"vod-detail-id-(\d+).html"
r1=re.compile(p1)
dic={}
m=0
for x in soup.find_all('a'):
    if re.search(r1,str(x)):
         h1=r1.findall(str(x))
         url1="http://zuidazy.net/?m=vod-detail-id-"+h1[0]+".html"
         r=requests.get(url1,headers=kv)
         r=r.text
         demo1=r
         soup1=bs(demo1,"html.parser")
         p2=r'name="copy_sel" type="checkbox" value="(.+)"'
         r2=re.compile(p2)
         p3=r'target="_blank">(.+)<span>'
         r3=re.compile(p3)
         h2=[]
         h3=r3.findall(str(x))
         dic4={}
         for y in soup1.find_all('input'):
             if re.search(r2,str(y)):
                 h2+=r2.findall(str(y))
         i=len(h2)/3
         i=int(i)
         dic1={}
         dic2={}
         dic3={}
         for a in range(i):
             b="EP"+str(a+1)
             dic1.update({b:h2[a]})
         for a in range(i):
             b="EP"+str(a+1)
             dic2.update({b:h2[a+i]})
         for a in range(i):
             b="EP"+str(a+1)
             dic3.update({b:h2[a+2*i]})
         dic4.update({"播放源1":dic1})
         dic4.update({"播放源2":dic2})
         dic4.update({"下载源":dic3})
         dic.update({h3[0]:dic4})
         m+=1
if m==0:
    print()
else:
    d=json.JSONEncoder().encode(dic)
    print(d)
