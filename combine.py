import glob
import gzip
import json
from xml.etree import ElementTree

import fitparse

ACTIVITIES_GLOB = "/Users/matt/Google Drive/My Drive/strava data/activities/*"


def main():
    coords: list[list[list[float]]] = []
    count = 0
    for f in glob.glob(ACTIVITIES_GLOB):
        if f.endswith("fit.gz"):
            coords.append(get_fit_coords(f))
        else:
            root = ElementTree.parse(f).getroot()
            pts = root.findall(".//{http://www.topografix.com/GPX/1/1}trkpt")
            coords.append(
                [[float(pt.attrib["lon"]), float(pt.attrib["lat"])] for pt in pts]
            )
        count += 1
        if count % 10 == 0:
            print(f"processed: {count} activities")

    with open("result.geojson", "wb") as f:
        f.write(
            json.dumps(
                {
                    "type": "MultiLineString",
                    "coordinates": coords,
                },
                separators=(",", ":"),
            ).encode()
        )


# 2^32 / 360 - https://gis.stackexchange.com/questions/122186/convert-garmin-or-iphone-weird-gps-coordinates
CONVERSION = 11930465


def get_fit_coords(path: str):
    pts = []
    with gzip.open(path, "r") as f:
        fitfile = fitparse.FitFile(f.read())
        for record in fitfile.get_messages("record"):
            lat = record.get_value("position_lat")
            long = record.get_value("position_long")
            if lat and long:
                pts.append([long / CONVERSION, lat / CONVERSION])
    return pts


if __name__ == "__main__":
    main()
