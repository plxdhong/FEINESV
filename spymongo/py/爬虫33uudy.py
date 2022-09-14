# -*- coding: UTF-8 -*-
import sys
reload(sys)
sys.setdefaultencoding('utf-8') 
import requests
from bs4 import BeautifulSoup as bs
import re
import json
import sys
url="http://www.156zy.cc/index.php?m=vod-search"
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
         url1="http://www.156zy.cc/?m=vod-detail-id-"+h1[0]+".html"
         r=requests.get(url1,headers=kv)
         r=r.text
         demo1=r
         soup1=bs(demo1,"html.parser")
         p2=r'name="copy_sel" type="checkbox" value="(.+)"'
         r2=re.compile(p2)
         h2=[]
         num=0
         for y in soup1.find_all('input'):
             if re.search(r2,str(y)):
                 h2+=r2.findall(str(y))
         p3=r'33u'
         r3=re.compile(p3)
         for z in soup1.find_all('h3'):
             if re.search(r3,z.string):
                 num+=1
         if num==2:
             i=len(h2)/2
             i=int(i)
             dic1={}
             dic2={}
             for a in range(i):
                 b="EP"+str(a+1)
                 dic1.update({b:h2[a]})
             for a in range(i):
                 b="EP"+str(a+1)
                 dic2.update({b:h2[a+i]})
             dic3={}
             dic3.update({"播放源1":dic1})
             dic3.update({"播放源2":dic2})
             dic.update({x.string:dic3})
             m+=1
         else:
             i=int(len(h2))
             dic1={}
             for a in range(i):
                 b="EP"+str(a+1)
                 dic1.update({b:h2[a]})
             dic3={}
             dic3.update({"播放源1":dic1})
             dic.update({x.string:dic3})
             m+=1  
if m==0:
    print()
else:
    d=json.JSONEncoder().encode(dic)
    print(d)
