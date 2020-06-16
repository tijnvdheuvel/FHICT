import requests


def api_call(url, family):
    api_location = "/api/v1/by_location/" + family
    command = url + api_location
    all_devices = []
    payload = {"history": 30, "num_scanners": 1}
    print("Calling: ", command)
    re = requests.get(command, params=payload)
    print("Call successful\nHttp status code: ", re.status_code)
    r = re.json()
    if re.status_code == 200 and r["success"] == True:
        i = 0
        while i < len(r["locations"]):

            name = r['locations'][i]["location"]
            lat = r['locations'][i]["gps"]["lat"]
            lon = r['locations'][i]["gps"]["lon"]
            tot = r['locations'][i]["total"]
            if "floor_0" in name:
                t = 0
            if "floor_1" in name:
                t = 1
            if "floor_2" in name:
                t = 2
            if "floor_3" in name:
                t = 3
            if "floor_4" in name:
                t = 4
            if "floor_5" in name:
                t = 5
            else:
                t = 0
            all_devices[t[0]].append(lat)
            all_devices[t[1]].append(lon)
            all_devices[t[2]].append(tot)
            i += 1
        print("API parsed")
        return all_devices
    else:
        print("External error")
        exit(5)

