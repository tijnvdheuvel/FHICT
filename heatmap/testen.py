import pylab as pl
from matplotlib import pyplot as plt
from mpl_toolkits.axes_grid1 import make_axes_locatable
from datetime import datetime
import call
import sys

demo_db = [[51.45111, 51.451247, 51.451438, 51.451441], [5.479317, 5.479945, 5.480229, 5.479414], [10, 5, 2, 20]]
y = demo_db[0]  # location scanners LAT
x = demo_db[1]  # location scanners LONG
users = demo_db[2]  # n devices at location
# 1px=5.217391 cm   (verhouding BG: 1774 x 617 PX)
# Boundary box
m = max(users)
min_lon = 5.47923
min_lat = 51.45139
max_lon = 5.48055
max_lat = 51.45098

DotScale = 5000  # int(round(9.82 * plotXmax, 0))

# API settings:
bbox = (min_lon, min_lat, max_lon, max_lat)
# Current time:
now = datetime.now()
current_time = now.strftime("%H:%M.%S sec")
current_date = now.strftime("%x")


def startup():
    if len(sys.argv) > 1:
        if sys.argv[1] == "detached" and sys.argv[2] and sys.argv[3]:
            db = get_data(sys.argv[2], sys.argv[3])
            n = 0
            while n < len(db):
                floor = n
                title = "Heatmap floor: ", floor
                t = str(floor)
                imclean = "floorplan_" + t + ".png"
                time = now.strftime("%H:%M.%S sec")
                heatmap(lat=db[floor][1], long=db[floor][0], c=db[floor][2], cmap=pl.cm.rainbow, boundry=bbox, vmin=1,
                    vmax=m, imname=imclean, s=DotScale, alph=0.4, pname=title, time=time,
                    description=("Heatmap of floor ", floor))
                n +=1
        elif sys.argv[1] == "user":
            introduction()

        elif sys.argv[1] == "help":
            print("Welcome to the heatmap generator")
            print("This is part of the FIND3 project on Fontys Hogeschool ICT Eindhoven.\n")
            print("Please ensure that the floorplan you're using has the format: floorplan_0.png")
            print("For user mode, user parameter [user]. \n"
              "If you'd like to run this without further interaction use parameter [live] followed by syntax as "
              "[http:localhost:8008] and [familyname] so: heatmap.py live http:localhost:8005 familyname")
        else:
            print("invalid input check parameter [help] for further explanations1")
    else:
        print("invalid input check parameter [help] for further explanations")


def get_data(url, family):
    print("Calling api at: ", url, family)
    db = call.api_call(url, family)
    return db


def heatmap(lat, long, c, cmap, vmin, vmax, imname, alph, s, pname, boundry, time, description):
    print("Starting heatmap generation at: ", current_date, "|", current_time)
    fig, ax = plt.subplots()
    ax.set_xlim(left=boundry[0], right=boundry[2],)
    ax.set_ylim(top=boundry[3], bottom=boundry[1])
    img = plt.imread(imname)
    ax.imshow(img, extent=[boundry[0], boundry[2], boundry[1], boundry[3]], origin=[5.47923, 51.45108])
    plt.title(pname)
    plot = ax.scatter(lat, long, s=s, alpha=alph, c=c, cmap=cmap, vmin=vmin, vmax=vmax, )

    plt.axis('off')

    divider = make_axes_locatable(ax)
    cax = divider.append_axes("right", size="5%", pad=0.05)
    cbar = plt.colorbar(plot, cax=cax)
    cbar.ax.get_yaxis().labelpad = 15
    cbar.ax.set_ylabel('aantal personen', rotation=270)

    a = "heatmap_" + imname
     #im.figsize = (18.5, 6.5)
    plt.savefig(fname=a, dpi=320, transparent=True, metadata={"Description": description, "Author": "Proftaak 19/20",
                                                              "Creation Time": time})

    print("Heatmap of [" + imname + "] done.")


def demo_heatmap(lat, long, c, cmap, vmin, vmax, imname, alph, s, pname, boundry, time):
    print("Starting demo heatmap generation at: ", current_date, "|", current_time)
    fig, ax = plt.subplots()
    ax.set_xlim(left=boundry[0], right=boundry[2],)
    ax.set_ylim(top=boundry[3], bottom=boundry[1])
    img = plt.imread(imname)
    ax.imshow(img, extent=[boundry[0], boundry[2], boundry[1], boundry[3]], origin=[5.47923, 51.45108])
    plt.title(pname)
    plot = ax.scatter(lat, long, s=s, alpha=alph, c=c, cmap=cmap, vmin=vmin, vmax=vmax, )

    plt.axis('off')

    divider = make_axes_locatable(ax)
    cax = divider.append_axes("right", size="5%", pad=0.05)
    cbar = plt.colorbar(plot, cax=cax)
    cbar.ax.get_yaxis().labelpad = 15
    cbar.ax.set_ylabel('aantal personen', rotation=270)

    a = "demo_heatmap_" + imname
     #im.figsize = (18.5, 6.5)
    plt.savefig(fname=a, dpi=320, transparent=True, metadata={"Title": "Joeeee", "Author": "Proftaak 19/20",
                                                              "Creation Time": time})

    print("Heatmap of [" + imname + "] done.")


def introduction():
    print("Welcome to the heatmap generator")
    print("This is part of the FIND3 project on Fontys Hogeschool ICT Eindhoven.\n")
    print("Please ensure that the floorplan you're using has the format: floorplan_0.png")
    print("Please specify the type of operation (demo or live) [e.g. live]")
    user_in = input()
    if user_in == "live":
        print("You've chose live mode\nPlease ensure that you have a HTTP connection to your FIND3 server.\n============")
        print("Please specify the address of the FIND3 server [e.g. http://localhost:8005]")
        user_in = input()
        u = user_in
        print("Input host: ", user_in)
        print("Please specify the family [e.g. testdb]")
        user_in = input()
        f = user_in
        print("Input family: ", user_in, "\n============")
        db = get_data(u, f)
        print("Please specify the floor you want to use: (e.g. 1)")
        n = 0
        while n < len(db):
            print("| ", n)
            n +=1
        user_in = input()
        floor = int(user_in)
        print("Floor used: ", user_in)
        title = "Heatmap floor: ", floor
        tussen = str(floor)
        imclean = "floorplan_" + tussen + ".png"
        heatmap(lat=db[floor][1], long=db[floor][0], c=db[floor][2], cmap=pl.cm.rainbow, boundry=bbox, vmin=1, vmax=m,
                imname=imclean, s=DotScale, alph=0.4, pname=title, time=current_time, description=("Heatmap of floor ",
                                                                                                   floor))
    else:
        print("You've chose demo mode")
        demo_heatmap(lat=demo_db[1], long=demo_db[0], c=demo_db[2], cmap=pl.cm.rainbow, boundry=bbox, vmin=1, vmax=m,
                imname="floorplan_0.png", s=DotScale, alph=0.4, pname="Heatmap Begane Grond", time=current_time)


startup()
#FHICT 19/20 Infra | Proftaak Groep B, Team A: Tijn, Daan, Jimmy, Nassim, Mika, Luc, Thomas