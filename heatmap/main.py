
import pylab as pl
from matplotlib import pyplot as plt
from matplotlib import cm
from mpl_toolkits.axes_grid1 import make_axes_locatable

from heatmap import call

db = [[650, 800, 300, 1400], [200, 600, 400, 100], [10, 5, 2, 20], [0]]
x = db[0]  # location scanners LAT
y = db[1]  # location scanners LONG
u = db[2]  # n devices at location
# f = db[3]  # floor of the data set

# 1px=5.217391 cm   (verhouding BG: 1774 x 617 PX)



plotXmax = 1700
plotYmax = 610
DotScale = int(round(9.82 * plotXmax, 0))
#  DotScale = int(round(3.82 * plotXmax, 0))
# Omreken factor coördinaten berekenen
m = max(u)
ur = "https://cloud.internalpositioning.com"
fam = "testdb"


def get_data(url, family):
    api_db = call.api_call(url, family)
    final_db = []
    i = 0
    if i < len(api_db):
        if "verdieping_0" in api_db["location"]:
            x_cord = api_db[""]
            final_db[0].append(api)


def Heatmap(lat, long, u, cmap, vmin, vmax, imname, alph, s, plotXmax, plotYmax, pname):
    plt.title(pname)
    fig, ax = plt.subplots()
    img = plt.imread(imname)
    ax.imshow(img, extent=[0, plotXmax, 0, plotYmax])

    plt.title(pname)
    im = plt.scatter(lat, long, s=s, alpha=alph, c=u, cmap=cmap, vmin=vmin, vmax=vmax)

    #plt.axis('off')
    divider = make_axes_locatable(ax)
    cax = divider.append_axes("right", size="5%", pad=0.05)
    cbar = plt.colorbar(im, cax=cax)
    cbar.ax.get_yaxis().labelpad = 15
    cbar.ax.set_ylabel('aantal personen', rotation=270)
    a = "gevulde_" + imname

    plt.savefig(a, dpi=256, transparent=True)
    print("Heatmap [" + imname + "] done.")


Heatmap(lat=x,long=y,u=u,cmap=pl.cm.rainbow, vmin=1, vmax=m, imname="floorplan_0.png", s=DotScale, alph=0.4, plotYmax=plotYmax, plotXmax=plotXmax, pname="Heatmap Begane Grond")

#FHICT 19/20 Infra | Proftaak Groep B, Team A: Tijn, Daan, Jimmy, Nassim, Mika, Luc, Thomas



