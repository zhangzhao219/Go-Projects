import pandas as pd
import matplotlib.pyplot as plt

data = pd.read_csv("temp.csv")
print(sum(data["count"])*(24*3600/288)*320/1024/1024/1024/2)
# myFig = plt.figure()
myFig = data.plot()
myFig.figure.savefig("myName.jpg")