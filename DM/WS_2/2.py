import requests
from bs4 import BeautifulSoup
import pandas as pd
import os

link = 'https://en.wikipedia.org/wiki/List_of_Tamil_films_of_2018'
saveto = "C:\\Users\\TEMP.CS2K16.025\\Documents"

resp = requests.get(link)
if str(resp.status_code).startswith('2'):
    soup = BeautifulSoup(resp.text, 'lxml')
    tables = soup.findAll('table', {'class':'wikitable sortable'})
    print("Number of tables : ", len(tables))
    for index, table in enumerate(tables):
        df = pd.read_html(table.prettify(), flavor='bs4')
        df = df[0]
        new_header = df.iloc[0]
        df = df[1:]
        df.columns = new_header 
        name = table.find('caption')
        if not name:
            name = "table_" + str(index)
        else:
            name = name.text.strip()
        df.to_csv(os.path.join(saveto, name + ".csv"))
else:
    print(resp)
