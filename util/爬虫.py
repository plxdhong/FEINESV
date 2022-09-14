import requests
from bs4 import BeautifulSoup as bs
import re
import json
url="http://kuyunzyw.net/search.asp"
kv={'user-agent':'Mozilla/5.0'}
name=input('输入搜索名称：')
s=name.encode('gb2312')
params={
    "searchword":s}
html=requests.post(url,data=params,headers=kv)
data = html.text.encode("latin1").decode("gbk",'ignore')
demo=data
soup=bs(demo,"html.parser")
p1=r"/detail/\?(\d+).html"
r1=re.compile(p1)
dic={}
m=0
for x in soup.find_all('a'):
    if re.search(r1,str(x)):
         h1=r1.findall(str(x))
         url1="http://kuyunzyw.net/detail/?"+h1[0]+".html"
         r=requests.get(url1,headers=kv)
         r=r.text.encode("latin1").decode("gbk",'ignore')
         demo1=r
         soup1=bs(demo1,"html.parser")
         p2=r'name="copy_yah" type="checkbox" value="(.+)"/'
         r2=re.compile(p2)
         h2=[]
         for y in soup1.find_all('input'):
             if re.search(r2,str(y)):
                 h2+=r2.findall(str(y))
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
if m==0:
    dic.update({name:'无搜索结果'})
d=json.JSONEncoder().encode(dic)
print(d)
