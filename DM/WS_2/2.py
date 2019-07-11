import requests
from bs4 import BeautifulSoup
import pandas as pd

link = 'https://en.wikipedia.org/wiki/List_of_Tamil_films_of_2018'

resp = requests.get(link)
if str(resp.status_code).startswith('2'):
    soup = BeautifulSoup(resp.text, 'lxml')
    tables = soup.findAll('table', {'class':'wikitable sortable'})
    for table in tables:
        df = pd.read_html(table)
        

else:
    print(resp)
